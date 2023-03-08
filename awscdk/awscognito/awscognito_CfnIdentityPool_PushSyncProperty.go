package awscognito


// `PushSync` is a property of the [AWS::Cognito::IdentityPool](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html) resource that defines the configuration options to be applied to an Amazon Cognito identity pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pushSyncProperty := &pushSyncProperty{
//   	applicationArns: []*string{
//   		jsii.String("applicationArns"),
//   	},
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnIdentityPool_PushSyncProperty struct {
	// The ARNs of the Amazon SNS platform applications that could be used by clients.
	ApplicationArns *[]*string `field:"optional" json:"applicationArns" yaml:"applicationArns"`
	// An IAM role configured to allow Amazon Cognito to call Amazon SNS on behalf of the developer.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

