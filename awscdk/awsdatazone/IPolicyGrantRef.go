package awsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PolicyGrant.
// Experimental.
type IPolicyGrantRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a PolicyGrant resource.
	// Experimental.
	PolicyGrantRef() *PolicyGrantReference
}

// The jsii proxy for IPolicyGrantRef
type jsiiProxy_IPolicyGrantRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IPolicyGrantRef) PolicyGrantRef() *PolicyGrantReference {
	var returns *PolicyGrantReference
	_jsii_.Get(
		j,
		"policyGrantRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyGrantRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyGrantRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

