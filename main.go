package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"kubedirlint/summary"
	"os"
	"path/filepath"
	"sort"
)

var changeDirFlag = flag.String("C", ".", "change the directory before linting")

func main() {
	flag.Parse()

	err := os.Chdir(*changeDirFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fs, err := getFiles()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to find YAML files: %s\n", err)
		os.Exit(1)
	}

	if len(fs) == 0 {
		fmt.Fprint(os.Stderr, "got 0 YAML files\n")
		os.Exit(1)
	}

	t := &summary.Template{}
	es := checkFiles(t, fs)

	if len(es) != 0 {
		format := "got %d errors\n"
		if len(es) == 1 {
			format = "got %d error\n"
		}
		fmt.Printf(format, len(es))

		w := bufio.NewWriter(os.Stderr)
		for _, e := range es {
			fmt.Fprintln(w, e)
		}
		w.Flush()

		os.Exit(1)
	}

	format := "all %d files are OK\n"
	if len(fs) == 1 {
		format = "%d file is OK\n"
	}
	fmt.Printf(format, len(fs))
}

func getFiles() ([]string, error) {
	var files []string
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		ext := filepath.Ext(path)
		regular := d.Type().IsRegular()
		if regular && (ext == ".yaml" || ext == ".yml") {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	sort.Strings(files)
	return files, nil
}

func checkFiles(t *summary.Template, fs []string) (es []error) {
	for _, f := range fs {
		if err := checkFile(t, f); err != nil {
			es = append(es, err)
		}
	}

	return
}

func checkFile(t *summary.Template, f string) error {
	d, err := os.ReadFile(f)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", f, err)
	}

	s, err := summary.Summarize(d)
	if err != nil {
		return fmt.Errorf("failed to summarize %s: %w", f, err)
	}

	p, err := t.SuggestPath(s)
	if err != nil {
		return fmt.Errorf("failed to propose path for %s: %w", f, err)
	}

	if p != f {
		return fmt.Errorf("file %s has wrong path, should be %s", f, p)
	}
	return nil
}
