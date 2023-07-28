package awshealthlake


// The identity provider configuration for the datastore.
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
	// Type of Authorization Strategy.
	//
	// The two types of supported Authorization strategies are SMART_ON_FHIR_V1 and AWS_AUTH.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-identityproviderconfiguration.html#cfn-healthlake-fhirdatastore-identityproviderconfiguration-authorizationstrategy
	//
	AuthorizationStrategy *string `field:"required" json:"authorizationStrategy" yaml:"authorizationStrategy"`
	// Flag to indicate if fine-grained authorization will be enabled for the datastore.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-identityproviderconfiguration.html#cfn-healthlake-fhirdatastore-identityproviderconfiguration-finegrainedauthorizationenabled
	//
	FineGrainedAuthorizationEnabled interface{} `field:"optional" json:"fineGrainedAuthorizationEnabled" yaml:"fineGrainedAuthorizationEnabled"`
	// The Amazon Resource Name (ARN) of the Lambda function that will be used to decode the access token created by the authorization server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-identityproviderconfiguration.html#cfn-healthlake-fhirdatastore-identityproviderconfiguration-idplambdaarn
	//
	IdpLambdaArn *string `field:"optional" json:"idpLambdaArn" yaml:"idpLambdaArn"`
	// The JSON metadata elements for identity provider configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-identityproviderconfiguration.html#cfn-healthlake-fhirdatastore-identityproviderconfiguration-metadata
	//
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
}

