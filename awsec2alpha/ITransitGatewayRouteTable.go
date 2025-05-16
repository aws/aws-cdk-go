package awsec2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awsec2alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Transit Gateway Route Table.
// Experimental.
type ITransitGatewayRouteTable interface {
	awscdk.IResource
	awsec2.IRouteTable
	// Associate the provided Attachments with this route table.
	//
	// Returns: ITransitGatewayRouteTableAssociation.
	// Experimental.
	AddAssociation(id *string, transitGatewayAttachment ITransitGatewayAttachment) ITransitGatewayRouteTableAssociation
	// Add a blackhole route to this route table.
	//
	// Returns: ITransitGatewayRoute.
	// Experimental.
	AddBlackholeRoute(id *string, destinationCidr *string) ITransitGatewayRoute
	// Add an active route to this route table.
	//
	// Returns: ITransitGatewayRoute.
	// Experimental.
	AddRoute(id *string, transitGatewayAttachment ITransitGatewayAttachment, destinationCidr *string) ITransitGatewayRoute
	// Enable propagation from the provided Attachments to this route table.
	//
	// Returns: ITransitGatewayRouteTablePropagation.
	// Experimental.
	EnablePropagation(id *string, transitGatewayAttachment ITransitGatewayAttachment) ITransitGatewayRouteTablePropagation
}

// The jsii proxy for ITransitGatewayRouteTable
type jsiiProxy_ITransitGatewayRouteTable struct {
	internal.Type__awscdkIResource
	internal.Type__awsec2IRouteTable
}

func (i *jsiiProxy_ITransitGatewayRouteTable) AddAssociation(id *string, transitGatewayAttachment ITransitGatewayAttachment) ITransitGatewayRouteTableAssociation {
	if err := i.validateAddAssociationParameters(id, transitGatewayAttachment); err != nil {
		panic(err)
	}
	var returns ITransitGatewayRouteTableAssociation

	_jsii_.Invoke(
		i,
		"addAssociation",
		[]interface{}{id, transitGatewayAttachment},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITransitGatewayRouteTable) AddBlackholeRoute(id *string, destinationCidr *string) ITransitGatewayRoute {
	if err := i.validateAddBlackholeRouteParameters(id, destinationCidr); err != nil {
		panic(err)
	}
	var returns ITransitGatewayRoute

	_jsii_.Invoke(
		i,
		"addBlackholeRoute",
		[]interface{}{id, destinationCidr},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITransitGatewayRouteTable) AddRoute(id *string, transitGatewayAttachment ITransitGatewayAttachment, destinationCidr *string) ITransitGatewayRoute {
	if err := i.validateAddRouteParameters(id, transitGatewayAttachment, destinationCidr); err != nil {
		panic(err)
	}
	var returns ITransitGatewayRoute

	_jsii_.Invoke(
		i,
		"addRoute",
		[]interface{}{id, transitGatewayAttachment, destinationCidr},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITransitGatewayRouteTable) EnablePropagation(id *string, transitGatewayAttachment ITransitGatewayAttachment) ITransitGatewayRouteTablePropagation {
	if err := i.validateEnablePropagationParameters(id, transitGatewayAttachment); err != nil {
		panic(err)
	}
	var returns ITransitGatewayRouteTablePropagation

	_jsii_.Invoke(
		i,
		"enablePropagation",
		[]interface{}{id, transitGatewayAttachment},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITransitGatewayRouteTable) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_ITransitGatewayRouteTable) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayRouteTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayRouteTable) RouteTableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeTableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayRouteTable) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

