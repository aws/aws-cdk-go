package awssecretsmanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Options to add a rotation schedule to a secret.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var fn function
//
//   secret := secretsmanager.NewSecret(this, jsii.String("Secret"))
//
//   secret.addRotationSchedule(jsii.String("RotationSchedule"), &RotationScheduleOptions{
//   	RotationLambda: fn,
//   	AutomaticallyAfter: awscdk.Duration_Days(jsii.Number(15)),
//   	RotateImmediatelyOnUpdate: jsii.Boolean(false),
//   })
//
type RotationScheduleOptions struct {
	// Specifies the number of days after the previous rotation before Secrets Manager triggers the next automatic rotation.
	//
	// The minimum value is 4 hours.
	// The maximum value is 1000 days.
	//
	// A value of zero (`Duration.days(0)`) will not create RotationRules.
	// Default: Duration.days(30)
	//
	AutomaticallyAfter awscdk.Duration `field:"optional" json:"automaticallyAfter" yaml:"automaticallyAfter"`
	// Hosted rotation.
	// Default: - either `rotationLambda` or `hostedRotation` must be specified.
	//
	HostedRotation HostedRotation `field:"optional" json:"hostedRotation" yaml:"hostedRotation"`
	// Specifies whether to rotate the secret immediately or wait until the next scheduled rotation window.
	// Default: true.
	//
	RotateImmediatelyOnUpdate *bool `field:"optional" json:"rotateImmediatelyOnUpdate" yaml:"rotateImmediatelyOnUpdate"`
	// A Lambda function that can rotate the secret.
	// Default: - either `rotationLambda` or `hostedRotation` must be specified.
	//
	RotationLambda awslambda.IFunction `field:"optional" json:"rotationLambda" yaml:"rotationLambda"`
}

