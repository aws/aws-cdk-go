package awssecretsmanager


// Creates a new Lambda rotation function based on one of the [Secrets Manager rotation function templates](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html) .
//
// You must specify `Transform: AWS::SecretsManager-2020-07-23` at the beginning of the CloudFormation template.
//
// For Amazon RDS master user credentials, see [AWS::RDS::DBCluster MasterUserSecret](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbcluster-masterusersecret.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostedRotationLambdaProperty := &HostedRotationLambdaProperty{
//   	RotationType: jsii.String("rotationType"),
//
//   	// the properties below are optional
//   	ExcludeCharacters: jsii.String("excludeCharacters"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	MasterSecretArn: jsii.String("masterSecretArn"),
//   	MasterSecretKmsKeyArn: jsii.String("masterSecretKmsKeyArn"),
//   	RotationLambdaName: jsii.String("rotationLambdaName"),
//   	Runtime: jsii.String("runtime"),
//   	SuperuserSecretArn: jsii.String("superuserSecretArn"),
//   	SuperuserSecretKmsKeyArn: jsii.String("superuserSecretKmsKeyArn"),
//   	VpcSecurityGroupIds: jsii.String("vpcSecurityGroupIds"),
//   	VpcSubnetIds: jsii.String("vpcSubnetIds"),
//   }
//
type CfnRotationSchedule_HostedRotationLambdaProperty struct {
	// The rotation template to base the rotation function on, one of the following:.
	//
	// - `MySQLSingleUser` to use the template [SecretsManagerRDSMySQLRotationSingleUser](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html#sar-template-mysql-singleuser) .
	// - `MySQLMultiUser` to use the template [SecretsManagerRDSMySQLRotationMultiUser](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html#sar-template-mysql-multiuser) .
	// - `PostgreSQLSingleUser` to use the template [SecretsManagerRDSPostgreSQLRotationSingleUser](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html#sar-template-postgre-singleuser)
	// - `PostgreSQLMultiUser` to use the template [SecretsManagerRDSPostgreSQLRotationMultiUser](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html#sar-template-postgre-multiuser) .
	// - `OracleSingleUser` to use the template [SecretsManagerRDSOracleRotationSingleUser](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html#sar-template-oracle-singleuser) .
	// - `OracleMultiUser` to use the template [SecretsManagerRDSOracleRotationMultiUser](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html#sar-template-oracle-multiuser) .
	// - `MariaDBSingleUser` to use the template [SecretsManagerRDSMariaDBRotationSingleUser](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html#sar-template-mariadb-singleuser) .
	// - `MariaDBMultiUser` to use the template [SecretsManagerRDSMariaDBRotationMultiUser](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html#sar-template-mariadb-multiuser) .
	// - `SQLServerSingleUser` to use the template [SecretsManagerRDSSQLServerRotationSingleUser](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html#sar-template-sqlserver-singleuser) .
	// - `SQLServerMultiUser` to use the template [SecretsManagerRDSSQLServerRotationMultiUser](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html#sar-template-sqlserver-multiuser) .
	// - `RedshiftSingleUser` to use the template [SecretsManagerRedshiftRotationSingleUsr](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html#sar-template-redshift-singleuser) .
	// - `RedshiftMultiUser` to use the template [SecretsManagerRedshiftRotationMultiUser](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html#sar-template-redshift-multiuser) .
	// - `MongoDBSingleUser` to use the template [SecretsManagerMongoDBRotationSingleUser](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html#sar-template-mongodb-singleuser) .
	// - `MongoDBMultiUser` to use the template [SecretsManagerMongoDBRotationMultiUser](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html#sar-template-mongodb-multiuser) .
	RotationType *string `field:"required" json:"rotationType" yaml:"rotationType"`
	// A string of the characters that you don't want in the password.
	ExcludeCharacters *string `field:"optional" json:"excludeCharacters" yaml:"excludeCharacters"`
	// The ARN of the KMS key that Secrets Manager uses to encrypt the secret.
	//
	// If you don't specify this value, then Secrets Manager uses the key `aws/secretsmanager` . If `aws/secretsmanager` doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The ARN of the secret that contains superuser credentials, if you use the [Alternating users rotation strategy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets_strategies.html#rotating-secrets-two-users) . CloudFormation grants the execution role for the Lambda rotation function `GetSecretValue` permission to the secret in this property. For more information, see [Lambda rotation function execution role permissions for Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets-required-permissions-function.html) .
	//
	// You must create the superuser secret before you can set this property.
	//
	// You must also include the superuser secret ARN as a key in the JSON of the rotating secret so that the Lambda rotation function can find it. CloudFormation does not hardcode secret ARNs in the Lambda rotation function, so you can use the function to rotate multiple secrets. For more information, see [JSON structure of Secrets Manager secrets](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_secret_json_structure.html) .
	//
	// You can specify `MasterSecretArn` or `SuperuserSecretArn` but not both. They represent the same superuser secret.
	MasterSecretArn *string `field:"optional" json:"masterSecretArn" yaml:"masterSecretArn"`
	// The ARN of the KMS key that Secrets Manager used to encrypt the superuser secret, if you use the [alternating users strategy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets_strategies.html#rotating-secrets-two-users) and the superuser secret is encrypted with a customer managed key. You don't need to specify this property if the superuser secret is encrypted using the key `aws/secretsmanager` . CloudFormation grants the execution role for the Lambda rotation function `Decrypt` , `DescribeKey` , and `GenerateDataKey` permission to the key in this property. For more information, see [Lambda rotation function execution role permissions for Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets-required-permissions-function.html) .
	//
	// You can specify `MasterSecretKmsKeyArn` or `SuperuserSecretKmsKeyArn` but not both. They represent the same superuser secret KMS key .
	MasterSecretKmsKeyArn *string `field:"optional" json:"masterSecretKmsKeyArn" yaml:"masterSecretKmsKeyArn"`
	// The name of the Lambda rotation function.
	RotationLambdaName *string `field:"optional" json:"rotationLambdaName" yaml:"rotationLambdaName"`
	// The Python runtime version associated with the Lambda function.
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// The ARN of the secret that contains superuser credentials, if you use the [Alternating users rotation strategy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets_strategies.html#rotating-secrets-two-users) . CloudFormation grants the execution role for the Lambda rotation function `GetSecretValue` permission to the secret in this property. For more information, see [Lambda rotation function execution role permissions for Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets-required-permissions-function.html) .
	//
	// You must create the superuser secret before you can set this property.
	//
	// You must also include the superuser secret ARN as a key in the JSON of the rotating secret so that the Lambda rotation function can find it. CloudFormation does not hardcode secret ARNs in the Lambda rotation function, so you can use the function to rotate multiple secrets. For more information, see [JSON structure of Secrets Manager secrets](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_secret_json_structure.html) .
	//
	// You can specify `MasterSecretArn` or `SuperuserSecretArn` but not both. They represent the same superuser secret.
	SuperuserSecretArn *string `field:"optional" json:"superuserSecretArn" yaml:"superuserSecretArn"`
	// The ARN of the KMS key that Secrets Manager used to encrypt the superuser secret, if you use the [alternating users strategy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets_strategies.html#rotating-secrets-two-users) and the superuser secret is encrypted with a customer managed key. You don't need to specify this property if the superuser secret is encrypted using the key `aws/secretsmanager` . CloudFormation grants the execution role for the Lambda rotation function `Decrypt` , `DescribeKey` , and `GenerateDataKey` permission to the key in this property. For more information, see [Lambda rotation function execution role permissions for Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets-required-permissions-function.html) .
	//
	// You can specify `MasterSecretKmsKeyArn` or `SuperuserSecretKmsKeyArn` but not both. They represent the same superuser secret KMS key .
	SuperuserSecretKmsKeyArn *string `field:"optional" json:"superuserSecretKmsKeyArn" yaml:"superuserSecretKmsKeyArn"`
	// A comma-separated list of security group IDs applied to the target database.
	//
	// The template applies the same security groups as on the Lambda rotation function that is created as part of this stack.
	VpcSecurityGroupIds *string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
	// A comma separated list of VPC subnet IDs of the target database network.
	//
	// The Lambda rotation function is in the same subnet group.
	VpcSubnetIds *string `field:"optional" json:"vpcSubnetIds" yaml:"vpcSubnetIds"`
}

