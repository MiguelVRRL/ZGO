package managing

import (
  "net"
)

type worker struct {
  config Config
  listener net.Listener
}

