package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModelPackageGroup.
// Experimental.
type IModelPackageGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IModelPackageGroupRef
type jsiiProxy_IModelPackageGroupRef struct {
	internal.Type__constructsIConstruct
}

