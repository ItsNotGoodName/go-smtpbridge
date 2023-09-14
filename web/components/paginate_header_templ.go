// Code generated by templ@v0.2.334 DO NOT EDIT.

package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"net/url"
	"strconv"

	"github.com/ItsNotGoodName/smtpbridge/pkg/pagination"
	"github.com/ItsNotGoodName/smtpbridge/web/helpers"
	"github.com/ItsNotGoodName/smtpbridge/web/routes"
)

type PaginateHeaderProps struct {
	Route      routes.Route
	Query      url.Values
	PageResult pagination.PageResult
	Ascending  bool
}

func PaginateHeader(props PaginateHeaderProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<div class=\"flex flex-col items-center justify-between gap-4 sm:flex-row\"><div class=\"mr-auto flex flex-row items-center gap-2\"><div class=\"dropdown\"><label tabindex=\"0\" class=\"btn btn-sm btn-outline\">")
		if err != nil {
			return err
		}
		var var_2 string = strconv.Itoa(props.PageResult.PerPage)
		_, err = templBuffer.WriteString(templ.EscapeString(var_2))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label><ul tabindex=\"0\" class=\"menu dropdown-content bg-base-100 rounded-box z-50 w-52 p-2 shadow-lg\">")
		if err != nil {
			return err
		}
		for _, option := range limits {
			_, err = templBuffer.WriteString("<li><a href=\"")
			if err != nil {
				return err
			}
			var var_3 templ.SafeURL = props.Route.URLQuery(helpers.Query(props.Query, "perPage", option))
			_, err = templBuffer.WriteString(templ.EscapeString(string(var_3)))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\">")
			if err != nil {
				return err
			}
			var var_4 string = strconv.Itoa(option)
			_, err = templBuffer.WriteString(templ.EscapeString(var_4))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</a></li>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</ul></div>")
		if err != nil {
			return err
		}
		if props.Ascending {
			_, err = templBuffer.WriteString("<a class=\"btn btn-sm btn-outline\" href=\"")
			if err != nil {
				return err
			}
			var var_5 templ.SafeURL = props.Route.URLQuery(helpers.Query(props.Query, "ascending", ""))
			_, err = templBuffer.WriteString(templ.EscapeString(string(var_5)))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\">")
			if err != nil {
				return err
			}
			var_6 := `ASC`
			_, err = templBuffer.WriteString(var_6)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</a>")
			if err != nil {
				return err
			}
		} else {
			_, err = templBuffer.WriteString("<a class=\"btn btn-sm btn-outline\" href=\"")
			if err != nil {
				return err
			}
			var var_7 templ.SafeURL = props.Route.URLQuery(helpers.Query(props.Query, "ascending", "1"))
			_, err = templBuffer.WriteString(templ.EscapeString(string(var_7)))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\">")
			if err != nil {
				return err
			}
			var_8 := `DESC`
			_, err = templBuffer.WriteString(var_8)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</a>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</div><div class=\"ml-auto flex flex-row items-center justify-between gap-4\"><div class=\"join\"><a class=\"join-item btn btn-sm btn-outline\"")
		if err != nil {
			return err
		}
		if !props.PageResult.HasPrevious() {
			_, err = templBuffer.WriteString(" disabled")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString(" href=\"")
		if err != nil {
			return err
		}
		var var_9 templ.SafeURL = props.Route.URLQuery((helpers.Query(props.Query, "page", props.PageResult.Previous())))
		_, err = templBuffer.WriteString(templ.EscapeString(string(var_9)))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		var_10 := `Previous`
		_, err = templBuffer.WriteString(var_10)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a><button class=\"join-item btn btn-sm btn-outline\">")
		if err != nil {
			return err
		}
		var var_11 string = strconv.Itoa(props.PageResult.Page)
		_, err = templBuffer.WriteString(templ.EscapeString(var_11))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(" ")
		if err != nil {
			return err
		}
		var_12 := `/ `
		_, err = templBuffer.WriteString(var_12)
		if err != nil {
			return err
		}
		var var_13 string = strconv.Itoa(props.PageResult.TotalPages)
		_, err = templBuffer.WriteString(templ.EscapeString(var_13))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</button><a class=\"join-item btn btn-sm btn-outline\"")
		if err != nil {
			return err
		}
		if !props.PageResult.HasNext() {
			_, err = templBuffer.WriteString(" disabled")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString(" href=\"")
		if err != nil {
			return err
		}
		var var_14 templ.SafeURL = props.Route.URLQuery(helpers.Query(props.Query, "page", props.PageResult.Next()))
		_, err = templBuffer.WriteString(templ.EscapeString(string(var_14)))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		var_15 := `Next`
		_, err = templBuffer.WriteString(var_15)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></div></div></div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
