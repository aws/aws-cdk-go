package awsidentitystore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsidentitystore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Group.
// Experimental.
type IGroupRef interface {
	constructs.IConstruct
	// A reference to a Group resource.
	// Experimental.
	GroupRef() *GroupReference
}

// The jsii proxy for IGroupRef
type jsiiProxy_IGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGroupRef) GroupRef() *GroupReference {
	var returns *GroupReference
	_jsii_.Get(
		j,
		"groupRef",
		&returns,
	)
	return returns
}

