package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/nesv/pm"
	"github.com/spf13/cobra"
)

var UnpackCmd = &cobra.Command{
	Use:   "unpack [name] [version]",
	Short: "Unpack a cached version of a package",
	Run:   runUnpack,
}

func runUnpack(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		log.Fatalln("not enough arguments")
	}

	pkgFilename := fmt.Sprintf("%s-%s-%s-%s.tar.gz",
		args[0], args[1], runtime.GOOS, runtime.GOARCH)
	pkgPath := filepath.Join(rootCacheDir, pkgFilename)

	if err := unpack(pkgPath); err != nil {
		log.Fatalln(err)
	}
}

func unpack(pkgPath string) error {
	pkgFilename := filepath.Base(pkgPath)

	f, err := os.Open(pkgPath)
	if err != nil {
		return fmt.Errorf("error: package file %q is not cached", pkgFilename)
	}
	defer f.Close()

	log.Println("unpacking", pkgFilename)

	unpackedFiles, err := pm.Unpack(rootBaseDir, f)
	if err != nil {
		return err
	}

	for _, fname := range unpackedFiles {
		log.Println("unpacked", fname)
	}

	return nil
}