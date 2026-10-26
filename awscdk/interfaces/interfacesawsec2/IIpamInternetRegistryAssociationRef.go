package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IpamInternetRegistryAssociation.
// Experimental.
type IIpamInternetRegistryAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a IpamInternetRegistryAssociation resource.
	// Experimental.
	IpamInternetRegistryAssociationRef() *IpamInternetRegistryAssociationReference
}

// The jsii proxy for IIpamInternetRegistryAssociationRef
type jsiiProxy_IIpamInternetRegistryAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IIpamInternetRegistryAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IIpamInternetRegistryAssociationRef) IpamInternetRegistryAssociationRef() *IpamInternetRegistryAssociationReference {
	var returns *IpamInternetRegistryAssociationReference
	_jsii_.Get(
		j,
		"ipamInternetRegistryAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamInternetRegistryAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamInternetRegistryAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

