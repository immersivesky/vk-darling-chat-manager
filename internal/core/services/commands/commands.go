package commands

import (
	"regexp"

	"gitlab.com/immersivesky/affinitycm-vk/internal/core/ports"
)

func NewCommandsService(db ports.DB) (ports.Commands, ports.Commands) {
	pingRegexp, pingCmd := regexp.MustCompile("(?i)^(?:ping)$"), new(PingCmd)

	return map[*regexp.Regexp]ports.Command{
			pingRegexp: pingCmd,
			regexp.MustCompile("(?i)^(?:help|menu|помощь|команды)$"):   new(MenuCmd),
			regexp.MustCompile(`(?i)^(plugins|плагины)(?:\s(.*))?$`):   &PluginsCmd{db: db},
			regexp.MustCompile(`(?i)^(control|настройки)(?:\s(.*))?$`): new(ControlCmd),
			regexp.MustCompile(`(?i)^(profile|профиль)(?:\s(.*))?$`):   new(ProfileCmd),
			regexp.MustCompile(`(?i)^(e?ban)(?:\s(.*))?$`):             new(BanCmd),
		}, map[*regexp.Regexp]ports.Command{
			pingRegexp: pingCmd,
		}
}
