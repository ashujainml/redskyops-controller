package webhook

import (
	"github.com/gramLabs/okeanos/pkg/webhook/patch"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, patch.Add)
}