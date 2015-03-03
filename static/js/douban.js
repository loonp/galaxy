(function() {
	var b = window.white_site_list || /^https?:\/\/([\w]+\.douban\.com|web[0-9]?\.qq\.com|hao\.qq\.com|(hao\.)*360\.cn|so\.com|www\.soso\.com)(\:[\d]+)?\//;
	if (self !== top && document.referrer.search(b) === -1) {
		top.location = self.location
	}
})();
Do = (typeof Do === "undefined") ?
function(b) {
	setTimeout(b, 0)
} : Do;
Douban = new Object();
Douban.errdetail = ["", "未知错误", "文件过大", "信息不全", "域名错误", "分类错误", "用户错误", "权限不足", "没有文件", "保存文件错误", "不支持的文件格式", "超时", "文件格式有误", "", "添加文件出错", "已经达到容量上限", "不存在的相册", "删除失败", "错误的MP3文件", "有禁用的内容,请修改重试"];
var trace = function(b) {
		if (!/^http:\/\/(www|movie|music\.|book|douban\.fm)/.test(location.href) && window.console && window.console.log) {
			console.log(b)
		}
	};

Douban.init_collect_btn = function(b) {
	$(b).click(function(h) {
		h.preventDefault();
		if ($("#hiddendialog").length) {
			show_dialog($("#hiddendialog").html());
			load_event_monitor($("#dialog"));
			return
		}
		show_dialog(null);
		var f = $(this).attr("name").split("-"),
			j = f[0],
			c = f[1],
			k = f[2],
			g = f[3],
			d = "/j/subject/" + c + "/interest?" + (k ? "interest=" + k : "") + (g ? "&rating=" + g : "") + (j == "cbtn" ? "&cmt=1" : "");
		$.getJSON(d, function(e) {
			if (!$("#dialog").length) {
				return
			}
			var o = $("<div></div>");
			o.get(0).innerHTML = e.html;
			var w = e.tags;
			var s = w.join(" ");
			$("input[name=tags]", o).val((s.length > 1) ? s + " " : s);
			var n = {};
			$.each(w, function(x, r) {
				n[r.toLowerCase()] = true
			});
			populate_tag_btns("我的标签:", $("#mytags", o), e.my_tags, n);
			populate_tag_btns("常用标签:", $("#populartags", o), e.popular_tags, n);
			if (j == "pbtn" || j == "cbtn") {
				$("form", o).data("reload", 1)
			}
			$("#dialog").html(o);
			$("#showtags").click(function() {
				if ($("#advtags").is(":hidden")) {
					$(this).html("缩起 ▲");
					$("#advtags").show();
					$("#foldcollect").val("U")
				} else {
					$(this).html($(this).attr("rel"));
					$("#advtags").hide();
					$("#foldcollect").val("F")
				}
				$(this).blur();
				refine_dialog()
			});				 
		});
		return false
	})
};
Douban.init_nine_collect_btn = function(b) {
	$(b).click(function() {
		var e = $(this).attr("name").split("-");
		var f = e[0],
			c = e[1],
			g = e[2];
		var d = "/j/subject/" + c + "/interest";
		$.getJSON(d, g && {
			interest: g
		}, function(l) {
			var j = $("<div></div>").html(l.html);
			var h = l.tags;
			var k = h.join(" ");
			$("input[name=tags]", j).val((k.length > 1) ? k + " " : k);
			var n = {};
			$.each(h, function(p, o) {
				n[o.toLowerCase()] = true
			});
			populate_tag_btns("我的标签(点击添加):", $("#mytags", j), l.my_tags, n);
			populate_tag_btns("常用的标签(点击添加):", $("#populartags", j), l.popular_tags, n);
			if (f == "pbtn") {
				$("form", j).data("reload", 1)
			}
			$("#collect_form_" + c).html("").append('<p class="ul"></p>').append(j);
			load_event_monitor($("#collect_form_" + c))
		});
		return false
	})
};

