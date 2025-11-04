package awsivs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EncoderConfiguration.
// Experimental.
type IEncoderConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a EncoderConfiguration resource.
	// Experimental.
	EncoderConfigurationRef() *EncoderConfigurationReference
}

// The jsii proxy for IEncoderConfigurationRef
type jsiiProxy_IEncoderConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IEncoderConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEncoderConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

