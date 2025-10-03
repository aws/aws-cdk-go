package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModelPackage.
// Experimental.
type IModelPackageRef interface {
	constructs.IConstruct
}

// The jsii proxy for IModelPackageRef
type jsiiProxy_IModelPackageRef struct {
	internal.Type__constructsIConstruct
}

