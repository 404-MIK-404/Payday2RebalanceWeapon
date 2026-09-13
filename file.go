package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func createFolder(folder string) error {
	err := os.MkdirAll(folder, 0755)
	if err != nil {
		return fmt.Errorf("Failed to create folder %q: %w", folder, err)
	}
	fmt.Println("Folder created:", folder)
	return nil
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open %s: %w", src, err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create %s: %w", dst, err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy %s -> %s: %w", src, dst, err)
	}

	return nil
}

func renderTemplate(src string, values map[string]string) error {
	if filepath.Base(src) != "Weapon_Template_Rebalance.lua" {
		return nil
	}

	fileTemplate, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open %q: %w", src, err)
	}

	tmpPath := src + ".tmp"
	tmpFile, err := os.Create(tmpPath)
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}

	writer := bufio.NewWriter(tmpFile)
	scanner := bufio.NewScanner(fileTemplate)

	for scanner.Scan() {
		line := scanner.Text()

		for key, value := range values {
			placeholder := "${" + key + "}"
			line = strings.ReplaceAll(line, placeholder, value)
		}

		writer.WriteString(line + "\n")
	}

	if err := scanner.Err(); err != nil {
		tmpFile.Close()
		return fmt.Errorf("failed to read %q: %w", src, err)
	}

	writer.Flush()
	tmpFile.Close()
	fileTemplate.Close()

	err = os.Rename(tmpPath, src)
	if err != nil {
		return fmt.Errorf("failed to rename %q: %w", src, err)
	}

	return nil
}
