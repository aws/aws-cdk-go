package awslicensemanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslicensemanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a License.
// Experimental.
type ILicenseRef interface {
	constructs.IConstruct
	// A reference to a License resource.
	// Experimental.
	LicenseRef() *LicenseReference
}

// The jsii proxy for ILicenseRef
type jsiiProxy_ILicenseRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILicenseRef) LicenseRef() *LicenseReference {
	var returns *LicenseReference
	_jsii_.Get(
		j,
		"licenseRef",
		&returns,
	)
	return returns
}

