package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IpPoolRouteTableAssociation.
// Experimental.
type IIpPoolRouteTableAssociationRef interface {
	constructs.IConstruct
	// A reference to a IpPoolRouteTableAssociation resource.
	// Experimental.
	IpPoolRouteTableAssociationRef() *IpPoolRouteTableAssociationReference
}

// The jsii proxy for IIpPoolRouteTableAssociationRef
type jsiiProxy_IIpPoolRouteTableAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IIpPoolRouteTableAssociationRef) IpPoolRouteTableAssociationRef() *IpPoolRouteTableAssociationReference {
	var returns *IpPoolRouteTableAssociationReference
	_jsii_.Get(
		j,
		"ipPoolRouteTableAssociationRef",
		&returns,
	)
	return returns
}

