package automate

import "strconv"

type path struct {
	a int
	b int
}

// Automate ....
type Automate struct {
	cursor    int
	symbols   []string
	paths     []path
	fnBegin   func()
	fnChange  func(a, b int, c string)
	fnDone    func()
	fnError   func()
	fnByPass  func(char string)
	isRunning bool
	steps     []int
}

// New ...
func New(symbols ...string) *Automate {
	auto := &Automate{
		cursor:    -1,
		paths:     []path{},
		symbols:   []string{},
		fnBegin:   func() {},
		fnChange:  func(a, b int, c string) {},
		fnDone:    func() {},
		fnError:   func() {},
		fnByPass:  func(c string) {},
		isRunning: true,
		steps:     []int{},
	}

	auto.Symbols(symbols...)

	return auto
}

// Empty ...
func Empty() *Automate {
	return New()
}

// OnByPass ...
func (a *Automate) OnByPass(fn func(c string)) {
	a.fnByPass = fn
}

// Lasts ...
func (a *Automate) Lasts(n int) []int {
	steps := []int{}
	size := len(a.steps) - 1 - n

	if size < 0 {
		size = -1
	}

	for i := len(a.steps) - 1; i > size; i-- {
		steps = append(steps, a.steps[i])
	}

	return steps
}

// Symbols ...
func (a *Automate) Symbols(symbols ...string) *Automate {
	if len(symbols) == 1 && len(symbols[0]) > 1 {
		for i := 0; i < len(symbols[0]); i++ {
			a.symbols = append(a.symbols, string(symbols[0][i]))
		}
	} else {
		a.symbols = symbols
	}

	return a
}

// AddPath ...
func (a *Automate) AddPath(i, f int) *Automate {
	a.paths = append(a.paths, path{i, f})

	return a
}

// Resume ...
func (a *Automate) Resume() *Automate {
	a.isRunning = true

	return a
}

// Stop ...
func (a *Automate) Stop() *Automate {
	a.isRunning = false

	return a
}

// Reset ...
func (a *Automate) Reset() *Automate {
	a.Resume()
	a.cursor = -1

	return a
}

// OnError ...
func (a *Automate) OnError(fn func()) *Automate {
	a.fnError = fn

	return a
}

// OnBegin ...
func (a *Automate) OnBegin(fn func()) *Automate {
	a.fnBegin = fn

	return a
}

// IsDone ...
func (a *Automate) IsDone() bool {
	return a.cursor == a.paths[len(a.paths)-1].b
}

// OnDone ...
func (a *Automate) OnDone(fn func()) *Automate {
	a.fnDone = fn

	return a
}

// OnChange ...
func (a *Automate) OnChange(fn func(a, b int, c string)) *Automate {
	a.fnChange = fn

	return a
}

// Run ...
func (a *Automate) Run(char string) bool {
	if !a.isRunning {
		return true
	}
	a.fnByPass(char)
	indexes := a.indexes(char)

	if a.cursor == -1 {
		if len(indexes) > 0 && indexes[0] == 0 {
			a.cursor = 0
			a.fnBegin()
			a.fnChange(-1, 0, char)
		}

		return true
	}

	path := a.getPath(a.cursor, indexes)
	if path != nil {
		a.fnChange(path.a, path.b, char)
		a.steps = append(a.steps, path.b)

		a.cursor = path.b

		if a.IsDone() {
			a.fnDone()
		}

		return true
	}

	a.fnError()
	return false
}

func (a *Automate) getPath(_a int, indexes []int) *path {
	for j := 0; j < len(indexes); j++ {
		b := indexes[j]

		for i := 0; i < len(a.paths); i++ {
			path := a.paths[i]
			if path.a == _a && path.b == b {
				return &path
			}
		}
	}

	return nil
}

func (a *Automate) indexes(char string) []int {
	indexes := []int{}

	for i := 0; i < len(a.symbols); i++ {
		if a.symbols[i] == char {
			indexes = append(indexes, i)
		}
	}

	if _, err := strconv.ParseInt(char, 10, 64); err == nil {
		for i := 0; i < len(a.symbols); i++ {
			if a.symbols[i] == "[1-9]" {
				indexes = append(indexes, i)
			}
		}
	}

	if len(indexes) > 0 {
		return indexes
	}

	for i := 0; i < len(a.symbols); i++ {
		if a.symbols[i] == "*" {
			indexes = append(indexes, i)
		}
	}

	return indexes
}
