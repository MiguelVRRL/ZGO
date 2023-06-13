package managing

type Master struct {

  // list of servers working
  servers []worker
  // signal that take the key: CTRL + c for sleep the programm
  shutdown chan bool
  // states of servers indicating if the server may be to sleep 
  states []chan bool

}


func (m Master) Delegate()  {

}
