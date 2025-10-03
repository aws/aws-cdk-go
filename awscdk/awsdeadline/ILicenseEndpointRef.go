package awsdeadline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LicenseEndpoint.
// Experimental.
type ILicenseEndpointRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILicenseEndpointRef
type jsiiProxy_ILicenseEndpointRef struct {
	internal.Type__constructsIConstruct
}

