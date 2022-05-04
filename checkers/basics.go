package checkers

import (
	"errors"
	"fmt"
	"os"

	dotenv "github.com/GalacticDocs/unitylogger/blob/GOlang/config"
)



func init(open int, close int, replace string) (string, error) {
	filter_empty("\x1b["+open+"m", "\x1b["+close+"m", replace)
}
