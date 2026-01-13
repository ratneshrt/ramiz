package language

type LanguageConfig struct {
	CompileCmd []string
	RunCmd     []string
	FileName   string
}

var Registry = map[string]LanguageConfig{
	"python": {
		RunCmd:   []string{"python3", "main.py"},
		FileName: "main.py",
	},
	"c": {
		CompileCmd: []string{"gcc", "main.c", "-o", "main"},
		RunCmd:     []string{"./main"},
		FileName:   "main.c",
	},
	"go": {
		RunCmd:   []string{"go", "run", "main.go"},
		FileName: "main.go",
	},
}
