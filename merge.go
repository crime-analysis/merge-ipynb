// Package merge provides a Merge function that merges iPython notebook cells

package merge

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sync"
)

// Merge merges the iPython notebooks by reading them from the provided Readers
// and writes the merged output JSON to the specified Writer
func Merge(w io.Writer, r ...io.Reader) error {
	if len(r) == 0 {
		return errors.New("no input readers provided")
	}

	if len(r) == 1 {
		_, err := io.Copy(w, r[0])
		return err
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
			return err
		}
	}

	baseCells, ok := notebooks[0]["cells"].([]interface{})
	if !ok {
		return errors.New("first notebook does not have expected format")
	}

	for i := 1; i < len(notebooks); i += 1 {
		cells, ok := notebooks[i]["cells"].([]interface{})
		if !ok {
			return fmt.Errorf("notebook #%d does not have expected format", i+1)
		}
		baseCells = append(baseCells, cells...)
	}

	notebooks[0]["cells"] = baseCells
	return json.NewEncoder(w).Encode(notebooks[0])
}
