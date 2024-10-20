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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<nav><div class=\"navbar bg-base-100\"><div class=\"flex-1\"><a class=\"btn btn-ghost text-xl\">Placeholder</a></div><div class=\"flex-none gap-2\"><div class=\"form-control\"><input type=\"text\" placeholder=\"Search\" class=\"input input-bordered w-24 md:w-auto\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if ctx.Value("isLoggedIn") == false {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a href=\"/register\" hx-boost=\"true\" class=\"btn btn-ghost\">Register</a> <a href=\"/login\" hx-boost=\"true\" class=\"btn btn-ghost\">Login</a> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if ctx.Value("isLoggedIn") == true {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"dropdown dropdown-end\"><div tabindex=\"0\" role=\"button\" class=\"btn btn-ghost btn-circle avatar\"><div class=\"w-10 rounded-full\"><img alt=\"Tailwind CSS Navbar component\" src=\"https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp\"></div></div><ul tabindex=\"0\" class=\"mt-3 z-[1] p-2 shadow menu menu-sm dropdown-content bg-base-100 rounded-box w-52\"><li><div>Profile</div></li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if ctx.Value("role") == "admin" {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><a href=\"/user/admin/manage-users\" hx-boost=\"true\">Manage Users</a></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form action=\"/logout\" method=\"post\" hx-boost=\"true\" class=\"justify-between w-full text-left\"><li><button type=\"submit\">Logout</button></li></form></ul></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
