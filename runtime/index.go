package runtime

import "runtime/debug"

func MainGetPanicAndLoop(f func()) {
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					debug.PrintStack()
				}
			}()

			f()
		}()
	}
}

func MainGetPanic(f func()) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
		}
	}()

	f()
}
