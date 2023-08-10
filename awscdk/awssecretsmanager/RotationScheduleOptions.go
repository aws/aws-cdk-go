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
	// A value of zero will disable automatic rotation - `Duration.days(0)`.
	AutomaticallyAfter awscdk.Duration `field:"optional" json:"automaticallyAfter" yaml:"automaticallyAfter"`
	// Hosted rotation.
	HostedRotation HostedRotation `field:"optional" json:"hostedRotation" yaml:"hostedRotation"`
	// Specifies whether to rotate the secret immediately or wait until the next scheduled rotation window.
	RotateImmediatelyOnUpdate *bool `field:"optional" json:"rotateImmediatelyOnUpdate" yaml:"rotateImmediatelyOnUpdate"`
	// A Lambda function that can rotate the secret.
	RotationLambda awslambda.IFunction `field:"optional" json:"rotationLambda" yaml:"rotationLambda"`
}

