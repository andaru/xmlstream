package xmlstream

import (
	"context"
	"encoding/xml"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXMLSM(t *testing.T) {
	for _, tc := range []struct {
		name      string
		input     string
		wantToken []xml.Token
	}{
		{
			name:  "nested tokens",
			input: "<foo><bar>baz</bar></foo>",
			wantToken: []xml.Token{
				xml.StartElement{Name: xml.Name{Local: "foo"}, Attr: []xml.Attr{}},
				xml.StartElement{Name: xml.Name{Local: "bar"}, Attr: []xml.Attr{}},
				xml.CharData("baz"),
				xml.EndElement{Name: xml.Name{Local: "bar"}},
				xml.EndElement{Name: xml.Name{Local: "foo"}},
			},
		},
		{
			name:  "space CDATA",
			input: "  ",
			wantToken: []xml.Token{
				xml.CharData("  "),
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			check := assert.New(t)
			d := xml.NewDecoder(strings.NewReader(tc.input))
			var gotToken []xml.Token
			parser := func(ctx context.Context, sm *StateMachine) StateFn {
				var decoderErr error
				var token xml.Token
				for {
					if token, decoderErr = d.Token(); decoderErr != nil {
						break
					}
					gotToken = append(gotToken, xml.CopyToken(token))
				}
				return BailWithError(decoderErr)
			}
			sm := &StateMachine{Begin: parser}
			sm.End = IgnoreEOF

			check.NoError(sm.Run(context.Background()))
			if check.Len(gotToken, len(tc.wantToken)) {
				for i := range tc.wantToken {
					check.Equal(tc.wantToken[i], gotToken[i])
				}
			}
		})
	}
}
