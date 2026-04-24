package awscdkdsqlalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkdsqlalpha/v2/internal"
)

// Represents an Aurora DSQL cluster.
// Experimental.
type ICluster interface {
	awscdk.IResource
	// Grant the given identity the specified actions.
	// See: https://docs.aws.amazon.com/aurora-dsql/latest/userguide/authentication-authorization.html
	//
	// [disable-awslint:no-grants].
	//
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permission to connect to the database.
	//
	// [disable-awslint:no-grants].
	// Experimental.
	GrantConnect(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permission to connect to the database with admin role.
	//
	// [disable-awslint:no-grants].
	// Experimental.
	GrantConnectAdmin(grantee awsiam.IGrantable) awsiam.Grant
	// Arn of the cluster.
	// Experimental.
	ClusterArn() *string
	// Connection endpoint for the cluster.
	// Experimental.
	ClusterEndpoint() *string
	// Identifier of the cluster.
	// Experimental.
	ClusterIdentifier() *string
	// VPC endpoint service name for the cluster.
	// Experimental.
	VpcEndpointServiceName() *string
}

// The jsii proxy for ICluster
type jsiiProxy_ICluster struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ICluster) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ICluster) GrantConnect(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantConnectParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantConnect",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ICluster) GrantConnectAdmin(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantConnectAdminParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantConnectAdmin",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ICluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) VpcEndpointServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointServiceName",
		&returns,
	)
	return returns
}

