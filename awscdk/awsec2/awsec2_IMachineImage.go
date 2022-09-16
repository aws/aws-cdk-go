package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for classes that can select an appropriate machine image to use.
type IMachineImage interface {
	// Return the image to use in the given context.
	GetImage(scope constructs.Construct) *MachineImageConfig
}

// The jsii proxy for IMachineImage
type jsiiProxy_IMachineImage struct {
	_ byte // padding
}

func (i *jsiiProxy_IMachineImage) GetImage(scope constructs.Construct) *MachineImageConfig {
	if err := i.validateGetImageParameters(scope); err != nil {
		panic(err)
	}
	var returns *MachineImageConfig

	_jsii_.Invoke(
		i,
		"getImage",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

