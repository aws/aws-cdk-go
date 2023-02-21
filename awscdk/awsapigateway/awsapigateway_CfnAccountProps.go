package awsapigateway


// Properties for defining a `CfnAccount`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccountProps := &cfnAccountProps{
//   	cloudWatchRoleArn: jsii.String("cloudWatchRoleArn"),
//   }
//
type CfnAccountProps struct {
	// The ARN of an Amazon CloudWatch role for the current Account.
	CloudWatchRoleArn *string `field:"optional" json:"cloudWatchRoleArn" yaml:"cloudWatchRoleArn"`
}

