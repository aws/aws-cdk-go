package awsec2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awsec2alpha/v2/internal"
)

// Represents a Transit Gateway Route Table Association.
// Experimental.
type ITransitGatewayAssociation interface {
	awscdk.IResource
	// The ID of the transit gateway route table association.
	// Experimental.
	TransitGatewayAssociationId() *string
}

// The jsii proxy for ITransitGatewayAssociation
type jsiiProxy_ITransitGatewayAssociation struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ITransitGatewayAssociation) TransitGatewayAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayAssociationId",
		&returns,
	)
	return returns
}

