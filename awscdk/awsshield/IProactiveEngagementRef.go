package awsshield

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsshield/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProactiveEngagement.
// Experimental.
type IProactiveEngagementRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ProactiveEngagement resource.
	// Experimental.
	ProactiveEngagementRef() *ProactiveEngagementReference
}

// The jsii proxy for IProactiveEngagementRef
type jsiiProxy_IProactiveEngagementRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IProactiveEngagementRef) ProactiveEngagementRef() *ProactiveEngagementReference {
	var returns *ProactiveEngagementReference
	_jsii_.Get(
		j,
		"proactiveEngagementRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProactiveEngagementRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProactiveEngagementRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

