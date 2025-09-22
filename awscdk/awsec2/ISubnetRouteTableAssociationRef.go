package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubnetRouteTableAssociation.
// Experimental.
type ISubnetRouteTableAssociationRef interface {
	constructs.IConstruct
	// A reference to a SubnetRouteTableAssociation resource.
	// Experimental.
	SubnetRouteTableAssociationRef() *SubnetRouteTableAssociationReference
}

// The jsii proxy for ISubnetRouteTableAssociationRef
type jsiiProxy_ISubnetRouteTableAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISubnetRouteTableAssociationRef) SubnetRouteTableAssociationRef() *SubnetRouteTableAssociationReference {
	var returns *SubnetRouteTableAssociationReference
	_jsii_.Get(
		j,
		"subnetRouteTableAssociationRef",
		&returns,
	)
	return returns
}

