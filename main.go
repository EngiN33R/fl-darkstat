package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/darklab8/fl-darkstat/darkstat/builder"
	"github.com/darklab8/fl-darkstat/darkstat/linker"
	"github.com/darklab8/fl-darkstat/darkstat/web"
	"github.com/darklab8/go-utils/goutils/utils/time_measure"
)

type Action string

func (a Action) ToStr() string { return string(a) }

const (
	Build Action = "build"
	Web   Action = "web"
)

func main() {
	var action string
	flag.StringVar(&action, "act", string(Web),
		fmt.Sprintln("action to run. Possible choices...",
			strings.Join([]string{Build.ToStr(), Web.ToStr()}, ", ")),
	)
	flag.Parse()
	fmt.Println("act:", action)

	switch Action(action) {

	case Build:
		linker.NewLinker().Link().BuildAll().RenderToLocal()
	case Web:
		var fs *builder.Filesystem
		time_measure.TimeMeasure(func(m *time_measure.TimeMeasurer) {
			var linked_build *builder.Builder
			time_measure.TimeMeasure(func(m *time_measure.TimeMeasurer) {
				linked_build = linker.NewLinker().Link()
			}, time_measure.WithMsg("linking stuff"))
			time_measure.TimeMeasure(func(m *time_measure.TimeMeasurer) {
				fs = linked_build.BuildAll()
			}, time_measure.WithMsg("building stuff"))
		}, time_measure.WithMsg("total time for web"))
		web.NewWeb(fs).Serve()
	}

}
