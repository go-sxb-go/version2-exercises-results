package files

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

type CountCharWorker struct {
	dir  string
	char rune
	sum  int

	sumChan chan int
	errs    chan error
}

// This function will count the amount of a given character of a type in all files of a directory
func CountCharInFiles(dir string, c rune) (int, error) {
	worker := &CountCharWorker{
		dir:     dir,
		char:    c,
		sum:     0,
		sumChan: make(chan int),
		errs:    make(chan error),
	}

	go worker.BrowseFile()
	go worker.ComputeSum()

	for err := range worker.errs {
		if err != nil {
			return -1, err
		}
	}

	return worker.sum, nil
}

func (w *CountCharWorker) ComputeSum() {
	defer close(w.errs)
	for value := range w.sumChan {
		w.sum += value
	}
}

func (w *CountCharWorker) HandleFile(file string, wg *sync.WaitGroup) {
	defer wg.Done()
	content, err := ioutil.ReadFile(file)
	if err != nil {
		w.errs <- err
		return
	}

	w.sumChan <- bytes.Count(content, []byte(string(w.char)))
}

func (w *CountCharWorker) BrowseFile() {
	defer close(w.sumChan)
	wg := &sync.WaitGroup{}

	err := filepath.Walk(w.dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		wg.Add(1)
		go w.HandleFile(path, wg)
		return nil
	})
	if err != nil {
		w.errs <- err
	}
	wg.Wait()
}
