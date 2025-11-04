package awslicensemanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslicensemanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Grant.
// Experimental.
type IGrantRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Grant resource.
	// Experimental.
	GrantRef() *GrantReference
}

// The jsii proxy for IGrantRef
type jsiiProxy_IGrantRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IGrantRef) GrantRef() *GrantReference {
	var returns *GrantReference
	_jsii_.Get(
		j,
		"grantRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGrantRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGrantRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

