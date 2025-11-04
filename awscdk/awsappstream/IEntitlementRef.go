package awsappstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Entitlement.
// Experimental.
type IEntitlementRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Entitlement resource.
	// Experimental.
	EntitlementRef() *EntitlementReference
}

// The jsii proxy for IEntitlementRef
type jsiiProxy_IEntitlementRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IEntitlementRef) EntitlementRef() *EntitlementReference {
	var returns *EntitlementReference
	_jsii_.Get(
		j,
		"entitlementRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEntitlementRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEntitlementRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

