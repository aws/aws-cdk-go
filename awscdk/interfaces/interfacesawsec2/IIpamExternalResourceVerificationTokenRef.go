package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IpamExternalResourceVerificationToken.
// Experimental.
type IIpamExternalResourceVerificationTokenRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a IpamExternalResourceVerificationToken resource.
	// Experimental.
	IpamExternalResourceVerificationTokenRef() *IpamExternalResourceVerificationTokenReference
}

// The jsii proxy for IIpamExternalResourceVerificationTokenRef
type jsiiProxy_IIpamExternalResourceVerificationTokenRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IIpamExternalResourceVerificationTokenRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IIpamExternalResourceVerificationTokenRef) IpamExternalResourceVerificationTokenRef() *IpamExternalResourceVerificationTokenReference {
	var returns *IpamExternalResourceVerificationTokenReference
	_jsii_.Get(
		j,
		"ipamExternalResourceVerificationTokenRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamExternalResourceVerificationTokenRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpamExternalResourceVerificationTokenRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

