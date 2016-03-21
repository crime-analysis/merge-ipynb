package merge

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sync"
)

func Merge(w io.Writer, r ...io.Reader) (int, error) {
	if len(r) == 0 {
		return 0, errors.New("no input readers provided")
	}

	if len(r) == 1 {
		i, err := io.Copy(w, r[0])
		return int(i), err
	}

	notebooks := make([]map[string]interface{}, len(r))
	wg := sync.WaitGroup{}
	wg.Add(len(r))
	ch := make(chan error, len(r))

	// unmarshal
	for i, reader := range r {
		reader := reader
		i := i

		go func() {
			defer wg.Done()
			ch <- json.NewDecoder(reader).Decode(&notebooks[i])
		}()
	}

	wg.Wait()
	close(ch)

	for err := range ch {
		if err != nil {
			return 0, err
		}
	}

	// the first notebook is used as the base
	fmt.Printf("%+v", notebooks[0])
	return 0, nil
}
