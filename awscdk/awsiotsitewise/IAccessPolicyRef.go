package awsiotsitewise

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotsitewise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessPolicy.
// Experimental.
type IAccessPolicyRef interface {
	constructs.IConstruct
	// A reference to a AccessPolicy resource.
	// Experimental.
	AccessPolicyRef() *AccessPolicyReference
}

// The jsii proxy for IAccessPolicyRef
type jsiiProxy_IAccessPolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAccessPolicyRef) AccessPolicyRef() *AccessPolicyReference {
	var returns *AccessPolicyReference
	_jsii_.Get(
		j,
		"accessPolicyRef",
		&returns,
	)
	return returns
}

