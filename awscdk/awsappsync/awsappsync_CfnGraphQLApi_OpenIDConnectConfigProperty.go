package awsappsync


// The `OpenIDConnectConfig` property type specifies the optional authorization configuration for using an OpenID Connect compliant service with your GraphQL endpoint for an AWS AppSync GraphQL API.
//
// `OpenIDConnectConfig` is a property of the [AWS::AppSync::GraphQLApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlapi.html) property type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openIDConnectConfigProperty := &openIDConnectConfigProperty{
//   	authTtl: jsii.Number(123),
//   	clientId: jsii.String("clientId"),
//   	iatTtl: jsii.Number(123),
//   	issuer: jsii.String("issuer"),
//   }
//
type CfnGraphQLApi_OpenIDConnectConfigProperty struct {
	// The number of milliseconds that a token is valid after being authenticated.
	AuthTtl *float64 `field:"optional" json:"authTtl" yaml:"authTtl"`
	// The client identifier of the Relying party at the OpenID identity provider.
	//
	// This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so that AWS AppSync can validate against multiple client identifiers at a time.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The number of milliseconds that a token is valid after it's issued to a user.
	IatTtl *float64 `field:"optional" json:"iatTtl" yaml:"iatTtl"`
	// The issuer for the OIDC configuration.
	//
	// The issuer returned by discovery must exactly match the value of `iss` in the ID token.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
}

