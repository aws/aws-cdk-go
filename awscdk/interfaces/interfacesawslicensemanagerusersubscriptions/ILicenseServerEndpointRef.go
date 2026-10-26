package interfacesawslicensemanagerusersubscriptions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslicensemanagerusersubscriptions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LicenseServerEndpoint.
// Experimental.
type ILicenseServerEndpointRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LicenseServerEndpoint resource.
	// Experimental.
	LicenseServerEndpointRef() *LicenseServerEndpointReference
}

// The jsii proxy for ILicenseServerEndpointRef
type jsiiProxy_ILicenseServerEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ILicenseServerEndpointRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ILicenseServerEndpointRef) LicenseServerEndpointRef() *LicenseServerEndpointReference {
	var returns *LicenseServerEndpointReference
	_jsii_.Get(
		j,
		"licenseServerEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILicenseServerEndpointRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILicenseServerEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

