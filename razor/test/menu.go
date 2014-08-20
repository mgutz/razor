// DO NOT EDIT! Auto-generated by github.com/mgutz/razor

package cases

import (
	"github.com/mgutz/razor/razor"
)

// Menu is generated
func Menu() *razor.SafeBuffer {
	_buffer := razor.NewSafeBuffer()
	locals := razor.Locals
	if locals != nil {
		// avoids not declared error if locals is not used
	}
	_buffer.WriteString("<ul class=\"nav nav-sidebar\">\n	<li role=\"presentation\" class=\"dropdown-header\">用户管理</li>\n	<li><a href=\"/admin/user\">查看用户</a></li>\n	<li><a href=\"/admin/user/create\">添加用户</a></li>\n	<li role=\"presentation\" class=\"divider\"></li>\n	<li role=\"presentation\" class=\"dropdown-header\">公文管理</li>\n	<li><a href=\"#\">收文管理</a></li>\n	<li><a href=\"#\">收文登记</a></li>\n	<li><a href=\"#\">发送公文</a></li>\n	<li><a href=\"#\">发文管理</a></li>\n	<li><a href=\"#\">发文登记</a></li>\n</ul>\n<ul class=\"nav nav-sidebar\">\n	<li><a href=\"\">领导审批</a></li>\n	<li><a href=\"\">流程监控</a></li>\n	<li role=\"presentation\" class=\"divider\"></li>\n	<li role=\"presentation\" class=\"dropdown-header\">其它</li>\n	<li><a href=\"\">添加日程</a></li>\n	<li><a href=\"\">公共通讯录</a></li>\n	<li><a href=\"\">添加联系人</a></li>\n	<li><a href=\"\">投票</a></li>\n</ul>")

	return _buffer
}
