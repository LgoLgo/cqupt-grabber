package client

import (
	"context"
	"strconv"
	"strings"

	"github.com/LgoLgo/cqupt-grabber/cqupt"
	model "github.com/LgoLgo/cqupt-grabber/model"
)

type Runner interface {
	Mode() Mode
	ForceBlockSearch() bool
	Start(ctx context.Context, cookie string, keywords []string, classNo string, logf func(level LogLevel, msg string))
}

func NewRunner(mode Mode) Runner {
	switch mode {
	case ModeSecXk:
		return &secRunner{tool: cqupt.NewForSecXk()}
	case ModeSmallTerm:
		return &smallRunner{tool: cqupt.NewForSmallTerm()}
	default:
		return &commonRunner{tool: cqupt.New()}
	}
}

type commonRunner struct {
	tool *cqupt.Engine
}

func (r *commonRunner) Mode() Mode             { return ModeCommon }
func (r *commonRunner) ForceBlockSearch() bool { return true }

func (r *commonRunner) Start(ctx context.Context, cookie string, keywords []string, _ string, logf func(level LogLevel, msg string)) {
	logf(LogInfo, "Common：BlockSearch 搜课中...")
	var loads []string = r.tool.Queryer.BlockSearch(cookie, keywords)
	logf(LogInfo, "Common：loads 数量: "+strconv.Itoa(len(loads)))

	if len(loads) == 0 {
		logf(LogWarn, "Common：loads 为空，结束")
		return
	}

	logf(LogInfo, "Common：开始抢课...")
	r.tool.Grabber.LoopRobWithCtx(ctx, cookie, loads)
	logf(LogSuccess, "Common：LoopRob 结束")
}

type smallRunner struct {
	tool *cqupt.SmallEngine
}

func (r *smallRunner) Mode() Mode             { return ModeSmallTerm }
func (r *smallRunner) ForceBlockSearch() bool { return true }

func (r *smallRunner) Start(ctx context.Context, cookie string, keywords []string, classNo string, logf func(level LogLevel, msg string)) {
	classNo = strings.TrimSpace(classNo)
	if classNo == "" {
		logf(LogError, "SmallTerm：classNo 不能为空")
		return
	}

	logf(LogInfo, "SmallTerm：BlockSearch 搜课中...")
	var loads []model.MetaData = r.tool.Queryer.BlockSearch(cookie, keywords, classNo)
	logf(LogInfo, "SmallTerm：loads 数量: "+strconv.Itoa(len(loads)))

	if len(loads) == 0 {
		logf(LogWarn, "SmallTerm：loads 为空，结束")
		return
	}

	logf(LogInfo, "SmallTerm：开始抢课...")
	r.tool.Grabber.LoopRobWithCtx(ctx, cookie, loads)
	logf(LogSuccess, "SmallTerm：LoopRob 返回结束")
}

type secRunner struct {
	tool *cqupt.SecEngine
}

func (r *secRunner) Mode() Mode             { return ModeSecXk }
func (r *secRunner) ForceBlockSearch() bool { return true }

func (r *secRunner) Start(ctx context.Context, cookie string, keywords []string, _ string, logf func(level LogLevel, msg string)) {
	logf(LogInfo, "SecXk：BlockSearch 搜课中...")
	var loads []model.SecCourseData = r.tool.Queryer.BlockSearch(cookie, keywords)
	logf(LogInfo, "SecXk：loads 数量: "+strconv.Itoa(len(loads)))

	if len(loads) == 0 {
		logf(LogWarn, "SecXk：loads 为空，结束")
		return
	}

	logf(LogInfo, "SecXk：开始抢课...")
	r.tool.Grabber.LoopRobWithCtx(ctx, cookie, loads)
	logf(LogSuccess, "SecXk：LoopRob 返回结束")
}
