package streams

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/TykTechnologies/tyk/header"
	"io"
	"net/http"
)

func (s *Stream) DoRequest(done <-chan int, ch <-chan http.Request) {
	go func() {
		for {
			select {
			case <-done:
				return
			case req := <-ch:
				go func(r *http.Request) {
					r = r.WithContext(context.Background())
					r.Header.Set(header.ContentType, header.ApplicationJSON)

					// todo:лучше создать свой клиент
					resp, err := http.DefaultClient.Do(r)
					defer func(Body io.ReadCloser) {
						err := Body.Close()
						if err != nil {
							s.logger.Error("response body close error: ", err)
						}
					}(resp.Body)
					if err != nil {
						s.logger.Error("http request error: ", err)
						return
					}

					body, _ := io.ReadAll(resp.Body)

					//res, err := helpers.Fetch(r)
					//body, _ := io.ReadAll(res.Body)

					var prettyJSON bytes.Buffer
					err = json.Indent(&prettyJSON, body, "", "\t")
					if err != nil {
						s.logger.Debug("JSON parse error: ", err)
						return
					}
					s.logger.Debug("response Body:", string(prettyJSON.Bytes()))
				}(&req)
			}
		}
	}()
}
