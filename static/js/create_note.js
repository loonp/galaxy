function openLinkDialog(a) {
}(function(e) {
	var f = {
		delimeter: /[,，]/,
		count_limit: 5,
		max_letters: 40,
		warn: {
			count_limit: "最多只能添加5个标签哦！",
			max_letters: "单个标签超过字数限制"
		},
		tag_tmpl: '<span class="tag"><i>{{tag}}</i><a href="#">X</a></span>'
	};
	var a = {
		COMMA: 44,
		COMMA_ZH: 65292,
		BACKSPACE: 8,
		SPACE: 32,
		END: 35,
		HOME: 36,
		LEFT: 37,
		RIGHT: 39,
		A: 65
	};
	var c = {
		getCharcode: function(h) {
			var g = h.charCode;
			return (typeof g === "number" && g !== 0) ? g : h.keyCode
		},
		getStrLength: function(h) {
			var g = 0;
			e(h.split("")).each(function(j, k) {
				g += (k.charCodeAt(0) > 0 && k.charCodeAt(0) < 128) ? 1 : 2
			});
			return g
		},
		getCaretPos: function(g) {
			return "selectionStart" in g ? g.selectionStart : Math.abs(document.selection.createRange().moveStart("character", -g.value.length))
		},
		setCaretPos: function(g, h) {
			if (g.setSelectionRange) {
				g.setSelectionRange(h, h)
			} else {
				if (g.createTextRange) {
					g.createTextRange().move("character", h)
				}
			}
			g.focus()
		},
		getCaretText: function(g, h) {
			var i = c.getCaretPos(g);
			if (!i) {
				return
			}
			return g.value.substring(i - h, i)
		}
	};
	var d = function(g) {
			this.text = g;
			this.node = null
		};
	d.prototype = {
		constructor: d,
		build: function(h, i, g) {
			this.node = [e.fn.insertAfter, e.fn.insertBefore][i === "after" ? 0 : 1].call(e(h.replace("{{tag}}", this.text)), g)
		},
		editable: function(g) {
			g.tag_writer.addClass("small-tags-writer").width(this.node.outerWidth()).insertBefore(this.node);
			g.tag_input.val(this.text).focus()
		}
	};
	var b = function(h, g) {
			this.taglist = [];
			this.dom = {
				tag_editor: e(".tags-editor", h),
				tag_writer: e(".tags-writer", h),
				tag_input: e('input[type="text"]', h)
			};
			this.options = g;
			this._events = e({});
			this.init(this.dom)
		};
	b.prototype = {
		constructor: b,
		insert: function(i, j) {
			i = e.trim(i);
			var h = this.indexOf(i);
			if (i.length === 0) {
				return false
			}
			if (h > -1) {
				this.fire("error", {
					type: "exist",
					msg: h
				});
				return false
			}
			if (this.taglist.length === this.options.count_limit) {
				this.fire("error", {
					msg: this.options.warn.count_limit
				});
				return false
			}
			if (c.getStrLength(i) > this.options.max_letters) {
				this.fire("error", {
					msg: this.options.warn.max_letters
				});
				return false
			}
			var g = new d(i);
			g.build(this.options.tag_tmpl, j, this.dom.tag_writer);
			this.taglist.push(g);
			this.fire("change");
			return true
		},
		indexOf: function(g) {
			var h = -1;
			e(this.taglist).each(function(k, j) {
				if (j.text.toLowerCase() === g.toLowerCase()) {
					h = k;
					return false
				}
			});
			return h
		},
		getTag: function(g) {
			var h = -1;
			e(this.taglist).each(function(k, j) {
				if (j.node[0] === g[0]) {
					h = k;
					return false
				}
			});
			return h
		},
		get: function(g) {
			return this.taglist[g >= 0 ? g : this.taglist.length + g]
		},
		remove: function(g) {
			this.taglist = e.grep(this.taglist, function(h) {
				if (h.text === g) {
					h.node.remove()
				}
				return h.text !== g
			});
			this.fire("change")
		},
		removeAllTags: function() {
			e(this.taglist).each(function(h, g) {
				g.node.remove()
			});
			this.taglist = [];
			this.fire("change")
		},
		renderFromString: function(i) {
			var h = this;
			var g = true;
			e(i.split(this.options.delimeter)).each(function(k, j) {
				j && (g = g && h.insert(j))
			});
			return g
		},
		plainString: function() {
			var g = Array.prototype.slice;
			return e(g.call(this.taglist, g.call(arguments, 1))).map(function(j, h) {
				return h.text
			}).get().join(",")
		},
		updateOrder: function() {
			var g = e(".tag", this.dom.tag_editor);
			this.taglist.sort(function(h, i) {
				return g.index(h.node) > g.index(i.node) ? 1 : -1
			})
		},
		resize: function(g) {
			g = g || this.dom.tag_input.val().length + 2 + "em";
			this.dom.tag_writer.width(g)
		},
		on: function() {
			this._events.bind.apply(this._events, arguments);
			return this
		},
		fire: function() {
			this._events.trigger.apply(this._events, arguments)
		},
		init: function(h) {
			var g = this;
			this.renderFromString(h.tag_input.val());
			h.tag_input.val("");
			h.tag_input.bind("focus", function() {
				h.tag_editor.addClass("focus")
			}).bind("blur", function() {
				h.tag_editor.removeClass("focus");
				if (h.tag_input.val() !== "") {
					g.renderFromString(h.tag_input.val());
					h.tag_input.val("")
				}
				g.updateOrder();
				h.tag_writer.removeClass("small-tags-writer").css("width", "auto");
				g.get(-1) && h.tag_writer.insertAfter(g.get(-1).node)
			});
			h.tag_input.keyup(function() {
				var i = h.tag_input.val(),
					k = c.getCaretText(h.tag_input[0], 1),
					j;
				k = k && k.charCodeAt(0);
				if (k === a.SPACE) {
					j = c.getCaretText(h.tag_input[0], i.length);
					g.renderFromString(j) && h.tag_input.val(i.substring(j.length, i.length))
				}
				g.resize()
			}).keydown(function(l) {
				var i = c.getCharcode(l),
					k, j;
				e.each(a, function(m, n) {
					if (i === n) {
						k = m;
						return false
					}
				});
				j = g.handleKeydown[k];
				j && j.call(g, {
					event: l,
					caret: c.getCaretPos(h.tag_input[0]),
					old_val: h.tag_input.val(),
					prevTag: h.tag_writer.prev(".tag"),
					nextTag: h.tag_writer.next(".tag")
				})
			});
			h.tag_editor.delegate(".tag a", "click", function(i) {
				i.preventDefault();
				i.stopPropagation();
				g.remove(e(this).parent(".tag").find("i").text())
			}).delegate(".tag", "click", function(j) {
				var i = e(this).find("i").text();
				g.get(g.indexOf(i)).editable(h);
				g.remove(i)
			})
		},
		handleKeydown: {
			BACKSPACE: function(i) {
				if (i.caret !== 0) {
					return
				}
				var h = this.getTag(i.prevTag);
				if (h > -1) {
					var g = this.get(h);
					this.dom.tag_input.val(g.text + " " + this.dom.tag_input.val());
					c.setCaretPos(this.dom.tag_input[0], g.text.length + 1);
					this.remove(g.text)
				}
				this.resize()
			},
			LEFT: function(i) {
				if (i.caret !== 0) {
					return
				}
				var h = this.getTag(i.prevTag);
				if (h > -1) {
					var g = this.get(h);
					this.dom.tag_writer.addClass("small-tags-writer").width(i.prevTag.outerWidth()).after(i.prevTag);
					this.dom.tag_input.val(g.text);
					this.remove(g.text);
					this.insert(i.old_val, "after")
				}
				this.resize()
			},
			RIGHT: function(i) {
				if (i.caret !== i.old_val.length || !i.nextTag.length) {
					return
				}
				var h = this.getTag(i.nextTag);
				if (h > -1) {
					var g = this.get(h);
					this.dom.tag_writer.width(i.nextTag.outerWidth()).before(i.nextTag);
					this.dom.tag_input.val(g.text);
					this.remove(g.text);
					this.renderFromString(i.old_val)
				}
				c.setCaretPos(this.dom.tag_input[0], 0);
				this.resize()
			},
			HOME: function(g) {
				this.dom.tag_input.blur();
				if (this.get(0)) {
					this.get(0).editable(this.dom);
					this.remove(this.get(0).text)
				}
			},
			END: function(g) {
				this.dom.tag_input.blur();
				if (this.get(-1)) {
					this.get(-1).editable(this.dom);
					this.remove(this.get(-1).text)
				}
			},
			A: function(h) {
				if (!h.event.ctrlKey) {
					return
				}
				var g = this.dom.tag_input.val();
				var i = this.dom.tag_writer.prevAll(".tag").find("i").map(function(k, j) {
					return e(j).text()
				}).get().reverse();
				if (e.trim(g)) {
					i.push(g)
				}
				e.merge(i, this.dom.tag_writer.nextAll(".tag").find("i").map(function(k, j) {
					return e(j).text()
				}).get());
				this.dom.tag_input.val(i.join(","));
				this.resize();
				this.removeAllTags()
			}
		}
	};
	e.fn.TagEditor = function(g) {
		return this.pushStack(this.map(function() {
			return new b(e(this), e.extend(f, g || {}))
		}))
	}
})(jQuery);

