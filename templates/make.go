package templates

func MakefileTemplate() string {
	return `
build:
		go build -o aoc -ldflags "-s -w"
`
}
