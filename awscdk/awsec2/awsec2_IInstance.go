package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2/internal"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Experimental.
type IInstance interface {
	IConnectable
	awsiam.IGrantable
	awscdk.IResource
	// The availability zone the instance was launched in.
	// Experimental.
	InstanceAvailabilityZone() *string
	// The instance's ID.
	// Experimental.
	InstanceId() *string
	// Private DNS name for this instance.
	// Experimental.
	InstancePrivateDnsName() *string
	// Private IP for this instance.
	// Experimental.
	InstancePrivateIp() *string
	// Publicly-routable DNS name for this instance.
	//
	// (May be an empty string if the instance does not have a public name).
	// Experimental.
	InstancePublicDnsName() *string
	// Publicly-routable IP  address for this instance.
	//
	// (May be an empty string if the instance does not have a public IP).
	// Experimental.
	InstancePublicIp() *string
}

// The jsii proxy for IInstance
type jsiiProxy_IInstance struct {
	jsiiProxy_IConnectable
	internal.Type__awsiamIGrantable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IInstance) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IInstance) InstanceAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstance) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstance) InstancePrivateDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePrivateDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstance) InstancePrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePrivateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstance) InstancePublicDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePublicDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstance) InstancePublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instancePublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstance) Connections() Connections {
	var returns Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstance) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstance) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstance) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstance) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

