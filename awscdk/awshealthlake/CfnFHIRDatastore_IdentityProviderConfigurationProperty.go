package awshealthlake


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-identityproviderconfiguration.html#cfn-healthlake-fhirdatastore-identityproviderconfiguration-authorizationstrategy
	//
	AuthorizationStrategy *string `field:"required" json:"authorizationStrategy" yaml:"authorizationStrategy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-identityproviderconfiguration.html#cfn-healthlake-fhirdatastore-identityproviderconfiguration-finegrainedauthorizationenabled
	//
	FineGrainedAuthorizationEnabled interface{} `field:"optional" json:"fineGrainedAuthorizationEnabled" yaml:"fineGrainedAuthorizationEnabled"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-identityproviderconfiguration.html#cfn-healthlake-fhirdatastore-identityproviderconfiguration-idplambdaarn
	//
	IdpLambdaArn *string `field:"optional" json:"idpLambdaArn" yaml:"idpLambdaArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-identityproviderconfiguration.html#cfn-healthlake-fhirdatastore-identityproviderconfiguration-metadata
	//
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
}

