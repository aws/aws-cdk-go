package awsapigateway


// Properties for defining a `CfnAccount`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccountProps := &CfnAccountProps{
//   	CloudWatchRoleArn: jsii.String("cloudWatchRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-account.html
//
type CfnAccountProps struct {
	// The ARN of an Amazon CloudWatch role for the current Account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-account.html#cfn-apigateway-account-cloudwatchrolearn
	//
	CloudWatchRoleArn *string `field:"optional" json:"cloudWatchRoleArn" yaml:"cloudWatchRoleArn"`
}

