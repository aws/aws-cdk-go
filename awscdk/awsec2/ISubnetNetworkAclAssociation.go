package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
)

// A SubnetNetworkAclAssociation.
type ISubnetNetworkAclAssociation interface {
	awscdk.IResource
	// ID for the current SubnetNetworkAclAssociation.
	SubnetNetworkAclAssociationAssociationId() *string
}

// The jsii proxy for ISubnetNetworkAclAssociation
type jsiiProxy_ISubnetNetworkAclAssociation struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ISubnetNetworkAclAssociation) SubnetNetworkAclAssociationAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetNetworkAclAssociationAssociationId",
		&returns,
	)
	return returns
}

