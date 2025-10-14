package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GroupVersion.
// Experimental.
type IGroupVersionRef interface {
	constructs.IConstruct
	// A reference to a GroupVersion resource.
	// Experimental.
	GroupVersionRef() *GroupVersionReference
}

// The jsii proxy for IGroupVersionRef
type jsiiProxy_IGroupVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGroupVersionRef) GroupVersionRef() *GroupVersionReference {
	var returns *GroupVersionReference
	_jsii_.Get(
		j,
		"groupVersionRef",
		&returns,
	)
	return returns
}

