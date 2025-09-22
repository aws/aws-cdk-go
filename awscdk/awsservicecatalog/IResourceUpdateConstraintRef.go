package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceUpdateConstraint.
// Experimental.
type IResourceUpdateConstraintRef interface {
	constructs.IConstruct
	// A reference to a ResourceUpdateConstraint resource.
	// Experimental.
	ResourceUpdateConstraintRef() *ResourceUpdateConstraintReference
}

// The jsii proxy for IResourceUpdateConstraintRef
type jsiiProxy_IResourceUpdateConstraintRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResourceUpdateConstraintRef) ResourceUpdateConstraintRef() *ResourceUpdateConstraintReference {
	var returns *ResourceUpdateConstraintReference
	_jsii_.Get(
		j,
		"resourceUpdateConstraintRef",
		&returns,
	)
	return returns
}

