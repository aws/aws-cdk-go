package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A SubnetNetworkAclAssociation.
type ISubnetNetworkAclAssociation interface {
	awscdk.IResource
	ISubnetNetworkAclAssociationRef
	// ID for the current SubnetNetworkAclAssociation.
	SubnetNetworkAclAssociationAssociationId() *string
}

// The jsii proxy for ISubnetNetworkAclAssociation
type jsiiProxy_ISubnetNetworkAclAssociation struct {
	internal.Type__awscdkIResource
	jsiiProxy_ISubnetNetworkAclAssociationRef
}

func (i *jsiiProxy_ISubnetNetworkAclAssociation) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
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

func (j *jsiiProxy_ISubnetNetworkAclAssociation) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnetNetworkAclAssociation) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnetNetworkAclAssociation) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnetNetworkAclAssociation) SubnetNetworkAclAssociationRef() *SubnetNetworkAclAssociationReference {
	var returns *SubnetNetworkAclAssociationReference
	_jsii_.Get(
		j,
		"subnetNetworkAclAssociationRef",
		&returns,
	)
	return returns
}

