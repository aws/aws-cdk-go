package awscomprehend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscomprehend/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Flywheel.
// Experimental.
type IFlywheelRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Flywheel resource.
	// Experimental.
	FlywheelRef() *FlywheelReference
}

// The jsii proxy for IFlywheelRef
type jsiiProxy_IFlywheelRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IFlywheelRef) FlywheelRef() *FlywheelReference {
	var returns *FlywheelReference
	_jsii_.Get(
		j,
		"flywheelRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlywheelRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlywheelRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

