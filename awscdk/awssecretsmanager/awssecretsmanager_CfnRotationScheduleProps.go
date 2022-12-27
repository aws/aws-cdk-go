package awssecretsmanager


// Properties for defining a `CfnRotationSchedule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRotationScheduleProps := &cfnRotationScheduleProps{
//   	secretId: jsii.String("secretId"),
//
//   	// the properties below are optional
//   	hostedRotationLambda: &hostedRotationLambdaProperty{
//   		rotationType: jsii.String("rotationType"),
//
//   		// the properties below are optional
//   		excludeCharacters: jsii.String("excludeCharacters"),
//   		kmsKeyArn: jsii.String("kmsKeyArn"),
//   		masterSecretArn: jsii.String("masterSecretArn"),
//   		masterSecretKmsKeyArn: jsii.String("masterSecretKmsKeyArn"),
//   		rotationLambdaName: jsii.String("rotationLambdaName"),
//   		superuserSecretArn: jsii.String("superuserSecretArn"),
//   		superuserSecretKmsKeyArn: jsii.String("superuserSecretKmsKeyArn"),
//   		vpcSecurityGroupIds: jsii.String("vpcSecurityGroupIds"),
//   		vpcSubnetIds: jsii.String("vpcSubnetIds"),
//   	},
//   	rotateImmediatelyOnUpdate: jsii.Boolean(false),
//   	rotationLambdaArn: jsii.String("rotationLambdaArn"),
//   	rotationRules: &rotationRulesProperty{
//   		automaticallyAfterDays: jsii.Number(123),
//   		duration: jsii.String("duration"),
//   		scheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   }
//
type CfnRotationScheduleProps struct {
	// The ARN or name of the secret to rotate.
	//
	// To reference a secret also created in this template, use the [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html) function with the secret's logical ID.
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// Creates a new Lambda rotation function based on one of the [Secrets Manager rotation function templates](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html) . To use a rotation function that already exists, specify `RotationLambdaARN` instead.
	HostedRotationLambda interface{} `field:"optional" json:"hostedRotationLambda" yaml:"hostedRotationLambda"`
	// Specifies whether to rotate the secret immediately or wait until the next scheduled rotation window.
	//
	// The rotation schedule is defined in `RotationRules` .
	//
	// If you don't immediately rotate the secret, Secrets Manager tests the rotation configuration by running the [`testSecret` step](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotate-secrets_how.html) of the Lambda rotation function. The test creates an `AWSPENDING` version of the secret and then removes it.
	//
	// If you don't specify this value, then by default, Secrets Manager rotates the secret immediately.
	RotateImmediatelyOnUpdate interface{} `field:"optional" json:"rotateImmediatelyOnUpdate" yaml:"rotateImmediatelyOnUpdate"`
	// The ARN of an existing Lambda rotation function.
	//
	// To specify a rotation function that is also defined in this template, use the [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html) function.
	//
	// To create a new rotation function based on one of the [Secrets Manager rotation function templates](https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_available-rotation-templates.html) , specify `HostedRotationLambda` instead.
	RotationLambdaArn *string `field:"optional" json:"rotationLambdaArn" yaml:"rotationLambdaArn"`
	// A structure that defines the rotation configuration for this secret.
	RotationRules interface{} `field:"optional" json:"rotationRules" yaml:"rotationRules"`
}

