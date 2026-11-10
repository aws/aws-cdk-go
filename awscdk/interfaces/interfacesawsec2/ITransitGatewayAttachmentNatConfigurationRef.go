package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayAttachmentNatConfiguration.
// Experimental.
type ITransitGatewayAttachmentNatConfigurationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TransitGatewayAttachmentNatConfiguration resource.
	// Experimental.
	TransitGatewayAttachmentNatConfigurationRef() *TransitGatewayAttachmentNatConfigurationReference
}

// The jsii proxy for ITransitGatewayAttachmentNatConfigurationRef
type jsiiProxy_ITransitGatewayAttachmentNatConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITransitGatewayAttachmentNatConfigurationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ITransitGatewayAttachmentNatConfigurationRef) TransitGatewayAttachmentNatConfigurationRef() *TransitGatewayAttachmentNatConfigurationReference {
	var returns *TransitGatewayAttachmentNatConfigurationReference
	_jsii_.Get(
		j,
		"transitGatewayAttachmentNatConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayAttachmentNatConfigurationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayAttachmentNatConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

