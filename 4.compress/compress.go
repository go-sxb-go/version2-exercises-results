package compress

import "compress/gzip"
import "compress/lzw"
import "io"
import "io/ioutil"

// CompressTo takes a file and a compression mode as input
// And return the path to a temp file containing the compressed
// version of the input
// mode can be 'lzw', 'gzip', or 'none' which will just copy the input in the file
func CompressTo(input io.Reader, mode string) (string, error) {
	file, err := ioutil.TempFile("", "go-compress")
	if err != nil {
		return "", err
	}

	switch mode {
	case "lzw":
		lzwCompress(input, file)
	case "gzip":
		gzipCompress(input, file)
	case "none":
		noneCompress(input, file)
	}

	file.Close()
	return file.Name(), nil
}

func lzwCompress(input io.Reader, out io.Writer) error {
	writer := lzw.NewWriter(out, lzw.MSB, 8)
	_, err := io.Copy(writer, input)
	if err != nil {
		return err
	}
	writer.Close()
	return nil
}

func gzipCompress(input io.Reader, out io.Writer) error {
	writer := gzip.NewWriter(out)
	_, err := io.Copy(writer, input)
	if err != nil {
		return err
	}
	writer.Close()
	return nil
}

func noneCompress(input io.Reader, out io.Writer) error {
	_, err := io.Copy(out, input)
	return err
}
