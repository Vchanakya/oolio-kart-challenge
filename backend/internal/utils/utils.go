package utils

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

// HasCouponInAtLeastTwo scans all files concurrently and
// returns true when the coupon appears in ≥2 distinct files.
func HasCouponInAtLeastTwo(files []string, coupon string) bool {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	var hits int32 // atomic counter of matching files

	for _, fn := range files {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()
			if scanFile(ctx, filename, coupon) {
				// increment hits; if this is the 2nd match, stop everything else
				if atomic.AddInt32(&hits, 1) == 2 {
					cancel()
				}
			}
		}(fn)
	}

	wg.Wait()
	return atomic.LoadInt32(&hits) >= 2
}

// scanFile streams one file line-by-line, returns true on first match.
func scanFile(ctx context.Context, path, coupon string) bool {
	filePath := filepath.Join("coupons", path)
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return false
	}
	defer f.Close()

	var r io.Reader = f
	if filepath.Ext(path) == ".gz" {
		gr, err := gzip.NewReader(f)
		if err != nil {
			fmt.Println("Error opening gzip reader:", err)
			return false
		}
		defer gr.Close()
		r = gr
	}

	scanner := bufio.NewScanner(r)
	const maxCap = 4 * 1024 * 1024 // allow very long lines (4 MB)
	scanner.Buffer(make([]byte, 4*1024), maxCap)

	couponBytes := []byte(coupon)
	checkInterval := 100 // Check context every 100 lines
	lineCount := 0

	for scanner.Scan() {
		if lineCount%checkInterval == 0 {
			select {
			case <-ctx.Done(): // Another goroutine found matches
				return false
			default:
			}
		}
		lineCount++

		if bytes.EqualFold(scanner.Bytes(), couponBytes) {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return false
	}

	return false // End-of-file or no match
}
