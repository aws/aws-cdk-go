package awshealthlake


// The identity provider configuration that you gave when the data store was created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identityProviderConfigurationProperty := &IdentityProviderConfigurationProperty{
//   	AuthorizationStrategy: jsii.String("authorizationStrategy"),
//
//   	// the properties below are optional
//   	FineGrainedAuthorizationEnabled: jsii.Boolean(false),
//   	IdpLambdaArn: jsii.String("idpLambdaArn"),
//   	Metadata: jsii.String("metadata"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-identityproviderconfiguration.html
//
type CfnFHIRDatastore_IdentityProviderConfigurationProperty struct {
	// The authorization strategy that you selected when you created the data store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-identityproviderconfiguration.html#cfn-healthlake-fhirdatastore-identityproviderconfiguration-authorizationstrategy
	//
	AuthorizationStrategy *string `field:"required" json:"authorizationStrategy" yaml:"authorizationStrategy"`
	// If you enabled fine-grained authorization when you created the data store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-identityproviderconfiguration.html#cfn-healthlake-fhirdatastore-identityproviderconfiguration-finegrainedauthorizationenabled
	//
	FineGrainedAuthorizationEnabled interface{} `field:"optional" json:"fineGrainedAuthorizationEnabled" yaml:"fineGrainedAuthorizationEnabled"`
	// The Amazon Resource Name (ARN) of the Lambda function that you want to use to decode the access token created by the authorization server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-identityproviderconfiguration.html#cfn-healthlake-fhirdatastore-identityproviderconfiguration-idplambdaarn
	//
	IdpLambdaArn *string `field:"optional" json:"idpLambdaArn" yaml:"idpLambdaArn"`
	// The JSON metadata elements that you want to use in your identity provider configuration.
	//
	// Required elements are listed based on the launch specification of the SMART application. For more information on all possible elements, see [Metadata](https://docs.aws.amazon.com/https://build.fhir.org/ig/HL7/smart-app-launch/conformance.html#metadata) in SMART's App Launch specification.
	//
	// `authorization_endpoint` : The URL to the OAuth2 authorization endpoint.
	//
	// `grant_types_supported` : An array of grant types that are supported at the token endpoint. You must provide at least one grant type option. Valid options are `authorization_code` and `client_credentials` .
	//
	// `token_endpoint` : The URL to the OAuth2 token endpoint.
	//
	// `capabilities` : An array of strings of the SMART capabilities that the authorization server supports.
	//
	// `code_challenge_methods_supported` : An array of strings of supported PKCE code challenge methods. You must include the `S256` method in the array of PKCE code challenge methods.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-identityproviderconfiguration.html#cfn-healthlake-fhirdatastore-identityproviderconfiguration-metadata
	//
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
}

