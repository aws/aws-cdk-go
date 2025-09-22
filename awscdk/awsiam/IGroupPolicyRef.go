package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GroupPolicy.
// Experimental.
type IGroupPolicyRef interface {
	constructs.IConstruct
	// A reference to a GroupPolicy resource.
	// Experimental.
	GroupPolicyRef() *GroupPolicyReference
}

// The jsii proxy for IGroupPolicyRef
type jsiiProxy_IGroupPolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGroupPolicyRef) GroupPolicyRef() *GroupPolicyReference {
	var returns *GroupPolicyReference
	_jsii_.Get(
		j,
		"groupPolicyRef",
		&returns,
	)
	return returns
}

