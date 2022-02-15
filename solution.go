package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	msg := emoji.Sprint("Hello :world_map:")

	return msg
}
