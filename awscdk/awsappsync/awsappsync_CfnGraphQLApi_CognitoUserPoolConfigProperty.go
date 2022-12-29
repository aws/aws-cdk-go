package awsappsync


// Describes an Amazon Cognito user pool configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cognitoUserPoolConfigProperty := &cognitoUserPoolConfigProperty{
//   	appIdClientRegex: jsii.String("appIdClientRegex"),
//   	awsRegion: jsii.String("awsRegion"),
//   	userPoolId: jsii.String("userPoolId"),
//   }
//
type CfnGraphQLApi_CognitoUserPoolConfigProperty struct {
	// A regular expression for validating the incoming Amazon Cognito user pool app client ID.
	AppIdClientRegex *string `field:"optional" json:"appIdClientRegex" yaml:"appIdClientRegex"`
	// The AWS Region in which the user pool was created.
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// The user pool ID.
	UserPoolId *string `field:"optional" json:"userPoolId" yaml:"userPoolId"`
}

