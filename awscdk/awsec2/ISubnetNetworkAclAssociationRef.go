package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubnetNetworkAclAssociation.
// Experimental.
type ISubnetNetworkAclAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SubnetNetworkAclAssociation resource.
	// Experimental.
	SubnetNetworkAclAssociationRef() *SubnetNetworkAclAssociationReference
}

// The jsii proxy for ISubnetNetworkAclAssociationRef
type jsiiProxy_ISubnetNetworkAclAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ISubnetNetworkAclAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnetNetworkAclAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