function extendTagEditor(c, f) {
	var b = $(".tags-tip", f),
		e = $(".tags-folder span"),
		d = $(".tags-control", f),
		a = $(".tags-folder", f);
	c.on("error", function(h, g) {
		switch (g.type) {
		case "exist":
			var i = c.taglist[g.msg].node;
			i.addClass("tags-warn");
			setTimeout(function() {
				i.removeClass("tags-warn")
			}, 1000);
			c.dom.tag_input.val("");
			break;
		default:
			!b.is(":visible") && b.text(g.msg).fadeIn().delay(2000).fadeOut(function() {
				c.fire("error:hidden")
			});
			break
		}
	}).on("change", function(h) {

		var g = c.dom.tag_writer.find("label");
		e.each(function(k, j) {
			j = $(j);
			c.indexOf(j.text()) > -1 ? j.addClass("tag-selected") : j.removeClass("tag-selected")
		});
		c.taglist.length ? g.addClass("invisible") : g.removeClass("invisible")
	});
	d.delegate(".tags_more", "click", function(g) {
		g.preventDefault();
		a.addClass("tags-unfold");
		c.fire("plugin:expand");
		$(this).hide()
	});
	a.delegate("span", "click", function() {
		var g = $(this);
		if (g.hasClass("tag-selected")) {
			c.remove(g.text())
		} else {
			c.insert(g.text())
		}
	});
	c.dom.tag_writer.find("label").removeClass("invisible")
}
function popupTagEditor() {
	var e = $("#tag-editor-template").text(),
		d = $(".a_edit_tag"),
		c, b;
	d.click(function(f) {
		f.preventDefault();
		if (!c) {
			c = dui.Dialog({
				title: d.text(),
				width: 500,
				content: e
			}).open();
			c.node.addClass("dialog-tags");
			c.update();
			a(URL_UPDATE_TAG, DATA_TAG)
		} else {
			c.open()
		}
		if (!b) {
			var g = c.node.find(".tag-editor-bd");
			b = g.TagEditor().get(0);
			extendTagEditor(b, g);
			b.on("plugin:expand error error:hidden", function() {
				c.update()
			})
		}
	});

 
}
Do("template", "dialog", "modernizr.dnd", "uploadify", function() {
	var w = $("#form_note"),
		a = $(".control-panel", w),
		f = w.children(".images"),
		d = w.children(".videos"),
		h = $("#note_title"),
		n = $("#note_text"),
		s = w.find(".note-privacy"),
		g = w.find("#publish_note"),
		j = $("#cannot_reply"),
		c = "#note_text";
	var i = $("#note_id").val();
	var k = false;
	var b = function() {
			var A = $('<div class="error-msg tip"></div>'),
				y = w.find(".footer .error-msg"),
				z;
			A.appendTo("body");
			return {
				putAside: function(B, D) {
					clearTimeout(z);
					var E = B.offset(),
						C = B.outerWidth();
					A.text(D).css({
						top: E.top + 5,
						left: E.left + C + 5
					}).show();
					$("html, body").animate({
						scrollTop: E.top - 40
					});
					z = setTimeout(function() {
						A.hide()
					}, 5000)
				},
				showBoard: function(B) {
					y.text(B).show();
					setTimeout(function() {
						y.text("")
					}, 5000)
				},
				clear: function() {
					A.hide();
					y.hide()
				}
			}
		}();
 	function q() {
		var y = {};
		$.each(w.serializeArray(), function(z, A) {
			y[A.name] = A.value
		});
		return y
	}
	var x = $(".note-tags"),
		t, p;
	if (x.length) {
		t = x.TagEditor().get(0);
		p = $("#author_tags");
		extendTagEditor(t, x)
	}
	function o() {
		if (t) {
			p.val(t.plainString())
		}
		$.ajax({
			type: "post",
			url: AUTOSAVE_URL,
			data: q()
		})
	}
	n.keydown(function(y) {
		if (String.fromCharCode(y.which).toLowerCase() == "s" && y.ctrlKey) {
			y.preventDefault();
			o();
			return false
		}
	});
	setInterval(o, 1000 * 60 * 5);
	 
});