package client // import "github.com/fdurand/moby/client"

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/fdurand/moby/api/types"
	"github.com/fdurand/moby/api/types/filters"
	"github.com/fdurand/moby/api/types/swarm"
	"github.com/fdurand/moby/errdefs"
)

func TestNodeListError(t *testing.T) {
	client := &Client{
		client: newMockClient(errorMock(http.StatusInternalServerError, "Server error")),
	}

	_, err := client.NodeList(context.Background(), types.NodeListOptions{})
	if !errdefs.IsSystem(err) {
		t.Fatalf("expected a Server Error, got %[1]T: %[1]v", err)
	}
}

func TestNodeList(t *testing.T) {
	expectedURL := "/nodes"

	filters := filters.NewArgs()
	filters.Add("label", "label1")
	filters.Add("label", "label2")

	listCases := []struct {
		options             types.NodeListOptions
		expectedQueryParams map[string]string
	}{
		{
			options: types.NodeListOptions{},
			expectedQueryParams: map[string]string{
				"filters": "",
			},
		},
		{
			options: types.NodeListOptions{
				Filters: filters,
			},
			expectedQueryParams: map[string]string{
				"filters": `{"label":{"label1":true,"label2":true}}`,
			},
		},
	}
	for _, listCase := range listCases {
		client := &Client{
			client: newMockClient(func(req *http.Request) (*http.Response, error) {
				if !strings.HasPrefix(req.URL.Path, expectedURL) {
					return nil, fmt.Errorf("Expected URL '%s', got '%s'", expectedURL, req.URL)
				}
				query := req.URL.Query()
				for key, expected := range listCase.expectedQueryParams {
					actual := query.Get(key)
					if actual != expected {
						return nil, fmt.Errorf("%s not set in URL query properly. Expected '%s', got %s", key, expected, actual)
					}
				}
				content, err := json.Marshal([]swarm.Node{
					{
						ID: "node_id1",
					},
					{
						ID: "node_id2",
					},
				})
				if err != nil {
					return nil, err
				}
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewReader(content)),
				}, nil
			}),
		}

		nodes, err := client.NodeList(context.Background(), listCase.options)
		if err != nil {
			t.Fatal(err)
		}
		if len(nodes) != 2 {
			t.Fatalf("expected 2 nodes, got %v", nodes)
		}
	}
}
