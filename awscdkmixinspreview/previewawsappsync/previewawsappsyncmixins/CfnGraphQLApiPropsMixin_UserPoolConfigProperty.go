package previewawsappsyncmixins


// The `UserPoolConfig` property type specifies the optional authorization configuration for using Amazon Cognito user pools with your GraphQL endpoint for an AWS AppSync GraphQL API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   userPoolConfigProperty := &UserPoolConfigProperty{
//   	AppIdClientRegex: jsii.String("appIdClientRegex"),
//   	AwsRegion: jsii.String("awsRegion"),
//   	DefaultAction: jsii.String("defaultAction"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-userpoolconfig.html
//
type CfnGraphQLApiPropsMixin_UserPoolConfigProperty struct {
	// A regular expression for validating the incoming Amazon Cognito user pool app client ID.
	//
	// If this value isn't set, no filtering is applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-userpoolconfig.html#cfn-appsync-graphqlapi-userpoolconfig-appidclientregex
	//
	AppIdClientRegex *string `field:"optional" json:"appIdClientRegex" yaml:"appIdClientRegex"`
	// The AWS Region in which the user pool was created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-userpoolconfig.html#cfn-appsync-graphqlapi-userpoolconfig-awsregion
	//
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// The action that you want your GraphQL API to take when a request that uses Amazon Cognito user pool authentication doesn't match the Amazon Cognito user pool configuration.
	//
	// When specifying Amazon Cognito user pools as the default authentication, you must set the value for `DefaultAction` to `ALLOW` if specifying `AdditionalAuthenticationProviders` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-userpoolconfig.html#cfn-appsync-graphqlapi-userpoolconfig-defaultaction
	//
	DefaultAction *string `field:"optional" json:"defaultAction" yaml:"defaultAction"`
	// The user pool ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-userpoolconfig.html#cfn-appsync-graphqlapi-userpoolconfig-userpoolid
	//
	UserPoolId *string `field:"optional" json:"userPoolId" yaml:"userPoolId"`
}

