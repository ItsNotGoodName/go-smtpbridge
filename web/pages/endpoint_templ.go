// Code generated by templ@v0.2.334 DO NOT EDIT.

package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/ItsNotGoodName/smtpbridge/internal/models"
	c "github.com/ItsNotGoodName/smtpbridge/web/components"
	"github.com/ItsNotGoodName/smtpbridge/web/meta"
	"github.com/ItsNotGoodName/smtpbridge/web/routes"
)

type endpointListViewProps struct {
	Endpoints []models.Endpoint
}

func endpointListView(m meta.Meta, props endpointListViewProps) templ.Component {
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
		var_2 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			_, err = templBuffer.WriteString("<div class=\"border-base-200 breadcrumbs border-b p-4 text-xl font-bold\"><ul><li>")
			if err != nil {
				return err
			}
			var_3 := `Endpoints`
			_, err = templBuffer.WriteString(var_3)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</li></ul></div> <div class=\"flex flex-wrap gap-4 p-4\">")
			if err != nil {
				return err
			}
			for _, end := range props.Endpoints {
				_, err = templBuffer.WriteString("<div class=\"w-full sm:w-96\"><div class=\"card card-compact bg-base-100 border-base-200 border\"><div class=\"card-body\"><h2 class=\"card-title\">")
				if err != nil {
					return err
				}
				var var_4 string = end.Name
				_, err = templBuffer.WriteString(templ.EscapeString(var_4))
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</h2><div class=\"flex flex-col gap-2\"><div class=\"flex items-center justify-between\"><div>")
				if err != nil {
					return err
				}
				var_5 := `Kind`
				_, err = templBuffer.WriteString(var_5)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</div><div>")
				if err != nil {
					return err
				}
				var var_6 string = end.Kind
				_, err = templBuffer.WriteString(templ.EscapeString(var_6))
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</div></div><div class=\"flex items-center justify-between\"><div>")
				if err != nil {
					return err
				}
				var_7 := `Internal`
				_, err = templBuffer.WriteString(var_7)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</div><div><input type=\"checkbox\" class=\"toggle\"")
				if err != nil {
					return err
				}
				if end.Internal {
					_, err = templBuffer.WriteString(" checked")
					if err != nil {
						return err
					}
				}
				_, err = templBuffer.WriteString(" disabled></div></div><div class=\"flex items-center justify-between\"><div>")
				if err != nil {
					return err
				}
				var_8 := `Text Disable`
				_, err = templBuffer.WriteString(var_8)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</div><div><input type=\"checkbox\" class=\"toggle\"")
				if err != nil {
					return err
				}
				if end.TextDisable {
					_, err = templBuffer.WriteString(" checked")
					if err != nil {
						return err
					}
				}
				if end.Internal {
					_, err = templBuffer.WriteString(" disabled")
					if err != nil {
						return err
					}
				}
				_, err = templBuffer.WriteString("></div></div><div class=\"flex items-center justify-between\"><div>")
				if err != nil {
					return err
				}
				var_9 := `Attachment Disable`
				_, err = templBuffer.WriteString(var_9)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</div><div><input type=\"checkbox\" class=\"toggle\"")
				if err != nil {
					return err
				}
				if end.AttachmentDisable {
					_, err = templBuffer.WriteString(" checked")
					if err != nil {
						return err
					}
				}
				if end.Internal {
					_, err = templBuffer.WriteString(" disabled")
					if err != nil {
						return err
					}
				}
				_, err = templBuffer.WriteString("></div></div><div class=\"form-control\"><label class=\"label\"><span class=\"label-text\">")
				if err != nil {
					return err
				}
				var_10 := `Title Template`
				_, err = templBuffer.WriteString(var_10)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</span></label><textarea")
				if err != nil {
					return err
				}
				if end.Internal {
					_, err = templBuffer.WriteString(" disabled")
					if err != nil {
						return err
					}
				}
				_, err = templBuffer.WriteString(" class=\"textarea textarea-bordered h-24\">")
				if err != nil {
					return err
				}
				var var_11 string = end.TitleTemplate
				_, err = templBuffer.WriteString(templ.EscapeString(var_11))
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</textarea></div><div class=\"form-control\"><label class=\"label\"><span class=\"label-text\">")
				if err != nil {
					return err
				}
				var_12 := `Body Template`
				_, err = templBuffer.WriteString(var_12)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</span></label><textarea")
				if err != nil {
					return err
				}
				if end.Internal {
					_, err = templBuffer.WriteString(" disabled")
					if err != nil {
						return err
					}
				}
				_, err = templBuffer.WriteString(" class=\"textarea textarea-bordered h-24\">")
				if err != nil {
					return err
				}
				var var_13 string = end.BodyTemplate
				_, err = templBuffer.WriteString(templ.EscapeString(var_13))
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</textarea></div></div><div class=\"flex flex-col gap-2\" data-loading-states><div class=\"card-actions justify-end\"><button class=\"btn btn-success\" hx-post=\"")
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(templ.EscapeString(routes.EndpointTest(end.ID).URLString()))
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("\" hx-target=\"next div\" data-loading-disable><span data-loading-class=\"loading loading-spinner loading-sm\">")
				if err != nil {
					return err
				}
				var_14 := `Test`
				_, err = templBuffer.WriteString(var_14)
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</span></button></div><div></div></div></div></div></div>")
				if err != nil {
					return err
				}
			}
			_, err = templBuffer.WriteString("</div>")
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = c.LayoutDefault(m).Render(templ.WithChildren(ctx, var_2), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
