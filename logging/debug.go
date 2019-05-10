package logging

import (
	"fmt"
	"runtime"
)

var debugLog = ".logs/debug"

var debugDisabled = true

func DisableDebug() {
	debugDisabled = true
}
func EnableDebug() {
	debugDisabled = false
}

func Debug(format string, args ...interface{}) {
	if debugDisabled {
		return
	}
	frame := getFrame(1)
	fmt.Printf("%+v", frame.File)
	fmt.Printf(":%+v\n", frame.Line)
	fmt.Printf(" --- %v\n", frame.Function)

	fmt.Printf("     --- ")
	fmt.Printf(format, args...)
	fmt.Printf("\n")
}

func getFrame(skipFrames int) runtime.Frame {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame
}
