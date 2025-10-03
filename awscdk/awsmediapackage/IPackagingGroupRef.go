package awsmediapackage

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediapackage/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PackagingGroup.
// Experimental.
type IPackagingGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPackagingGroupRef
type jsiiProxy_IPackagingGroupRef struct {
	internal.Type__constructsIConstruct
}

