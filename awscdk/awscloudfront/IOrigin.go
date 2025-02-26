package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the concept of a CloudFront Origin.
//
// You provide one or more origins when creating a Distribution.
type IOrigin interface {
	// The method called when a given Origin is added (for the first time) to a Distribution.
	Bind(scope constructs.Construct, options *OriginBindOptions) *OriginBindConfig
}

// The jsii proxy for IOrigin
type jsiiProxy_IOrigin struct {
	_ byte // padding
}

func (i *jsiiProxy_IOrigin) Bind(scope constructs.Construct, options *OriginBindOptions) *OriginBindConfig {
	if err := i.validateBindParameters(scope, options); err != nil {
		panic(err)
	}
	var returns *OriginBindConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

