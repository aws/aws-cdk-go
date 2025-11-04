package awsmediaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FlowEntitlement.
// Experimental.
type IFlowEntitlementRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a FlowEntitlement resource.
	// Experimental.
	FlowEntitlementRef() *FlowEntitlementReference
}

// The jsii proxy for IFlowEntitlementRef
type jsiiProxy_IFlowEntitlementRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IFlowEntitlementRef) FlowEntitlementRef() *FlowEntitlementReference {
	var returns *FlowEntitlementReference
	_jsii_.Get(
		j,
		"flowEntitlementRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowEntitlementRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowEntitlementRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

