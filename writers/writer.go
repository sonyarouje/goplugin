package writers

import "goplugin/writers/common"

//Writer interface that each plugin should implement
type Writer interface {
	Write(plugger common.Plugger)
}
