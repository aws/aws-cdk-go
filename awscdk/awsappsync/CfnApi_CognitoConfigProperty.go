package awsappsync


// Describes an Amazon Cognito configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cognitoConfigProperty := &CognitoConfigProperty{
//   	AwsRegion: jsii.String("awsRegion"),
//   	UserPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	AppIdClientRegex: jsii.String("appIdClientRegex"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-cognitoconfig.html
//
type CfnApi_CognitoConfigProperty struct {
	// The AWS Region in which the user pool was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-cognitoconfig.html#cfn-appsync-api-cognitoconfig-awsregion
	//
	AwsRegion *string `field:"required" json:"awsRegion" yaml:"awsRegion"`
	// The user pool ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-cognitoconfig.html#cfn-appsync-api-cognitoconfig-userpoolid
	//
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// A regular expression for validating the incoming Amazon Cognito user pool app client ID.
	//
	// If this value isn't set, no filtering is applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-api-cognitoconfig.html#cfn-appsync-api-cognitoconfig-appidclientregex
	//
	AppIdClientRegex *string `field:"optional" json:"appIdClientRegex" yaml:"appIdClientRegex"`
}

