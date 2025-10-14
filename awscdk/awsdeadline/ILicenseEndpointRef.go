package awsdeadline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LicenseEndpoint.
// Experimental.
type ILicenseEndpointRef interface {
	constructs.IConstruct
	// A reference to a LicenseEndpoint resource.
	// Experimental.
	LicenseEndpointRef() *LicenseEndpointReference
}

// The jsii proxy for ILicenseEndpointRef
type jsiiProxy_ILicenseEndpointRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILicenseEndpointRef) LicenseEndpointRef() *LicenseEndpointReference {
	var returns *LicenseEndpointReference
	_jsii_.Get(
		j,
		"licenseEndpointRef",
		&returns,
	)
	return returns
}

