package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Represents the concept of a CloudFront Origin.
//
// You provide one or more origins when creating a Distribution.
// Experimental.
type IOrigin interface {
	// The method called when a given Origin is added (for the first time) to a Distribution.
	// Experimental.
	Bind(scope awscdk.Construct, options *OriginBindOptions) *OriginBindConfig
}

// The jsii proxy for IOrigin
type jsiiProxy_IOrigin struct {
	_ byte // padding
}

func (i *jsiiProxy_IOrigin) Bind(scope awscdk.Construct, options *OriginBindOptions) *OriginBindConfig {
	var returns *OriginBindConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{scope, options},
		&returns,
	)

	return returns
}

