package awssecretsmanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A secret rotation serverless application.
//
// Example:
//   var myUserSecret secret
//   var myMasterSecret secret
//   var myDatabase iConnectable
//   var myVpc vpc
//
//
//   secretsmanager.NewSecretRotation(this, jsii.String("SecretRotation"), &SecretRotationProps{
//   	Application: secretsmanager.SecretRotationApplication_MYSQL_ROTATION_MULTI_USER(),
//   	Secret: myUserSecret,
//   	 // The secret that will be rotated
//   	MasterSecret: myMasterSecret,
//   	 // The secret used for the rotation
//   	Target: myDatabase,
//   	Vpc: myVpc,
//   })
//
type SecretRotationApplication interface {
	// Whether the rotation application uses the mutli user scheme.
	IsMultiUser() *bool
	// Returns the application ARN for the current partition.
	//
	// Can be used in combination with a `CfnMapping` to automatically select the correct ARN based on the current partition.
	ApplicationArnForPartition(partition *string) *string
	// The semantic version of the app for the current partition.
	//
	// Can be used in combination with a `CfnMapping` to automatically select the correct version based on the current partition.
	SemanticVersionForPartition(partition *string) *string
}

// The jsii proxy struct for SecretRotationApplication
type jsiiProxy_SecretRotationApplication struct {
	_ byte // padding
}

func (j *jsiiProxy_SecretRotationApplication) IsMultiUser() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isMultiUser",
		&returns,
	)
	return returns
}


func NewSecretRotationApplication(applicationId *string, semanticVersion *string, options *SecretRotationApplicationOptions) SecretRotationApplication {
	_init_.Initialize()

	if err := validateNewSecretRotationApplicationParameters(applicationId, semanticVersion, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretRotationApplication{}

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		[]interface{}{applicationId, semanticVersion, options},
		&j,
	)

	return &j
}

func NewSecretRotationApplication_Override(s SecretRotationApplication, applicationId *string, semanticVersion *string, options *SecretRotationApplicationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		[]interface{}{applicationId, semanticVersion, options},
		s,
	)
}

func SecretRotationApplication_MARIADB_ROTATION_MULTI_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"MARIADB_ROTATION_MULTI_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_MARIADB_ROTATION_SINGLE_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"MARIADB_ROTATION_SINGLE_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_MONGODB_ROTATION_MULTI_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"MONGODB_ROTATION_MULTI_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_MONGODB_ROTATION_SINGLE_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"MONGODB_ROTATION_SINGLE_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_MYSQL_ROTATION_MULTI_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"MYSQL_ROTATION_MULTI_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_MYSQL_ROTATION_SINGLE_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"MYSQL_ROTATION_SINGLE_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_ORACLE_ROTATION_MULTI_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"ORACLE_ROTATION_MULTI_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_ORACLE_ROTATION_SINGLE_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"ORACLE_ROTATION_SINGLE_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_POSTGRES_ROTATION_MULTI_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"POSTGRES_ROTATION_MULTI_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_POSTGRES_ROTATION_SINGLE_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"POSTGRES_ROTATION_SINGLE_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_REDSHIFT_ROTATION_MULTI_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"REDSHIFT_ROTATION_MULTI_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_REDSHIFT_ROTATION_SINGLE_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"REDSHIFT_ROTATION_SINGLE_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_SQLSERVER_ROTATION_MULTI_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"SQLSERVER_ROTATION_MULTI_USER",
		&returns,
	)
	return returns
}

func SecretRotationApplication_SQLSERVER_ROTATION_SINGLE_USER() SecretRotationApplication {
	_init_.Initialize()
	var returns SecretRotationApplication
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_secretsmanager.SecretRotationApplication",
		"SQLSERVER_ROTATION_SINGLE_USER",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SecretRotationApplication) ApplicationArnForPartition(partition *string) *string {
	if err := s.validateApplicationArnForPartitionParameters(partition); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"applicationArnForPartition",
		[]interface{}{partition},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretRotationApplication) SemanticVersionForPartition(partition *string) *string {
	if err := s.validateSemanticVersionForPartitionParameters(partition); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"semanticVersionForPartition",
		[]interface{}{partition},
		&returns,
	)

	return returns
}

