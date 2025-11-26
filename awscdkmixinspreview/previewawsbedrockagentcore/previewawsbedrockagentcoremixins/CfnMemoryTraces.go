package previewawsbedrockagentcoremixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Builder for CfnMemoryLogsMixin to generate TRACES for CfnMemory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMemoryTraces := awscdkmixinspreview.Mixins.NewCfnMemoryTraces()
//
type CfnMemoryTraces interface {
	// Send traces to X-Ray.
	ToXRay() CfnMemoryLogsMixin
}

// The jsii proxy struct for CfnMemoryTraces
type jsiiProxy_CfnMemoryTraces struct {
	_ byte // padding
}

// Experimental.
func NewCfnMemoryTraces() CfnMemoryTraces {
	_init_.Initialize()

	j := jsiiProxy_CfnMemoryTraces{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryTraces",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewCfnMemoryTraces_Override(c CfnMemoryTraces) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryTraces",
		nil, // no parameters
		c,
	)
}

func (c *jsiiProxy_CfnMemoryTraces) ToXRay() CfnMemoryLogsMixin {
	var returns CfnMemoryLogsMixin

	_jsii_.Invoke(
		c,
		"toXRay",
		nil, // no parameters
		&returns,
	)

	return returns
}

