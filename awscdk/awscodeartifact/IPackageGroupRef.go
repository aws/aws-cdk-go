package awscodeartifact

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodeartifact/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PackageGroup.
// Experimental.
type IPackageGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPackageGroupRef
type jsiiProxy_IPackageGroupRef struct {
	internal.Type__constructsIConstruct
}

