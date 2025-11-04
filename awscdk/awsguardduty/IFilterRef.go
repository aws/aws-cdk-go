package awsguardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Filter.
// Experimental.
type IFilterRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Filter resource.
	// Experimental.
	FilterRef() *FilterReference
}

// The jsii proxy for IFilterRef
type jsiiProxy_IFilterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IFilterRef) FilterRef() *FilterReference {
	var returns *FilterReference
	_jsii_.Get(
		j,
		"filterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFilterRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFilterRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

