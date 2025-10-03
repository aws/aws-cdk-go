package awspanorama

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspanorama/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PackageVersion.
// Experimental.
type IPackageVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPackageVersionRef
type jsiiProxy_IPackageVersionRef struct {
	internal.Type__constructsIConstruct
}

