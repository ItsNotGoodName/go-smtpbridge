// Code generated by templ@v0.2.334 DO NOT EDIT.

package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"strconv"

	"github.com/ItsNotGoodName/smtpbridge/internal/models"
	"github.com/ItsNotGoodName/smtpbridge/web/icons"
	"github.com/ItsNotGoodName/smtpbridge/web/routes"
)

type RuleFormProps struct {
	Flash  templ.Component
	Create bool
	Data   RuleFormData
}

type RuleFormData struct {
	ID                  int64
	Internal            bool
	Name                string
	NameError           string
	Expression          string
	ExpressionErr       error
	Endpoints           []models.Endpoint
	EndpointsSelections []bool
}

func RuleForm(props RuleFormProps) templ.Component {
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
		_, err = templBuffer.WriteString("<form hx-post=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(props.Route().URLString()))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" class=\"flex flex-col gap-4\" data-loading-states>")
		if err != nil {
			return err
		}
		if props.Data.Internal {
			_, err = templBuffer.WriteString("<div class=\"alert alert-warning\">")
			if err != nil {
				return err
			}
			err = icons.Alert("h-6 w-6").Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("<span>")
			if err != nil {
				return err
			}
			var_2 := `Rule cannot be edited because it is internal.`
			_, err = templBuffer.WriteString(var_2)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</span></div>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("<div class=\"form-control\"><label class=\"label\"><span class=\"label-text\">")
		if err != nil {
			return err
		}
		var_3 := `Name`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</span></label><input")
		if err != nil {
			return err
		}
		if props.Data.Internal {
			_, err = templBuffer.WriteString(" disabled")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString(" name=\"name\" type=\"text\" placeholder=\"Name\" class=\"input input-bordered\" value=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(props.Data.Name))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		if props.Data.NameError != "" {
			_, err = templBuffer.WriteString("<label class=\"label\"><span class=\"label-text-alt text-error\">")
			if err != nil {
				return err
			}
			var var_4 string = props.Data.NameError
			_, err = templBuffer.WriteString(templ.EscapeString(var_4))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</span></label>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</div><div class=\"form-control\" data-loading-states data-loading-path=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(routes.RuleExpressionCheck().URLString()))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><label class=\"label\"><span class=\"label-text\">")
		if err != nil {
			return err
		}
		var_5 := `Expression`
		_, err = templBuffer.WriteString(var_5)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</span><span class=\"label-text-alt\"><span data-loading-class=\"loading loading-spinner loading-xs\"></span></span></label><textarea")
		if err != nil {
			return err
		}
		if props.Data.Internal {
			_, err = templBuffer.WriteString(" disabled")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString(" name=\"expression\" hx-post=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(routes.RuleExpressionCheck().URLString()))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" hx-trigger=\"keyup changed delay:100ms\" hx-target=\"next label\" class=\"textarea textarea-bordered h-24\" placeholder=\"Expression\">")
		if err != nil {
			return err
		}
		var var_6 string = props.Data.Expression
		_, err = templBuffer.WriteString(templ.EscapeString(var_6))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</textarea><label class=\"label\">")
		if err != nil {
			return err
		}
		err = RuleFormExpressionLabel(RuleFormExpressionLabelProps{
			Err: props.Data.ExpressionErr,
		}).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</label></div><fieldset><legend>")
		if err != nil {
			return err
		}
		var_7 := `Endpoints`
		_, err = templBuffer.WriteString(var_7)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</legend>")
		if err != nil {
			return err
		}
		for i, end := range props.Data.Endpoints {
			_, err = templBuffer.WriteString("<div class=\"form-control\"><label class=\"label cursor-pointer\"><span class=\"label-text\">")
			if err != nil {
				return err
			}
			var var_8 string = end.Name
			_, err = templBuffer.WriteString(templ.EscapeString(var_8))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</span><input")
			if err != nil {
				return err
			}
			if props.Data.Internal {
				_, err = templBuffer.WriteString(" disabled")
				if err != nil {
					return err
				}
			}
			_, err = templBuffer.WriteString(" type=\"checkbox\" class=\"toggle\" name=\"endpoints\" value=\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(templ.EscapeString(strconv.FormatInt(end.ID, 10)))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\"")
			if err != nil {
				return err
			}
			if props.Data.EndpointsSelections[i] {
				_, err = templBuffer.WriteString(" checked")
				if err != nil {
					return err
				}
			}
			_, err = templBuffer.WriteString("></label></div>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</fieldset><button")
		if err != nil {
			return err
		}
		if props.Data.Internal {
			_, err = templBuffer.WriteString(" disabled")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString(" type=\"submit\" class=\"btn btn-primary btn-block\" data-loading-disable><span data-loading-class=\"loading loading-spinner loading-xs\">")
		if err != nil {
			return err
		}
		if props.Create {
			var_9 := `Create Rule`
			_, err = templBuffer.WriteString(var_9)
			if err != nil {
				return err
			}
		} else {
			var_10 := `Update Rule`
			_, err = templBuffer.WriteString(var_10)
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</span></button>")
		if err != nil {
			return err
		}
		if props.Flash != nil {
			err = props.Flash.Render(ctx, templBuffer)
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</form>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

type RuleFormExpressionLabelProps struct {
	Err error
}

func RuleFormExpressionLabel(props RuleFormExpressionLabelProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_11 := templ.GetChildren(ctx)
		if var_11 == nil {
			var_11 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if props.Err != nil {
			_, err = templBuffer.WriteString("<span class=\"label-text-alt text-error\">")
			if err != nil {
				return err
			}
			var var_12 string = props.Err.Error()
			_, err = templBuffer.WriteString(templ.EscapeString(var_12))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</span>")
			if err != nil {
				return err
			}
		} else {
			_, err = templBuffer.WriteString("<span class=\"label-text-alt text-success\">")
			if err != nil {
				return err
			}
			var_13 := `Valid!`
			_, err = templBuffer.WriteString(var_13)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</span>")
			if err != nil {
				return err
			}
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
