package interfacesawslicensemanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslicensemanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a License.
// Experimental.
type ILicenseRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a License resource.
	// Experimental.
	LicenseRef() *LicenseReference
}

// The jsii proxy for ILicenseRef
type jsiiProxy_ILicenseRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_ILicenseRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILicenseRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

