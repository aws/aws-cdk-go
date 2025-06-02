package awssecretsmanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A hosted rotation.
//
// Example:
//   secret := secretsmanager.NewSecret(this, jsii.String("Secret"))
//
//   secret.addRotationSchedule(jsii.String("RotationSchedule"), &RotationScheduleOptions{
//   	HostedRotation: secretsmanager.HostedRotation_MysqlSingleUser(),
//   })
//
type HostedRotation interface {
	awsec2.IConnectable
	// Security group connections for this hosted rotation.
	Connections() awsec2.Connections
	// Binds this hosted rotation to a secret.
	Bind(secret ISecret, scope constructs.Construct) *CfnRotationSchedule_HostedRotationLambdaProperty
}

// The jsii proxy struct for HostedRotation
type jsiiProxy_HostedRotation struct {
	internal.Type__awsec2IConnectable
}

func (j *jsiiProxy_HostedRotation) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}


// MariaDB Multi User.
func HostedRotation_MariaDbMultiUser(options *MultiUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	if err := validateHostedRotation_MariaDbMultiUserParameters(options); err != nil {
		panic(err)
	}
	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"mariaDbMultiUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// MariaDB Single User.
func HostedRotation_MariaDbSingleUser(options *SingleUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	if err := validateHostedRotation_MariaDbSingleUserParameters(options); err != nil {
		panic(err)
	}
	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"mariaDbSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// MongoDB Multi User.
func HostedRotation_MongoDbMultiUser(options *MultiUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	if err := validateHostedRotation_MongoDbMultiUserParameters(options); err != nil {
		panic(err)
	}
	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"mongoDbMultiUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// MongoDB Single User.
func HostedRotation_MongoDbSingleUser(options *SingleUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	if err := validateHostedRotation_MongoDbSingleUserParameters(options); err != nil {
		panic(err)
	}
	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"mongoDbSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// MySQL Multi User.
func HostedRotation_MysqlMultiUser(options *MultiUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	if err := validateHostedRotation_MysqlMultiUserParameters(options); err != nil {
		panic(err)
	}
	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"mysqlMultiUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// MySQL Single User.
func HostedRotation_MysqlSingleUser(options *SingleUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	if err := validateHostedRotation_MysqlSingleUserParameters(options); err != nil {
		panic(err)
	}
	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"mysqlSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Oracle Multi User.
func HostedRotation_OracleMultiUser(options *MultiUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	if err := validateHostedRotation_OracleMultiUserParameters(options); err != nil {
		panic(err)
	}
	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"oracleMultiUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Oracle Single User.
func HostedRotation_OracleSingleUser(options *SingleUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	if err := validateHostedRotation_OracleSingleUserParameters(options); err != nil {
		panic(err)
	}
	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"oracleSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// PostgreSQL Multi User.
func HostedRotation_PostgreSqlMultiUser(options *MultiUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	if err := validateHostedRotation_PostgreSqlMultiUserParameters(options); err != nil {
		panic(err)
	}
	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"postgreSqlMultiUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// PostgreSQL Single User.
func HostedRotation_PostgreSqlSingleUser(options *SingleUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	if err := validateHostedRotation_PostgreSqlSingleUserParameters(options); err != nil {
		panic(err)
	}
	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"postgreSqlSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Redshift Multi User.
func HostedRotation_RedshiftMultiUser(options *MultiUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	if err := validateHostedRotation_RedshiftMultiUserParameters(options); err != nil {
		panic(err)
	}
	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"redshiftMultiUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Redshift Single User.
func HostedRotation_RedshiftSingleUser(options *SingleUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	if err := validateHostedRotation_RedshiftSingleUserParameters(options); err != nil {
		panic(err)
	}
	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"redshiftSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// SQL Server Multi User.
func HostedRotation_SqlServerMultiUser(options *MultiUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	if err := validateHostedRotation_SqlServerMultiUserParameters(options); err != nil {
		panic(err)
	}
	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"sqlServerMultiUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// SQL Server Single User.
func HostedRotation_SqlServerSingleUser(options *SingleUserHostedRotationOptions) HostedRotation {
	_init_.Initialize()

	if err := validateHostedRotation_SqlServerSingleUserParameters(options); err != nil {
		panic(err)
	}
	var returns HostedRotation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_secretsmanager.HostedRotation",
		"sqlServerSingleUser",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HostedRotation) Bind(secret ISecret, scope constructs.Construct) *CfnRotationSchedule_HostedRotationLambdaProperty {
	if err := h.validateBindParameters(secret, scope); err != nil {
		panic(err)
	}
	var returns *CfnRotationSchedule_HostedRotationLambdaProperty

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{secret, scope},
		&returns,
	)

	return returns
}

