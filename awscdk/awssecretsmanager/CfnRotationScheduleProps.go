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
	// The ARN or name of the secret to rotate.
	//
	// To reference a secret also created in this template, use the [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html) function with the secret's logical ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-rotationschedule.html#cfn-secretsmanager-rotationschedule-secretid
	//
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// Creates a new Lambda rotation function based on one of the [Secrets Manager rotation function templates](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html) . To use a rotation function that already exists, specify `RotationLambdaARN` instead.
	//
	// For Amazon RDS master user credentials, see [AWS::RDS::DBCluster MasterUserSecret](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbcluster-masterusersecret.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-rotationschedule.html#cfn-secretsmanager-rotationschedule-hostedrotationlambda
	//
	HostedRotationLambda interface{} `field:"optional" json:"hostedRotationLambda" yaml:"hostedRotationLambda"`
	// Specifies whether to rotate the secret immediately or wait until the next scheduled rotation window.
	//
	// The rotation schedule is defined in `RotationRules` .
	//
	// If you don't immediately rotate the secret, Secrets Manager tests the rotation configuration by running the [`testSecret` step](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotate-secrets_how.html) of the Lambda rotation function. The test creates an `AWSPENDING` version of the secret and then removes it.
	//
	// If you don't specify this value, then by default, Secrets Manager rotates the secret immediately.
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
	// To create a new rotation function based on one of the [Secrets Manager rotation function templates](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html) , specify `HostedRotationLambda` instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-rotationschedule.html#cfn-secretsmanager-rotationschedule-rotationlambdaarn
	//
	RotationLambdaArn *string `field:"optional" json:"rotationLambdaArn" yaml:"rotationLambdaArn"`
	// A structure that defines the rotation configuration for this secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-rotationschedule.html#cfn-secretsmanager-rotationschedule-rotationrules
	//
	RotationRules interface{} `field:"optional" json:"rotationRules" yaml:"rotationRules"`
}

