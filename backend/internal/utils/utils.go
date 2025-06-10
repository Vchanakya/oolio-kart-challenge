package utils

import (
	"bufio"
	"compress/gzip"
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
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
	filePath := ("coupons/" + path)
	f, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer f.Close()

	var r io.Reader = f
	if filepath.Ext(path) == ".gz" {
		gr, err := gzip.NewReader(f)
		if err != nil {
			return false
		}
		defer gr.Close()
		r = gr
	}

	scanner := bufio.NewScanner(r)
	const maxCap = 4 * 1024 * 1024 // allow very long lines (4 MB)
	scanner.Buffer(make([]byte, 64*1024), maxCap)

	capitals := strings.ToUpper(coupon)
	for scanner.Scan() {
		select {
		case <-ctx.Done(): // another goroutine already found two matches
			return false
		default:
		}
		if strings.ToUpper(scanner.Text()) == capitals {
			return true
		}
	}
	return false // end-of-file or scanner error
}
