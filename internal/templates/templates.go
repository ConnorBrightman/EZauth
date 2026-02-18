package templates

import (
	"embed"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

//go:embed demo/*
var templateFiles embed.FS

func GenerateTemplates() error {
	publicDir := "public"

	if err := os.MkdirAll(publicDir, 0755); err != nil {
		return err
	}

	sub, err := fs.Sub(templateFiles, "demo")
	if err != nil {
		return err
	}

	return fs.WalkDir(sub, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("❌ Error walking %s: %v", path, err)
			return nil
		}

		if path == "." {
			return nil
		}

		destPath := filepath.Join(publicDir, path)

		if d.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}

		data, err := fs.ReadFile(sub, path)
		if err != nil {
			log.Printf("❌ Failed to read %s: %v", path, err)
			return nil
		}

		if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
			return err
		}

		if err := os.WriteFile(destPath, data, 0644); err != nil {
			log.Printf("❌ Failed to write %s: %v", destPath, err)
		} else {
			log.Printf("✅ Created %s", destPath)
		}

		return nil
	})
}
