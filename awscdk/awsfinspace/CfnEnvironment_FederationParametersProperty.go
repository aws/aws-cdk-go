package awsfinspace


// Configuration information when authentication mode is FEDERATED.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   federationParametersProperty := &FederationParametersProperty{
//   	ApplicationCallBackUrl: jsii.String("applicationCallBackUrl"),
//   	AttributeMap: []interface{}{
//   		&AttributeMapItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	FederationProviderName: jsii.String("federationProviderName"),
//   	FederationUrn: jsii.String("federationUrn"),
//   	SamlMetadataDocument: jsii.String("samlMetadataDocument"),
//   	SamlMetadataUrl: jsii.String("samlMetadataUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html
//
type CfnEnvironment_FederationParametersProperty struct {
	// The redirect or sign-in URL that should be entered into the SAML 2.0 compliant identity provider configuration (IdP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-applicationcallbackurl
	//
	ApplicationCallBackUrl *string `field:"optional" json:"applicationCallBackUrl" yaml:"applicationCallBackUrl"`
	// SAML attribute name and value.
	//
	// The name must always be `Email` and the value should be set to the attribute definition in which user email is set. For example, name would be `Email` and value `http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress` . Please check your SAML 2.0 compliant identity provider (IdP) documentation for details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-attributemap
	//
	AttributeMap interface{} `field:"optional" json:"attributeMap" yaml:"attributeMap"`
	// Name of the identity provider (IdP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-federationprovidername
	//
	FederationProviderName *string `field:"optional" json:"federationProviderName" yaml:"federationProviderName"`
	// The Uniform Resource Name (URN).
	//
	// Also referred as Service Provider URN or Audience URI or Service Provider Entity ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-federationurn
	//
	FederationUrn *string `field:"optional" json:"federationUrn" yaml:"federationUrn"`
	// SAML 2.0 Metadata document from identity provider (IdP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-samlmetadatadocument
	//
	SamlMetadataDocument *string `field:"optional" json:"samlMetadataDocument" yaml:"samlMetadataDocument"`
	// Provide the metadata URL from your SAML 2.0 compliant identity provider (IdP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-samlmetadataurl
	//
	SamlMetadataUrl *string `field:"optional" json:"samlMetadataUrl" yaml:"samlMetadataUrl"`
}

