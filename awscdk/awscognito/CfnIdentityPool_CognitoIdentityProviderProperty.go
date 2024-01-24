package awscognito


// `CognitoIdentityProvider` is a property of the [AWS::Cognito::IdentityPool](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html) resource that represents an Amazon Cognito user pool and its client ID.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cognitoIdentityProviderProperty := &CognitoIdentityProviderProperty{
//   	ClientId: jsii.String("clientId"),
//   	ProviderName: jsii.String("providerName"),
//
//   	// the properties below are optional
//   	ServerSideTokenCheck: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitoidentityprovider.html
//
type CfnIdentityPool_CognitoIdentityProviderProperty struct {
	// The client ID for the Amazon Cognito user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitoidentityprovider.html#cfn-cognito-identitypool-cognitoidentityprovider-clientid
	//
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The provider name for an Amazon Cognito user pool.
	//
	// For example: `cognito-idp.us-east-2.amazonaws.com/us-east-2_123456789` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitoidentityprovider.html#cfn-cognito-identitypool-cognitoidentityprovider-providername
	//
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// TRUE if server-side token validation is enabled for the identity provider’s token.
	//
	// After you set the `ServerSideTokenCheck` to TRUE for an identity pool, that identity pool checks with the integrated user pools to make sure the user has not been globally signed out or deleted before the identity pool provides an OIDC token or AWS credentials for the user.
	//
	// If the user is signed out or deleted, the identity pool returns a 400 Not Authorized error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitoidentityprovider.html#cfn-cognito-identitypool-cognitoidentityprovider-serversidetokencheck
	//
	ServerSideTokenCheck interface{} `field:"optional" json:"serverSideTokenCheck" yaml:"serverSideTokenCheck"`
}

