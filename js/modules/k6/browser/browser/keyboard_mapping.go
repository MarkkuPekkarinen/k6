package browser

import (
	"fmt"

	"github.com/grafana/sobek"

	"go.k6.io/k6/js/modules/k6/browser/common"
	"go.k6.io/k6/js/modules/k6/browser/k6ext"
)

func mapKeyboard(vu moduleVU, kb *common.Keyboard) mapping {
	return mapping{
		"down": func(key string) *sobek.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				return nil, kb.Down(key) //nolint:wrapcheck
			})
		},
		"up": func(key string) *sobek.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				return nil, kb.Up(key) //nolint:wrapcheck
			})
		},
		"press": func(key string, opts sobek.Value) *sobek.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				kbopts, err := exportTo[common.KeyboardOptions](vu.Runtime(), opts)
				if err != nil {
					return nil, fmt.Errorf("parsing keyboard options: %w", err)
				}
				return nil, kb.Press(key, kbopts)
			})
		},
		"type": func(text string, opts sobek.Value) *sobek.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				kbopts, err := exportTo[common.KeyboardOptions](vu.Runtime(), opts)
				if err != nil {
					return nil, fmt.Errorf("parsing keyboard options: %w", err)
				}
				return nil, kb.Type(text, kbopts)
			})
		},
		"insertText": func(text string) *sobek.Promise {
			return k6ext.Promise(vu.Context(), func() (any, error) {
				return nil, kb.InsertText(text) //nolint:wrapcheck
			})
		},
	}
}