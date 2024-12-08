package command

import "fmt"

type Spark struct {
	Type string
}

func NewSpark() (*Spark, error) {
	cmd := &Spark{
		Type: KeySpark,
	}
	return cmd, nil
}

func (cmd *Spark) Print()          { fmt.Println(cmd.Type) }
func (cmd *Spark) GetType() string { return cmd.Type }
