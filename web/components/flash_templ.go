// Code generated by templ@v0.2.334 DO NOT EDIT.

package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
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
			_, err = templBuffer.WriteString("<div class=\"alert alert-success\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-6 w-6 shrink-0 stroke-current\" fill=\"none\" viewBox=\"0 0 24 24\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z\"></path></svg>")
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
			_, err = templBuffer.WriteString("<div class=\"alert alert-error\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-6 w-6 shrink-0 stroke-current\" fill=\"none\" viewBox=\"0 0 24 24\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z\"></path></svg>")
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