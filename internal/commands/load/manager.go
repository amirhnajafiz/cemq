package load

import (
	"github.com/amirhnajafiz/cemq/internal/mqtt"
	"github.com/amirhnajafiz/cemq/pkg/model"
)

// Manager handles load commands
type Manager interface {
	Generate(input *model.Load) error
}

func New(conn mqtt.CLI, logsFlag bool) Manager {
	return &load{
		logsFlag: logsFlag,
		conn:     conn,
	}
}
