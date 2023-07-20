package awscognito


// `CognitoStreams` is a property of the [AWS::Cognito::IdentityPool](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html) resource that defines configuration options for Amazon Cognito streams.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cognitoStreamsProperty := &CognitoStreamsProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	StreamingStatus: jsii.String("streamingStatus"),
//   	StreamName: jsii.String("streamName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitostreams.html
//
type CfnIdentityPool_CognitoStreamsProperty struct {
	// The Amazon Resource Name (ARN) of the role Amazon Cognito can assume to publish to the stream.
	//
	// This role must grant access to Amazon Cognito (cognito-sync) to invoke `PutRecord` on your Amazon Cognito stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitostreams.html#cfn-cognito-identitypool-cognitostreams-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Status of the Amazon Cognito streams.
	//
	// Valid values are: `ENABLED` or `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitostreams.html#cfn-cognito-identitypool-cognitostreams-streamingstatus
	//
	StreamingStatus *string `field:"optional" json:"streamingStatus" yaml:"streamingStatus"`
	// The name of the Amazon Cognito stream to receive updates.
	//
	// This stream must be in the developer's account and in the same Region as the identity pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitostreams.html#cfn-cognito-identitypool-cognitostreams-streamname
	//
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
}

