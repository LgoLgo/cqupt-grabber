package client

import (
	"context"
	"errors"
	"strings"
	"sync/atomic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func InitCTLs(a fyne.App, w fyne.Window) {
	modeSel := widget.NewSelect([]string{string(ModeCommon), string(ModeSecXk), string(ModeSmallTerm)}, nil)
	modeSel.SetSelected(string(ModeCommon))

	cookieEntry := widget.NewMultiLineEntry()
	cookieEntry.SetPlaceHolder("粘贴你的 cookie")
	cookieEntry.SetMinRowsVisible(2)

	classNoEntry := widget.NewEntry()
	classNoEntry.SetPlaceHolder("小学期班级号（仅 SmallTerm 需要）")
	classNoEntry.Disable()

	keywordsEntry := widget.NewMultiLineEntry()
	keywordsEntry.SetPlaceHolder("关键词（多行）：一行一个，用于 BlockSearch 搜课")
	keywordsEntry.SetMinRowsVisible(6)

	logRT := widget.NewRichText()
	logRT.Wrapping = fyne.TextWrapWord
	logRT.Scroll = 3

	logScroll := container.NewVScroll(logRT)
	logScroll.SetMinSize(fyne.NewSize(600, 220))

	appendLog := func(level LogLevel, msg string) {
		fyne.Do(func() {
			var colorName fyne.ThemeColorName
			var style fyne.TextStyle

			switch level {
			case LogError:
				colorName = theme.ColorNameError
				style = fyne.TextStyle{Bold: true}
			case LogWarn:
				colorName = theme.ColorNameWarning
			case LogSuccess:
				colorName = theme.ColorNamePrimary
				style = fyne.TextStyle{Bold: true}
			default:
				colorName = theme.ColorNameForeground
			}

			logRT.Segments = append(logRT.Segments, &widget.TextSegment{
				Text: msg + "\n",
				Style: widget.RichTextStyle{
					ColorName: colorName,
					TextStyle: style,
				},
			})
			logRT.Refresh()
			logScroll.ScrollToBottom()
		})
	}

	HookStdLog(appendLog)

	clearLogBtn := widget.NewButton("清空日志", func() {
		logRT.Segments = nil
		logRT.Refresh()
	})

	modeSel.OnChanged = func(v string) {
		m := Mode(v)
		if m == ModeSmallTerm {
			classNoEntry.Enable()
		} else {
			classNoEntry.Disable()
			classNoEntry.SetText("")
		}
	}

	var running int32
	var cancel context.CancelFunc

	startBtn := widget.NewButton("开始抢", func() {
		if !atomic.CompareAndSwapInt32(&running, 0, 1) {
			dialog.ShowInformation("提示", "已经在运行了", w)
			return
		}

		cookie := strings.TrimSpace(cookieEntry.Text)
		if cookie == "" {
			atomic.StoreInt32(&running, 0)
			dialog.ShowError(errors.New("参数错误：cookie 不能为空"), w)
			return
		}

		keywords := splitLines(keywordsEntry.Text)
		if len(keywords) == 0 {
			atomic.StoreInt32(&running, 0)
			dialog.ShowError(errors.New("参数错误：关键词不能为空"), w)
			return
		}

		mode := Mode(modeSel.Selected)
		r := NewRunner(mode)

		ctx, c := context.WithCancel(context.Background())
		cancel = c

		appendLog(LogInfo, "启动中...")
		appendLog(LogInfo, "模式: "+modeSel.Selected)

		go func() {
			defer atomic.StoreInt32(&running, 0)

			go func() {
				<-ctx.Done()
				appendLog(LogInfo, "已请求停止")
			}()

			r.Start(ctx, cookie, keywords, classNoEntry.Text, appendLog)
		}()
	})

	stopBtn := widget.NewButton("停止", func() {
		if atomic.LoadInt32(&running) == 0 {
			return
		}
		if cancel != nil {
			cancel()
		}
	})

	w.SetContent(container.NewVBox(
		widget.NewLabelWithStyle("cqupt-grabber", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewForm(
			widget.NewFormItem("模式", modeSel),
			widget.NewFormItem("Cookie", cookieEntry),
			widget.NewFormItem("ClassNo", classNoEntry),
		),
		widget.NewLabel("关键词（用于搜课）"),
		keywordsEntry,
		container.NewHBox(startBtn, stopBtn, clearLogBtn),
		widget.NewSeparator(),
		widget.NewLabelWithStyle("日志", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		logScroll,
	))
}
