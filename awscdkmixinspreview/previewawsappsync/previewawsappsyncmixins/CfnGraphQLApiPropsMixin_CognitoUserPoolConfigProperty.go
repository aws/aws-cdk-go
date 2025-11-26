package previewawsappsyncmixins


// Describes an Amazon Cognito user pool configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cognitoUserPoolConfigProperty := &CognitoUserPoolConfigProperty{
//   	AppIdClientRegex: jsii.String("appIdClientRegex"),
//   	AwsRegion: jsii.String("awsRegion"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-cognitouserpoolconfig.html
//
type CfnGraphQLApiPropsMixin_CognitoUserPoolConfigProperty struct {
	// A regular expression for validating the incoming Amazon Cognito user pool app client ID.
	//
	// If this value isn't set, no filtering is applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-cognitouserpoolconfig.html#cfn-appsync-graphqlapi-cognitouserpoolconfig-appidclientregex
	//
	AppIdClientRegex *string `field:"optional" json:"appIdClientRegex" yaml:"appIdClientRegex"`
	// The AWS Region in which the user pool was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-cognitouserpoolconfig.html#cfn-appsync-graphqlapi-cognitouserpoolconfig-awsregion
	//
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// The user pool ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-cognitouserpoolconfig.html#cfn-appsync-graphqlapi-cognitouserpoolconfig-userpoolid
	//
	UserPoolId *string `field:"optional" json:"userPoolId" yaml:"userPoolId"`
}

