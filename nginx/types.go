package nginx

type Submodule struct {
	Name        string
	Lines		[]string
	Submodules  []Submodule
}

type Root Submodule
