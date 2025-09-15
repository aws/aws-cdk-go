package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubnetNetworkAclAssociation.
// Experimental.
type ISubnetNetworkAclAssociationRef interface {
	constructs.IConstruct
	// A reference to a SubnetNetworkAclAssociation resource.
	// Experimental.
	SubnetNetworkAclAssociationRef() *SubnetNetworkAclAssociationReference
}

// The jsii proxy for ISubnetNetworkAclAssociationRef
type jsiiProxy_ISubnetNetworkAclAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISubnetNetworkAclAssociationRef) SubnetNetworkAclAssociationRef() *SubnetNetworkAclAssociationReference {
	var returns *SubnetNetworkAclAssociationReference
	_jsii_.Get(
		j,
		"subnetNetworkAclAssociationRef",
		&returns,
	)
	return returns
}

