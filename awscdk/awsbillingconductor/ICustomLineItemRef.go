package awsbillingconductor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbillingconductor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomLineItem.
// Experimental.
type ICustomLineItemRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CustomLineItem resource.
	// Experimental.
	CustomLineItemRef() *CustomLineItemReference
}

// The jsii proxy for ICustomLineItemRef
type jsiiProxy_ICustomLineItemRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICustomLineItemRef) CustomLineItemRef() *CustomLineItemReference {
	var returns *CustomLineItemReference
	_jsii_.Get(
		j,
		"customLineItemRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomLineItemRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomLineItemRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

