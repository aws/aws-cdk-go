package interfacesawsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomerGatewayAssociation.
// Experimental.
type ICustomerGatewayAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CustomerGatewayAssociation resource.
	// Experimental.
	CustomerGatewayAssociationRef() *CustomerGatewayAssociationReference
}

// The jsii proxy for ICustomerGatewayAssociationRef
type jsiiProxy_ICustomerGatewayAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ICustomerGatewayAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ICustomerGatewayAssociationRef) CustomerGatewayAssociationRef() *CustomerGatewayAssociationReference {
	var returns *CustomerGatewayAssociationReference
	_jsii_.Get(
		j,
		"customerGatewayAssociationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomerGatewayAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomerGatewayAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

