package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceGroupConfig.
// Experimental.
type IInstanceGroupConfigRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a InstanceGroupConfig resource.
	// Experimental.
	InstanceGroupConfigRef() *InstanceGroupConfigReference
}

// The jsii proxy for IInstanceGroupConfigRef
type jsiiProxy_IInstanceGroupConfigRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IInstanceGroupConfigRef) InstanceGroupConfigRef() *InstanceGroupConfigReference {
	var returns *InstanceGroupConfigReference
	_jsii_.Get(
		j,
		"instanceGroupConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceGroupConfigRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceGroupConfigRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

