package awsappstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Entitlement.
// Experimental.
type IEntitlementRef interface {
	constructs.IConstruct
	// A reference to a Entitlement resource.
	// Experimental.
	EntitlementRef() *EntitlementReference
}

// The jsii proxy for IEntitlementRef
type jsiiProxy_IEntitlementRef struct {
	internal.Type__constructsIConstruct
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

