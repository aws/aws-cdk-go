package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awsrds/internal"
)

// DB Proxy.
// Experimental.
type IDatabaseProxy interface {
	awscdk.IResource
	// Grant the given identity connection access to the proxy.
	// Experimental.
	GrantConnect(grantee awsiam.IGrantable, dbUser *string) awsiam.Grant
	// DB Proxy ARN.
	// Experimental.
	DbProxyArn() *string
	// DB Proxy Name.
	// Experimental.
	DbProxyName() *string
	// Endpoint.
	// Experimental.
	Endpoint() *string
}

// The jsii proxy for IDatabaseProxy
type jsiiProxy_IDatabaseProxy struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IDatabaseProxy) GrantConnect(grantee awsiam.IGrantable, dbUser *string) awsiam.Grant {
	if err := i.validateGrantConnectParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantConnect",
		[]interface{}{grantee, dbUser},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IDatabaseProxy) DbProxyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbProxyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseProxy) DbProxyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbProxyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseProxy) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

