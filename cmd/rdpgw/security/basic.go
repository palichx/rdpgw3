package security

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
)

var (
	Hosts         []string
	HostSelection string
)


var Grouphosts map[string][]string


func SetGroupHosts(Grouphoststest map[string][]string) {
        if len(Grouphoststest) < 32 {
                log.Printf("Hosts  %s by policy handler", Grouphoststest)
        }
      Grouphosts = Grouphoststest
      for key, hosts := range  Grouphoststest {
      log.Printf("%s:\n", key)
        for _, host := range hosts {
          log.Printf("  - %s\n", host)
        }
      }

}


func CheckHost(ctx context.Context, host string) (bool, error) {
	switch HostSelection {
	case "any":
		return true, nil
	case "signed":
		// todo get from context?
		return false, errors.New("cannot verify host in 'signed' mode as token data is missing")
	case "roundrobin", "unsigned":
		s := getTunnel(ctx)
		if s.User.UserName() == "" {
			return false, errors.New("no valid session info or username found in context")
		}

		log.Printf("Checking host for user %s", s.User.UserName())
		for _, h := range Hosts {
			h = strings.Replace(h, "{{ preferred_username }}", s.User.UserName(), 1)
			log.Printf("ADDITIONAL LOG1 %s", s.User)
			log.Printf("ADDITIONAL LOG2 %s", h)
			log.Printf("ADDITIONAL LOG3 %s", host)
			if h == host {
				return true, nil
			}
		}
		return false, fmt.Errorf("invalid host %s", host)
	}

	return false, errors.New("unrecognized host selection criteria")
}
