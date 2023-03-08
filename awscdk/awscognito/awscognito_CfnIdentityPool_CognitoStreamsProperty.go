package awscognito


// `CognitoStreams` is a property of the [AWS::Cognito::IdentityPool](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html) resource that defines configuration options for Amazon Cognito streams.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cognitoStreamsProperty := &cognitoStreamsProperty{
//   	roleArn: jsii.String("roleArn"),
//   	streamingStatus: jsii.String("streamingStatus"),
//   	streamName: jsii.String("streamName"),
//   }
//
type CfnIdentityPool_CognitoStreamsProperty struct {
	// The Amazon Resource Name (ARN) of the role Amazon Cognito can assume to publish to the stream.
	//
	// This role must grant access to Amazon Cognito (cognito-sync) to invoke `PutRecord` on your Amazon Cognito stream.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Status of the Amazon Cognito streams.
	//
	// Valid values are: `ENABLED` or `DISABLED` .
	StreamingStatus *string `field:"optional" json:"streamingStatus" yaml:"streamingStatus"`
	// The name of the Amazon Cognito stream to receive updates.
	//
	// This stream must be in the developer's account and in the same Region as the identity pool.
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

