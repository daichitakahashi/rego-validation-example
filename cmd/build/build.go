package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"flag"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/compile"
)

var (
	src, dstWasm, dstBundle string
)

func init() {
	flag.StringVar(&src, "src", "", "source directory")
	flag.StringVar(&dstWasm, "wasm", "", "filename of wasm")
	flag.StringVar(&dstBundle, "bundle", "", "filename of bundle")
	flag.Parse()
}

func main() {
	// see: https://github.com/open-policy-agent/opa/blob/main/cmd/build.go

	if src == "" {
		log.Fatal("source directory(-src) required")
	}
	if dstWasm == "" {
		log.Fatal("filename of wasm file(-wasm) required")
	}
	if dstBundle == "" {
		log.Fatal("filename of bundle file(-bundle) required")
	}

	buf := bytes.NewBuffer(nil)

	compiler := compile.New().
		WithCapabilities(ast.CapabilitiesForThisVersion()).
		WithTarget("wasm").
		WithAsBundle(true).
		WithOptimizationLevel(2).
		WithOutput(buf).
		WithEntrypoints("validation/email", "validation/domain").
		WithRegoAnnotationEntrypoints(true).
		WithPaths(src).
		WithFilter(func(abspath string, info fs.FileInfo, depth int) bool {
			return !info.IsDir() && strings.HasSuffix(abspath, "_test.rego")
		}).
		WithBundleVerificationKeyID("default")

	err := compiler.Build(context.Background())
	if err != nil {
		log.Panic(err)
	}

	gr, err := gzip.NewReader(buf)
	if err != nil {
		log.Panic(err)
	}
	defer gr.Close()
	tr := tar.NewReader(gr)

	wasm, err := os.Create(dstWasm)
	if err != nil {
		log.Panic(err)
	}
	defer wasm.Close()

	bundle, err := os.Create(dstBundle)
	if err != nil {
		log.Panic(err)
	}
	defer bundle.Close()
	gw := gzip.NewWriter(bundle)
	defer gw.Close()
	tw := tar.NewWriter(gw)

	for {
		h, err := tr.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Panic(err)
		}
		switch strings.TrimPrefix(h.Name, "/") {
		case "policy.wasm":
			_, err = io.Copy(wasm, tr)
			if err != nil {
				log.Panic(err)
			}
		case ".manifest":
			continue
		default:
			// Re-bundle data.json and ".rego" files
			err = tw.WriteHeader(h)
			if err != nil {
				log.Panic(err)
			}
			_, err = io.Copy(tw, tr)
			if err != nil {
				log.Panic(err)
			}
			err = tw.Flush()
			if err != nil {
				log.Panic(err)
			}
		}
	}
}
