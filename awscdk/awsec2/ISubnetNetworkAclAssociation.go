package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

// A SubnetNetworkAclAssociation.
type ISubnetNetworkAclAssociation interface {
	awscdk.IResource
	interfacesawsec2.ISubnetNetworkAclAssociationRef
	// ID for the current SubnetNetworkAclAssociation.
	SubnetNetworkAclAssociationAssociationId() *string
}

// The jsii proxy for ISubnetNetworkAclAssociation
type jsiiProxy_ISubnetNetworkAclAssociation struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsec2ISubnetNetworkAclAssociationRef
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

func (i *jsiiProxy_ISubnetNetworkAclAssociation) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISubnetNetworkAclAssociation) SubnetNetworkAclAssociationAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetNetworkAclAssociationAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubnetNetworkAclAssociation) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

func (j *jsiiProxy_ISubnetNetworkAclAssociation) SubnetNetworkAclAssociationRef() *interfacesawsec2.SubnetNetworkAclAssociationReference {
	var returns *interfacesawsec2.SubnetNetworkAclAssociationReference
	_jsii_.Get(
		j,
		"subnetNetworkAclAssociationRef",
		&returns,
	)
	return returns
}

