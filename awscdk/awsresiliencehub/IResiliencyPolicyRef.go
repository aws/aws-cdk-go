package awsresiliencehub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsresiliencehub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResiliencyPolicy.
// Experimental.
type IResiliencyPolicyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ResiliencyPolicy resource.
	// Experimental.
	ResiliencyPolicyRef() *ResiliencyPolicyReference
}

// The jsii proxy for IResiliencyPolicyRef
type jsiiProxy_IResiliencyPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IResiliencyPolicyRef) ResiliencyPolicyRef() *ResiliencyPolicyReference {
	var returns *ResiliencyPolicyReference
	_jsii_.Get(
		j,
		"resiliencyPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResiliencyPolicyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResiliencyPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

