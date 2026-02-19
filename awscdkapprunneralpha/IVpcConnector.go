package awscdkapprunneralpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdkapprunneralpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the App Runner VPC Connector.
// Experimental.
type IVpcConnector interface {
	awsec2.IConnectable
	awscdk.IResource
	// The ARN of the VPC connector.
	// Experimental.
	VpcConnectorArn() *string
	// The Name of the VPC connector.
	// Experimental.
	VpcConnectorName() *string
	// The revision of the VPC connector.
	// Experimental.
	VpcConnectorRevision() *float64
}

// The jsii proxy for IVpcConnector
type jsiiProxy_IVpcConnector struct {
	internal.Type__awsec2IConnectable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IVpcConnector) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IVpcConnector) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IVpcConnector) VpcConnectorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcConnector) VpcConnectorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcConnector) VpcConnectorRevision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vpcConnectorRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcConnector) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcConnector) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcConnector) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVpcConnector) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

