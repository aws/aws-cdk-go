package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomerGatewayAssociation.
// Experimental.
type ICustomerGatewayAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CustomerGatewayAssociation resource.
	// Experimental.
	CustomerGatewayAssociationRef() *CustomerGatewayAssociationReference
}

// The jsii proxy for ICustomerGatewayAssociationRef
type jsiiProxy_ICustomerGatewayAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ICustomerGatewayAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

