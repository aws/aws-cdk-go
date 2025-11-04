package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectionGroup.
// Experimental.
type IConnectionGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ConnectionGroup resource.
	// Experimental.
	ConnectionGroupRef() *ConnectionGroupReference
}

// The jsii proxy for IConnectionGroupRef
type jsiiProxy_IConnectionGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IConnectionGroupRef) ConnectionGroupRef() *ConnectionGroupReference {
	var returns *ConnectionGroupReference
	_jsii_.Get(
		j,
		"connectionGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectionGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectionGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

