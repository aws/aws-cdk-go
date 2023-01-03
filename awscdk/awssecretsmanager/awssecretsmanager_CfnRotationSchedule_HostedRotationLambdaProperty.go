package awssecretsmanager


// Creates a new Lambda rotation function based on one of the [Secrets Manager rotation function templates](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html) .
//
// You must specify `Transform: AWS::SecretsManager-2020-07-23` at the beginning of the CloudFormation template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hostedRotationLambdaProperty := &hostedRotationLambdaProperty{
//   	rotationType: jsii.String("rotationType"),
//
//   	// the properties below are optional
//   	excludeCharacters: jsii.String("excludeCharacters"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   	masterSecretArn: jsii.String("masterSecretArn"),
//   	masterSecretKmsKeyArn: jsii.String("masterSecretKmsKeyArn"),
//   	rotationLambdaName: jsii.String("rotationLambdaName"),
//   	superuserSecretArn: jsii.String("superuserSecretArn"),
//   	superuserSecretKmsKeyArn: jsii.String("superuserSecretKmsKeyArn"),
//   	vpcSecurityGroupIds: jsii.String("vpcSecurityGroupIds"),
//   	vpcSubnetIds: jsii.String("vpcSubnetIds"),
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
	// The ARN of the secret that contains elevated credentials.
	//
	// You must create the elevated secret before you can set this property. The Lambda rotation function uses this secret for the [Alternating users rotation strategy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets_strategies.html#rotating-secrets-two-users) .
	MasterSecretArn *string `field:"optional" json:"masterSecretArn" yaml:"masterSecretArn"`
	// The ARN of the KMS key that Secrets Manager uses to encrypt the elevated secret if you use the [alternating users strategy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets_strategies.html#rotating-secrets-two-users) . If you don't specify this value and you use the alternating users strategy, then Secrets Manager uses the key `aws/secretsmanager` . If `aws/secretsmanager` doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.
	MasterSecretKmsKeyArn *string `field:"optional" json:"masterSecretKmsKeyArn" yaml:"masterSecretKmsKeyArn"`
	// The name of the Lambda rotation function.
	RotationLambdaName *string `field:"optional" json:"rotationLambdaName" yaml:"rotationLambdaName"`
	// The ARN of the secret that contains elevated credentials.
	//
	// You must create the superuser secret before you can set this property. The Lambda rotation function uses this secret for the [Alternating users rotation strategy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets_strategies.html#rotating-secrets-two-users) .
	SuperuserSecretArn *string `field:"optional" json:"superuserSecretArn" yaml:"superuserSecretArn"`
	// The ARN of the KMS key that Secrets Manager uses to encrypt the elevated secret if you use the [alternating users strategy](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets_strategies.html#rotating-secrets-two-users) . If you don't specify this value and you use the alternating users strategy, then Secrets Manager uses the key `aws/secretsmanager` . If `aws/secretsmanager` doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.
	SuperuserSecretKmsKeyArn *string `field:"optional" json:"superuserSecretKmsKeyArn" yaml:"superuserSecretKmsKeyArn"`
	// A comma-separated list of security group IDs applied to the target database.
	//
	// The templates applies the same security groups as on the Lambda rotation function that is created as part of this stack.
	VpcSecurityGroupIds *string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
	// A comma separated list of VPC subnet IDs of the target database network.
	//
	// The Lambda rotation function is in the same subnet group.
	VpcSubnetIds *string `field:"optional" json:"vpcSubnetIds" yaml:"vpcSubnetIds"`
}

