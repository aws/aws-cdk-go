package awsapprunner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapprunner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AutoScalingConfiguration.
// Experimental.
type IAutoScalingConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a AutoScalingConfiguration resource.
	// Experimental.
	AutoScalingConfigurationRef() *AutoScalingConfigurationReference
}

// The jsii proxy for IAutoScalingConfigurationRef
type jsiiProxy_IAutoScalingConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IAutoScalingConfigurationRef) AutoScalingConfigurationRef() *AutoScalingConfigurationReference {
	var returns *AutoScalingConfigurationReference
	_jsii_.Get(
		j,
		"autoScalingConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

