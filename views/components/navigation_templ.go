// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Navigation() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<nav><div class=\"navbar bg-base-100\"><div class=\"flex-1\"><a class=\"btn btn-ghost text-xl\">Placeholder</a></div><div class=\"flex-none gap-2\"><div class=\"form-control\"><input type=\"text\" placeholder=\"Search\" class=\"input input-bordered w-24 md:w-auto\"></div><div class=\"dropdown dropdown-end\"><div tabindex=\"0\" role=\"button\" class=\"btn btn-ghost btn-circle avatar\"><div class=\"w-10 rounded-full\"><img alt=\"Tailwind CSS Navbar component\" src=\"https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg\"></div></div><ul tabindex=\"0\" class=\"mt-3 z-[1] p-2 shadow menu menu-sm dropdown-content bg-base-100 rounded-box w-52\"><li><a class=\"justify-between\">Profile <span class=\"badge\">New</span></a></li><li><a>Settings</a></li><li><a>Logout</a></li></ul></div></div></div></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
