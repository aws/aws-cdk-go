package awsapigatewayv2


// The `JWTConfiguration` property specifies the configuration of a JWT authorizer.
//
// Required for the `JWT` authorizer type. Supported only for HTTP APIs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jWTConfigurationProperty := &jWTConfigurationProperty{
//   	audience: []*string{
//   		jsii.String("audience"),
//   	},
//   	issuer: jsii.String("issuer"),
//   }
//
type CfnAuthorizer_JWTConfigurationProperty struct {
	// A list of the intended recipients of the JWT.
	//
	// A valid JWT must provide an `aud` that matches at least one entry in this list. See [RFC 7519](https://docs.aws.amazon.com/https://tools.ietf.org/html/rfc7519#section-4.1.3) . Required for the `JWT` authorizer type. Supported only for HTTP APIs.
	Audience *[]*string `field:"optional" json:"audience" yaml:"audience"`
	// The base domain of the identity provider that issues JSON Web Tokens.
	//
	// For example, an Amazon Cognito user pool has the following format: `https://cognito-idp. {region} .amazonaws.com/ {userPoolId}` . Required for the `JWT` authorizer type. Supported only for HTTP APIs.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
}

