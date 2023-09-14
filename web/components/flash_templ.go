// Code generated by templ@v0.2.334 DO NOT EDIT.

package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/ItsNotGoodName/smtpbridge/web/icons"
	"github.com/ItsNotGoodName/smtpbridge/web/routes"
)

type FlashType int

const (
	FlashTypeError FlashType = iota
	FlashTypeSuccess
)

type FlashProps struct {
	Component templ.Component
	Type      FlashType
}

func Flash(t FlashType, c templ.Component) templ.Component {
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
		_, err = templBuffer.WriteString("<div lass=\"flash\">")
		if err != nil {
			return err
		}
		switch t {
		case FlashTypeSuccess:
			_, err = templBuffer.WriteString("<div class=\"alert alert-success\">")
			if err != nil {
				return err
			}
			err = icons.CheckboxCircle("h-6 w-6").Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			err = c.Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div>")
			if err != nil {
				return err
			}
		default:
			_, err = templBuffer.WriteString("<div class=\"alert alert-error\">")
			if err != nil {
				return err
			}
			err = icons.CloseCircle("h-6 w-6").Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			err = c.Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</div>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func FlashMessage(message string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_2 := templ.GetChildren(ctx)
		if var_2 == nil {
			var_2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<span>")
		if err != nil {
			return err
		}
		var var_3 string = message
		_, err = templBuffer.WriteString(templ.EscapeString(var_3))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</span>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func FlashMessageLink(message string, route routes.Route) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_4 := templ.GetChildren(ctx)
		if var_4 == nil {
			var_4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<span>")
		if err != nil {
			return err
		}
		var var_5 string = message
		_, err = templBuffer.WriteString(templ.EscapeString(var_5))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</span><a class=\"btn btn-xs\" href=\"")
		if err != nil {
			return err
		}
		var var_6 templ.SafeURL = route.URL()
		_, err = templBuffer.WriteString(templ.EscapeString(string(var_6)))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		var_7 := `View`
		_, err = templBuffer.WriteString(var_7)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
