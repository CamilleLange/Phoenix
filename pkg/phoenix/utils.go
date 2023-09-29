package phoenix

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// ReplaceStringInFiles replaces a string in all Go files in a folder and its subfolders.
func ReplaceStringInFiles(rootDir, search, replace string) error {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	err := filepath.Walk(rootDir, func(filePath string, info os.FileInfo, err error) error {
		// Breaking recursion if previous iteration raise an error.
		if err != nil {
			return err
		}

		wg.Add(1)
		go func(filePath string, info os.FileInfo) {
			defer wg.Done()

			if strings.Contains(filePath, string(os.PathSeparator)+"vendor"+string(os.PathSeparator)) {
				return // ignoring vendor
			}

			if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
				var localErr error
				localErr = ReplaceStringIn(filePath, filePath+".tmp", search, replace)

				mu.Lock()
				defer mu.Unlock()

				if localErr != nil {
					err = localErr
				}
			}
		}(filePath, info)

		return err
	})

	if err != nil {
		return fmt.Errorf("error while replacing in files : %w", err)
	}

	wg.Wait()

	return nil
}

// ReplaceStringIn replaces a string in a file.
func ReplaceStringIn(inputFilePath, tmpFilePath, search, replace string) error {
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return fmt.Errorf("can't open the input file : %w", err)
	}
	defer inputFile.Close()

	tmpFile, err := os.Create(tmpFilePath)
	if err != nil {
		return fmt.Errorf("can't create the output file : %w", err)
	}
	defer tmpFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(tmpFile)

	for scanner.Scan() {
		newLine := strings.Replace(scanner.Text(), search, replace, -1)

		_, err := fmt.Fprintln(writer, newLine)
		if err != nil {
			return fmt.Errorf("can't write to the output file : %w", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("can't read the input file : %w", err)
	}

	// Ensure all data is written to the output file
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("can't flush file internal buffer : %w", err)
	}

	// Close the files
	err = inputFile.Close()
	if err != nil {
		return fmt.Errorf("can't close input file : %w", err)
	}

	err = tmpFile.Close()
	if err != nil {
		return fmt.Errorf("can't close output file : %w", err)
	}

	// Replace the input file with the output file
	err = os.Rename(tmpFilePath, inputFilePath)
	if err != nil {
		return fmt.Errorf("error replacing the input file : %w", err)
	}

	return nil
}
