package awssecretsmanager


// Properties for defining a `CfnRotationSchedule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRotationScheduleProps := &CfnRotationScheduleProps{
//   	SecretId: jsii.String("secretId"),
//
//   	// the properties below are optional
//   	ExternalSecretRotationMetadata: []interface{}{
//   		&ExternalSecretRotationMetadataItemProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ExternalSecretRotationRoleArn: jsii.String("externalSecretRotationRoleArn"),
//   	HostedRotationLambda: &HostedRotationLambdaProperty{
//   		RotationType: jsii.String("rotationType"),
//
//   		// the properties below are optional
//   		ExcludeCharacters: jsii.String("excludeCharacters"),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		MasterSecretArn: jsii.String("masterSecretArn"),
//   		MasterSecretKmsKeyArn: jsii.String("masterSecretKmsKeyArn"),
//   		RotationLambdaName: jsii.String("rotationLambdaName"),
//   		Runtime: jsii.String("runtime"),
//   		SuperuserSecretArn: jsii.String("superuserSecretArn"),
//   		SuperuserSecretKmsKeyArn: jsii.String("superuserSecretKmsKeyArn"),
//   		VpcSecurityGroupIds: jsii.String("vpcSecurityGroupIds"),
//   		VpcSubnetIds: jsii.String("vpcSubnetIds"),
//   	},
//   	RotateImmediatelyOnUpdate: jsii.Boolean(false),
//   	RotationLambdaArn: jsii.String("rotationLambdaArn"),
//   	RotationRules: &RotationRulesProperty{
//   		AutomaticallyAfterDays: jsii.Number(123),
//   		Duration: jsii.String("duration"),
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-rotationschedule.html
//
type CfnRotationScheduleProps struct {
	// The ARN or name of the secret to rotate. This is unique for each rotation schedule definition.
	//
	// To reference a secret also created in this template, use the [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html) function with the secret's logical ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-rotationschedule.html#cfn-secretsmanager-rotationschedule-secretid
	//
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// The list of metadata needed to successfully rotate a managed external secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-rotationschedule.html#cfn-secretsmanager-rotationschedule-externalsecretrotationmetadata
	//
	ExternalSecretRotationMetadata interface{} `field:"optional" json:"externalSecretRotationMetadata" yaml:"externalSecretRotationMetadata"`
	// The ARN of the IAM role that is used by Secrets Manager to rotate a managed external secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-rotationschedule.html#cfn-secretsmanager-rotationschedule-externalsecretrotationrolearn
	//
	ExternalSecretRotationRoleArn *string `field:"optional" json:"externalSecretRotationRoleArn" yaml:"externalSecretRotationRoleArn"`
	// Creates a new Lambda rotation function based on one of the [Secrets Manager rotation function templates](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html) . To use a rotation function that already exists, specify `RotationLambdaARN` instead.
	//
	// You must specify `Transform: AWS::SecretsManager-2024-09-16` at the beginning of the CloudFormation template. Transforms are macros hosted by AWS CloudFormation that help you create and manage complex infrastructure. The `Transform: AWS::SecretsManager-2024-09-16` transform automatically extends the CloudFormation stack to include a nested stack (of type `AWS::CloudFormation::Stack` ), which then creates and updates on your behalf during subsequent stack operations, the appropriate rotation Lambda function for your database or service. For general information on transforms, see the [AWS CloudFormation documentation.](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-reference.html)
	//
	// For Amazon RDS master user credentials, see [AWS::RDS::DBCluster MasterUserSecret](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbcluster-masterusersecret.html) .
	//
	// For Amazon Redshift admin user credentials, see [AWS::Redshift::Cluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-rotationschedule.html#cfn-secretsmanager-rotationschedule-hostedrotationlambda
	//
	HostedRotationLambda interface{} `field:"optional" json:"hostedRotationLambda" yaml:"hostedRotationLambda"`
	// Determines whether to rotate the secret immediately or wait until the next scheduled rotation window when the rotation schedule is updated.
	//
	// The rotation schedule is defined in `RotationRules` .
	//
	// The default for `RotateImmediatelyOnUpdate` is `true` . If you don't specify this value, Secrets Manager rotates the secret immediately.
	//
	// If you set `RotateImmediatelyOnUpdate` to `false` , Secrets Manager tests the rotation configuration by running the [`testSecret` step](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotate-secrets_how.html) of the Lambda rotation function. This test creates an `AWSPENDING` version of the secret and then removes it.
	//
	// > When changing an existing rotation schedule and setting `RotateImmediatelyOnUpdate` to `false` :
	// >
	// > - If using `AutomaticallyAfterDays` or a `ScheduleExpression` with `rate()` , the previously scheduled rotation might still occur.
	// > - To prevent unintended rotations, use a `ScheduleExpression` with `cron()` for granular control over rotation windows.
	//
	// Rotation is an asynchronous process. For more information, see [How rotation works](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotate-secrets_how.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-rotationschedule.html#cfn-secretsmanager-rotationschedule-rotateimmediatelyonupdate
	//
	RotateImmediatelyOnUpdate interface{} `field:"optional" json:"rotateImmediatelyOnUpdate" yaml:"rotateImmediatelyOnUpdate"`
	// The ARN of an existing Lambda rotation function.
	//
	// To specify a rotation function that is also defined in this template, use the [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html) function.
	//
	// For Amazon RDS master user credentials, see [AWS::RDS::DBCluster MasterUserSecret](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbcluster-masterusersecret.html) .
	//
	// For Amazon Redshift admin user credentials, see [AWS::Redshift::Cluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-cluster.html) .
	//
	// To create a new rotation function based on one of the [Secrets Manager rotation function templates](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html) , specify `HostedRotationLambda` instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-rotationschedule.html#cfn-secretsmanager-rotationschedule-rotationlambdaarn
	//
	RotationLambdaArn *string `field:"optional" json:"rotationLambdaArn" yaml:"rotationLambdaArn"`
	// A structure that defines the rotation configuration for this secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-rotationschedule.html#cfn-secretsmanager-rotationschedule-rotationrules
	//
	RotationRules interface{} `field:"optional" json:"rotationRules" yaml:"rotationRules"`
}

