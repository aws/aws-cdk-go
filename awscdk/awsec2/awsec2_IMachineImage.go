package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Interface for classes that can select an appropriate machine image to use.
// Experimental.
type IMachineImage interface {
	// Return the image to use in the given context.
	// Experimental.
	GetImage(scope awscdk.Construct) *MachineImageConfig
}

// The jsii proxy for IMachineImage
type jsiiProxy_IMachineImage struct {
	_ byte // padding
}

func (i *jsiiProxy_IMachineImage) GetImage(scope awscdk.Construct) *MachineImageConfig {
	var returns *MachineImageConfig

	_jsii_.Invoke(
		i,
		"getImage",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

