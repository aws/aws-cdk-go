package awsguardduty

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsguardduty/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Member.
// Experimental.
type IMemberRef interface {
	constructs.IConstruct
	// A reference to a Member resource.
	// Experimental.
	MemberRef() *MemberReference
}

// The jsii proxy for IMemberRef
type jsiiProxy_IMemberRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMemberRef) MemberRef() *MemberReference {
	var returns *MemberReference
	_jsii_.Get(
		j,
		"memberRef",
		&returns,
	)
	return returns
}

