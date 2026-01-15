package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrds"
	"github.com/aws/constructs-go/constructs/v10"
)

// A DB proxy endpoint.
type IDatabaseProxyEndpoint interface {
	interfacesawsrds.IDBProxyEndpointRef
	awscdk.IResource
	// DB Proxy Endpoint ARN.
	DbProxyEndpointArn() *string
	// DB Proxy Endpoint Name.
	DbProxyEndpointName() *string
	// Endpoint.
	Endpoint() *string
}

// The jsii proxy for IDatabaseProxyEndpoint
type jsiiProxy_IDatabaseProxyEndpoint struct {
	internal.Type__interfacesawsrdsIDBProxyEndpointRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IDatabaseProxyEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IDatabaseProxyEndpoint) DbProxyEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbProxyEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseProxyEndpoint) DbProxyEndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbProxyEndpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseProxyEndpoint) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseProxyEndpoint) DbProxyEndpointRef() *interfacesawsrds.DBProxyEndpointReference {
	var returns *interfacesawsrds.DBProxyEndpointReference
	_jsii_.Get(
		j,
		"dbProxyEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseProxyEndpoint) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseProxyEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseProxyEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

