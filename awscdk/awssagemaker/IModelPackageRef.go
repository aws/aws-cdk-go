package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModelPackage.
// Experimental.
type IModelPackageRef interface {
	constructs.IConstruct
	// A reference to a ModelPackage resource.
	// Experimental.
	ModelPackageRef() *ModelPackageReference
}

// The jsii proxy for IModelPackageRef
type jsiiProxy_IModelPackageRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IModelPackageRef) ModelPackageRef() *ModelPackageReference {
	var returns *ModelPackageReference
	_jsii_.Get(
		j,
		"modelPackageRef",
		&returns,
	)
	return returns
}

