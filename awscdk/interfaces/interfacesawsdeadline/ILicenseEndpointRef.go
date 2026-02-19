package interfacesawsdeadline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LicenseEndpoint.
// Experimental.
type ILicenseEndpointRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LicenseEndpoint resource.
	// Experimental.
	LicenseEndpointRef() *LicenseEndpointReference
}

// The jsii proxy for ILicenseEndpointRef
type jsiiProxy_ILicenseEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ILicenseEndpointRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_ILicenseEndpointRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILicenseEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

