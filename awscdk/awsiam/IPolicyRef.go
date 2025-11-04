package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Policy.
// Experimental.
type IPolicyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Policy resource.
	// Experimental.
	PolicyRef() *PolicyReference
}

// The jsii proxy for IPolicyRef
type jsiiProxy_IPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IPolicyRef) PolicyRef() *PolicyReference {
	var returns *PolicyReference
	_jsii_.Get(
		j,
		"policyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

