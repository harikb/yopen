package yopen

import (
	"compress/gzip"
	"io"
	"os"
	"path"
	"strings"
)

// Reader is returned by NewReader
type Reader struct {
	io.Reader              // exposed to external world. See http://bit.ly/1KPttpo
	fp        io.Reader    // used to store original source file or stream
	gz        *gzip.Reader // used to store gz.Reader reference
}

// Close the associated files.
func (r *Reader) Close() (err error) {
	if r.gz != nil {
		err = r.gz.Close()
		if err != nil {
			return
		}
	}
	// fp might as well be *os.File and avoid the cast below
	// however, I am keeping it open as brentp/xopen in case we
	// take non-file input.
	if c, ok := r.fp.(io.ReadCloser); ok {
		err = c.Close()
		if err != nil {
			return
		}
	}
	return
}

// Writer is returned by NewWriter
type Writer struct {
	io.Writer              // exposed to external world. See http://bit.ly/1KPttpo
	fp        io.Writer    // used to store original source file or stream
	gz        *gzip.Writer // used to store gz.Reader reference
	finalName string
}

// Close the associated files.
// If io.Writer is a ReadCloser, Close() the file.
// If io.Writer is a *os.File and there is a finalName, rename to that name.
func (w *Writer) Close() (err error) {

	if w.gz != nil {
		err = w.gz.Flush()
		if err != nil {
			return
		}
		err = w.gz.Close()
		if err != nil {
			return
		}
	}
	if c, ok := w.fp.(io.ReadCloser); ok {
		tempFN := ""
		if f, ok := w.fp.(*os.File); ok {
			tempFN = f.Name()
		}
		err = c.Close()
		if err != nil {
			return
		}
		if w.finalName != "" && tempFN != "" {
			err = os.Rename(tempFN, w.finalName)
			if err != nil {
				return
			}
		}
	}
	return
}

// NewReader open the file specified by the name or a .gz version
// of it if opening by the original file failed (this is to handle
// a concurrent process gzip-ing files).
func NewReader(filename string) (yr *Reader, err error) {

	var in *os.File
	gzFallback := false
	var gz *gzip.Reader

	if in, err = os.Open(filename); err != nil {
		if os.IsNotExist(err) && !strings.HasSuffix(filename, ".gz") {
			// try to open the .gz version
			in, err = os.Open(filename + ".gz")
			gzFallback = true
		}
		if err != nil {
			return
		}
	}

	if gzFallback || strings.HasSuffix(filename, ".gz") {
		// replace reader with a gzip reader
		gz, err = gzip.NewReader(in)
	}
	if gz != nil {
		yr = &Reader{gz, in, gz}
	} else {
		yr = &Reader{in, in, gz}
	}
	return
}

// NewWriter creates the file specified by the name or a .gz version
// of it depending on the name. It creates the file with dot-prefix
// and rename to original name at Close()
func NewWriter(filename string) (yw *Writer, err error) {

	var out *os.File
	var gz *gzip.Writer

	dir := path.Dir(filename)
	fn := path.Base(filename)
	tempFN := path.Join(dir, "."+fn)

	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return
	}
	out, err = os.Create(tempFN)
	if err != nil {
		return
	}
	if strings.HasSuffix(filename, ".gz") {
		gz = gzip.NewWriter(out)
	}
	if gz != nil {
		yw = &Writer{gz, out, gz, filename}
	} else {
		yw = &Writer{out, out, gz, filename}
	}
	return
}
