package awspanorama

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspanorama/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Package.
// Experimental.
type IPackageRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPackageRef
type jsiiProxy_IPackageRef struct {
	internal.Type__constructsIConstruct
}

