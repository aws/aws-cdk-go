package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SoftwarePackage.
// Experimental.
type ISoftwarePackageRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISoftwarePackageRef
type jsiiProxy_ISoftwarePackageRef struct {
	internal.Type__constructsIConstruct
}

