package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayPolicyTableEntry.
// Experimental.
type ITransitGatewayPolicyTableEntryRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TransitGatewayPolicyTableEntry resource.
	// Experimental.
	TransitGatewayPolicyTableEntryRef() *TransitGatewayPolicyTableEntryReference
}

// The jsii proxy for ITransitGatewayPolicyTableEntryRef
type jsiiProxy_ITransitGatewayPolicyTableEntryRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITransitGatewayPolicyTableEntryRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ITransitGatewayPolicyTableEntryRef) TransitGatewayPolicyTableEntryRef() *TransitGatewayPolicyTableEntryReference {
	var returns *TransitGatewayPolicyTableEntryReference
	_jsii_.Get(
		j,
		"transitGatewayPolicyTableEntryRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayPolicyTableEntryRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayPolicyTableEntryRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

