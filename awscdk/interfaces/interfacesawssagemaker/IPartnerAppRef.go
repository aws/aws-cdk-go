package interfacesawssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PartnerApp.
// Experimental.
type IPartnerAppRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PartnerApp resource.
	// Experimental.
	PartnerAppRef() *PartnerAppReference
}

// The jsii proxy for IPartnerAppRef
type jsiiProxy_IPartnerAppRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPartnerAppRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IPartnerAppRef) PartnerAppRef() *PartnerAppReference {
	var returns *PartnerAppReference
	_jsii_.Get(
		j,
		"partnerAppRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPartnerAppRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPartnerAppRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

