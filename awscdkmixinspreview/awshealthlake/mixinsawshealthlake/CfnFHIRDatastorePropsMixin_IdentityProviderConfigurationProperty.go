package mixinsawshealthlake


// The identity provider configuration selected when the data store was created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   identityProviderConfigurationProperty := &IdentityProviderConfigurationProperty{
//   	AuthorizationStrategy: jsii.String("authorizationStrategy"),
//   	FineGrainedAuthorizationEnabled: jsii.Boolean(false),
//   	IdpLambdaArn: jsii.String("idpLambdaArn"),
//   	Metadata: jsii.String("metadata"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-identityproviderconfiguration.html
//
type CfnFHIRDatastorePropsMixin_IdentityProviderConfigurationProperty struct {
	// The authorization strategy selected when the HealthLake data store is created.
	//
	// > HealthLake provides support for both SMART on FHIR V1 and V2 as described below.
	// >
	// > - `SMART_ON_FHIR_V1` – Support for only SMART on FHIR V1, which includes `read` (read/search) and `write` (create/update/delete) permissions.
	// > - `SMART_ON_FHIR` – Support for both SMART on FHIR V1 and V2, which includes `create` , `read` , `update` , `delete` , and `search` permissions.
	// > - `AWS_AUTH` – The default HealthLake authorization strategy; not affiliated with SMART on FHIR.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-identityproviderconfiguration.html#cfn-healthlake-fhirdatastore-identityproviderconfiguration-authorizationstrategy
	//
	AuthorizationStrategy *string `field:"optional" json:"authorizationStrategy" yaml:"authorizationStrategy"`
	// The parameter to enable SMART on FHIR fine-grained authorization for the data store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-identityproviderconfiguration.html#cfn-healthlake-fhirdatastore-identityproviderconfiguration-finegrainedauthorizationenabled
	//
	FineGrainedAuthorizationEnabled interface{} `field:"optional" json:"fineGrainedAuthorizationEnabled" yaml:"fineGrainedAuthorizationEnabled"`
	// The Amazon Resource Name (ARN) of the Lambda function to use to decode the access token created by the authorization server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-identityproviderconfiguration.html#cfn-healthlake-fhirdatastore-identityproviderconfiguration-idplambdaarn
	//
	IdpLambdaArn *string `field:"optional" json:"idpLambdaArn" yaml:"idpLambdaArn"`
	// The JSON metadata elements to use in your identity provider configuration.
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