jQuery.fn.extend({
	pos: function() {
		var c = this[0];
		if (c.offsetParent) {
			for (var d = 0, b = 0; c.offsetParent; c = c.offsetParent) {
				d += c.offsetLeft;
				b += c.offsetTop
			}
			return {
				x: d,
				y: b
			}
		} else {
			return {
				x: c.x,
				y: c.y
			}
		}
	},
	chop: function(g, b) {
		var d = [],
			c = [];
		for (var e = 0, f = this.length; e < f; e++) {
			if (!b != !g(this[e], e)) {
				d.push(this[e])
			} else {
				c.push(this[e])
			}
		}
		return [d, c]
	},
	sum: function(c, e) {
		var b = this.length,
			d = zero = e ? "" : 0;
		while (b) {
			d += this[--b][c] + (b && e || zero)
		}
		return d
	},
	set_len_limit: function(c) {
		var d = this.find(":submit:first");
		var e = d.attr("value");
		var b = function() {
				if (this.value && this.value.length > c) {
					d.attr("disabled", 1).attr("value", "字数不能超过" + c + "字")
				} else {
					d.removeAttr("disabled").attr("value", e)
				}
			};
		$("textarea", this).focus(b).blur(b).keydown(b).keyup(b)
	},
	display_limit: function(b, g) {
		var c = this,
			e, d = function(h) {
				var f = c.val();
				if (f == e) {
					return
				}
				if (f.length >= b) {
					c.val(f.substring(0, b))
				}
				g.text(b - c.val().length);
				e = c.val()
			};
		this.keyup(d);
		d()
	},
	set_caret: function() {
		if (!$.browser.msie) {
			return
		}
		var b = function() {
				this.p = document.selection.createRange().duplicate()
			};
		this.click(b).select(b).keyup(b)
	},
	insert_caret: function(c) {
		var k = this[0];
		if (document.all && k.createTextRange && k.p) {
			var j = k.p;
			j.text = j.text.charAt(j.text.length - 1) == "" ? c + "" : c
		} else {
			if (k.setSelectionRange) {
				var f = k.selectionStart;
				var h = k.selectionEnd;
				var g = k.value.substring(0, f);
				var d = k.value.substring(h);
				k.value = g + c + d;
				k.focus();
				var b = c.length;
				k.setSelectionRange(f + b, f + b);
				k.blur()
			} else {
				k.value += c
			}
		}
	},
	get_sel: function() {
		var b = this[0];
		return document.all && b.createTextRange && b.p ? b.p.text : b.setSelectionRange ? b.value.substring(b.selectionStart, b.selectionEnd) : ""
	},
	blur_hide: function() {
		var c = this,
			b = function() {
				return false
			};
		c.mousedown(b);
		$(document.body).mousedown(function() {
			c.hide().unbind("mousedown", b);
			$(document.body).unbind("mousedown", arguments.callee)
		});
		return this
	},
	yellow_fade: function() {
		var b = 0,
			d = 1,
			c = this;

		function e() {
			c.css({
				backgroundColor: "rgb(100%,100%," + b + "%)"
			});
			b += d;
			d += 0.5;
			if (b <= 100) {
				setTimeout(e, 35)
			} else {
				c.css({
					backgroundColor: ""
				})
			}
		}
		e();
		return this
	},
	hover_fold: function(d) {
		var b = {
			folder: [1, 3],
			unfolder: [0, 2]
		},
			c = function(e, f) {
				return function() {
					$("img", e).attr("src", "/pics/arrow1_" + f + ".png")
				}
			};
		return this.hover(c(this, b[d][0]), c(this, b[d][1]))
	},
	multiselect: function(d) {
		var g = function() {
				return true
			},
			f = d.onselect || g,
			e = d.onremove || g,
			c = d.onchange || g,
			h = d.selclass || "sel",
			b = d.values || [];
		return this.click(function() {
			var k = /id(\d*)/.exec(this.className)[1],
				j = $.inArray(k, b);
			if (j != -1) {
				if (!e(this)) {
					return
				}
				b.splice(j, 1);
				$(this).removeClass(h)
			} else {
				if (!f(this)) {
					return
				}
				b.push(k);
				$(this).addClass(h)
			}
			c(b);
			return false
		})
	},
	initDataInput: function() {
		var b = $(this);
		if (!b.val() || b.val() === b.attr("title")) {
			b.addClass("color-lightgray");
			b.val(b.attr("title"))
		}
		b.focus(function() {
			b.removeClass("color-lightgray");
			if (b.val() === b.attr("title")) {
				b.val("")
			}
		}).blur(function() {
			if (!b.val()) {
				b.addClass("color-lightgray");
				b.val(b.attr("title"))
			}
		})
	},
	setItemList: function(p) {
		var c = {},
			g = "",
			n = '<img class="gray-loader" src="/pics/spinner.gif" />',
			e = "/pics/spinner.gif",
			j = ".input-create",
			l = {
				keyup: function(q) {
					var o = q.target.value.replace(/ /g, "");
					if (q.keyCode === 13) {
						p.create.callback(c, g, o, p.limit)
					}
				}
			},
			d = document.body,
			k = new Image(),
			f = {
				create: {
					title: "新分组",
					tips: "创建新分组"
				}
			},
			p = $.extend(f, p),
			b = '<span class="create-new">' + p.create.title + "</span>",
			h = '<input class="input-create" type="text" value="" title="' + p.create.tips + '" maxlength="' + p.create.maxLen + '" />';
		k.src = e;
		$(this).click(function(o) {
			o.stopPropagation();
			c = this;
			sglist.hide();
			g = $.isFunction(p.target) ? p.target(c) : p.target;
			sgarrow.removeClass(CSS_ARROW_SELECT);
			$(c).addClass(CSS_ARROW_SELECT);
			$(CSS_SET_GROUP_LIST, this).show();
			$(j).focus();
			if ($.browser.msie && $.browser.version !== "8.0") {
				sgarrow.css("z-index", "");
				$(this).css("z-index", 10)
			}
		});
		$(CSS_SET_GROUP_LIST).delegate("li:not('.last')", "click", function(w) {
			w.preventDefault();
			var v = w.target,
				q = this,
				u = v.type === "checkbox" ? true : false,
				o = $(this).children("input"),
				s = $(this).children("input").val(),
				r = (u && o.attr("checked") || !u && !o.attr("checked")) ? "addtotag" : "removefromtag";
			if (!$(CSS_LOADER, this).length) {
				o.hide().after(n)
			}
			p.callback(q, r, u, g, s)
		});
		$(d).click(function(o) {
			$(CSS_SET_GROUP_LIST, this).hide();
			$(c).removeClass(CSS_ARROW_SELECT);
			if (newGroupNum && newGroupNum < p.limit) {
				$(j).replaceWith(b)
			}
		});
		$(CSS_SET_GROUP_LIST).delegate(".create-new", "click", function() {
			$(this).replaceWith(h);
			$(j).focus()
		});
		$(CSS_SET_GROUP_LIST).delegate(j, "keyup", function(o) {
			if ($.isFunction(l[o.type])) {
				l[o.type].call(this, o)
			}
		})
	}
});
 
 
 