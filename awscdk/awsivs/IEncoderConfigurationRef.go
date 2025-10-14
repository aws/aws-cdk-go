package awsivs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EncoderConfiguration.
// Experimental.
type IEncoderConfigurationRef interface {
	constructs.IConstruct
	// A reference to a EncoderConfiguration resource.
	// Experimental.
	EncoderConfigurationRef() *EncoderConfigurationReference
}

// The jsii proxy for IEncoderConfigurationRef
type jsiiProxy_IEncoderConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEncoderConfigurationRef) EncoderConfigurationRef() *EncoderConfigurationReference {
	var returns *EncoderConfigurationReference
	_jsii_.Get(
		j,
		"encoderConfigurationRef",
		&returns,
	)
	return returns
}

