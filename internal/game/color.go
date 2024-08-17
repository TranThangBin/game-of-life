package game

type ColorValue byte

const (
	BLACK ColorValue = iota
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

var (
	fgColor = [...][]byte{
		[]byte("\033[30m"), // Black
		[]byte("\033[31m"), // Red
		[]byte("\033[32m"), // Green
		[]byte("\033[33m"), // Yellow
		[]byte("\033[34m"), // Blue
		[]byte("\033[35m"), // Magenta
		[]byte("\033[36m"), // Cyan
		[]byte("\033[37m"), // White
	}

	bgColor = [...][]byte{
		[]byte("\033[40m"), // Black
		[]byte("\033[41m"), // Red
		[]byte("\033[42m"), // Green
		[]byte("\033[43m"), // Yellow
		[]byte("\033[44m"), // Blue
		[]byte("\033[45m"), // Magenta
		[]byte("\033[46m"), // Cyan
		[]byte("\033[47m"), // White
	}

	colorReset = []byte("\033[0m")
)

type ColorBuilder struct {
	textContent []byte
	fgColor     []byte
	bgColor     []byte
}

func NewColorBuilder(content []byte) ColorBuilder {
	return ColorBuilder{
		textContent: content,
	}
}

func (c ColorBuilder) WithFgColor(color ColorValue) ColorBuilder {
	c.fgColor = fgColor[color]
	return c
}

func (c ColorBuilder) WithBgColor(color ColorValue) ColorBuilder {
	c.bgColor = bgColor[color]
	return c
}

func (c ColorBuilder) Build() []byte {
	result := c.textContent
	var reset []byte
	if len(c.fgColor) != 0 {
		result = append(c.fgColor, result...)
		reset = colorReset
	}
	if len(c.bgColor) != 0 {
		result = append(c.bgColor, result...)
		reset = colorReset
	}
	return append(result, reset...)
}
