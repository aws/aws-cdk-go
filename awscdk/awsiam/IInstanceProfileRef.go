package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InstanceProfile.
// Experimental.
type IInstanceProfileRef interface {
	constructs.IConstruct
	// A reference to a InstanceProfile resource.
	// Experimental.
	InstanceProfileRef() *InstanceProfileReference
}

// The jsii proxy for IInstanceProfileRef
type jsiiProxy_IInstanceProfileRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IInstanceProfileRef) InstanceProfileRef() *InstanceProfileReference {
	var returns *InstanceProfileReference
	_jsii_.Get(
		j,
		"instanceProfileRef",
		&returns,
	)
	return returns
}

