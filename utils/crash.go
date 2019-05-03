package utils

func Crash(should bool, message string) {
	if should {
		panic(message)
	}
}
