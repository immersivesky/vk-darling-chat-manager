package commands

import (
	"regexp"

	"gitlab.com/immersivesky/affinitycm-vk/internal/core/ports"
)

func NewCommandsService() (ports.Commands, ports.Commands) {
	pingRegexp, pingCmd := regexp.MustCompile("(?i)^(?:ping)$"), new(PingCmd)

	return map[*regexp.Regexp]ports.Command{
			pingRegexp: pingCmd,
			regexp.MustCompile("(?i)^(?:help|menu|помощь|команды)$"): new(HelpCmd),
			regexp.MustCompile(`(?i)^(profile|профиль)(?:\s(.*))?$`): new(ProfileCmd),
			regexp.MustCompile(`(?i)^(e?ban)(?:\s(.*))?$`):           new(BanCmd),
		}, map[*regexp.Regexp]ports.Command{
			pingRegexp: pingCmd,
		}
}
