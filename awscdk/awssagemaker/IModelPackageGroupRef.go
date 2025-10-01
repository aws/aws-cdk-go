package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModelPackageGroup.
// Experimental.
type IModelPackageGroupRef interface {
	constructs.IConstruct
	// A reference to a ModelPackageGroup resource.
	// Experimental.
	ModelPackageGroupRef() *ModelPackageGroupReference
}

// The jsii proxy for IModelPackageGroupRef
type jsiiProxy_IModelPackageGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IModelPackageGroupRef) ModelPackageGroupRef() *ModelPackageGroupReference {
	var returns *ModelPackageGroupReference
	_jsii_.Get(
		j,
		"modelPackageGroupRef",
		&returns,
	)
	return returns
}

