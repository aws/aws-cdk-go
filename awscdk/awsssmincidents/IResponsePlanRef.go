package awsssmincidents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmincidents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResponsePlan.
// Experimental.
type IResponsePlanRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ResponsePlan resource.
	// Experimental.
	ResponsePlanRef() *ResponsePlanReference
}

// The jsii proxy for IResponsePlanRef
type jsiiProxy_IResponsePlanRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IResponsePlanRef) ResponsePlanRef() *ResponsePlanReference {
	var returns *ResponsePlanReference
	_jsii_.Get(
		j,
		"responsePlanRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResponsePlanRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResponsePlanRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

