package awssecretsmanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Hosted rotation type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostedRotationType := awscdk.Aws_secretsmanager.HostedRotationType_MARIADB_MULTI_USER()
//
type HostedRotationType interface {
	// Whether the rotation uses the mutli user scheme.
	IsMultiUser() *bool
	// The type of rotation.
	Name() *string
}

// The jsii proxy struct for HostedRotationType
type jsiiProxy_HostedRotationType struct {
	_ byte // padding
}

func (j *jsiiProxy_HostedRotationType) IsMultiUser() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isMultiUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HostedRotationType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func HostedRotationType_MARIADB_MULTI_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"MARIADB_MULTI_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_MARIADB_SINGLE_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"MARIADB_SINGLE_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_MONGODB_MULTI_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"MONGODB_MULTI_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_MONGODB_SINGLE_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"MONGODB_SINGLE_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_MYSQL_MULTI_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"MYSQL_MULTI_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_MYSQL_SINGLE_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"MYSQL_SINGLE_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_ORACLE_MULTI_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"ORACLE_MULTI_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_ORACLE_SINGLE_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"ORACLE_SINGLE_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_POSTGRESQL_MULTI_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"POSTGRESQL_MULTI_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_POSTGRESQL_SINGLE_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"POSTGRESQL_SINGLE_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_REDSHIFT_MULTI_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"REDSHIFT_MULTI_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_REDSHIFT_SINGLE_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"REDSHIFT_SINGLE_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_SQLSERVER_MULTI_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"SQLSERVER_MULTI_USER",
		&returns,
	)
	return returns
}

func HostedRotationType_SQLSERVER_SINGLE_USER() HostedRotationType {
	_init_.Initialize()
	var returns HostedRotationType
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.HostedRotationType",
		"SQLSERVER_SINGLE_USER",
		&returns,
	)
	return returns
}

