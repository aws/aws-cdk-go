package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SoftwarePackageVersion.
// Experimental.
type ISoftwarePackageVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISoftwarePackageVersionRef
type jsiiProxy_ISoftwarePackageVersionRef struct {
	internal.Type__constructsIConstruct
}

