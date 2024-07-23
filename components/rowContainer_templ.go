// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func RowContainer() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"grid gap-x-2 gap-y-2 border-t border-gray-200 sm:mt-16 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(GetData()) > 0 {
			for _, item := range GetData() {
				templ_7745c5c3_Err = AmiiboBox(item).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

type Release struct {
	AU string `json:"au"`
	EU string `json:"eu"`
	JP string `json:"jp"`
	NA string `json:"na"`
}

type Amiibo struct {
	AmiiboSeries string  `json:"amiiboSeries"`
	Character    string  `json:"character"`
	GameSeries   string  `json:"gameSeries"`
	Head         string  `json:"head"`
	Image        string  `json:"image"`
	Name         string  `json:"name"`
	Release      Release `json:"release"`
	Tail         string  `json:"tail"`
	Type         string  `json:"type"`
}

type Response struct {
	Amiibo []Amiibo `json:"amiibo"`
}

func GetData() []Amiibo {
	url := "https://www.amiiboapi.com/api/amiibo/"

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("HTTP GET request failed: %s\n", err)
		return []Amiibo{}
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %s\n", err)
		return []Amiibo{}
	}

	var data Response
	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		fmt.Printf("Failed to unmarshal response body: %s\n", err)
		return []Amiibo{}
	}

	// Filter out only the amiibo figures, not cards
	var amiiboFigures []Amiibo
	for _, item := range data.Amiibo {
		if item.Type == "Figure" {
			amiiboFigures = append(amiiboFigures, item)
		}
	}

	return amiiboFigures
}
