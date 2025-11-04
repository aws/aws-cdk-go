package awsdeadline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Limit.
// Experimental.
type ILimitRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Limit resource.
	// Experimental.
	LimitRef() *LimitReference
}

// The jsii proxy for ILimitRef
type jsiiProxy_ILimitRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ILimitRef) LimitRef() *LimitReference {
	var returns *LimitReference
	_jsii_.Get(
		j,
		"limitRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILimitRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILimitRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

