package awsfinspace


// Configuration information when authentication mode is FEDERATED.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributeMap interface{}
//
//   federationParametersProperty := &federationParametersProperty{
//   	applicationCallBackUrl: jsii.String("applicationCallBackUrl"),
//   	attributeMap: attributeMap,
//   	federationProviderName: jsii.String("federationProviderName"),
//   	federationUrn: jsii.String("federationUrn"),
//   	samlMetadataDocument: jsii.String("samlMetadataDocument"),
//   	samlMetadataUrl: jsii.String("samlMetadataUrl"),
//   }
//
type CfnEnvironment_FederationParametersProperty struct {
	// The redirect or sign-in URL that should be entered into the SAML 2.0 compliant identity provider configuration (IdP).
	ApplicationCallBackUrl *string `field:"optional" json:"applicationCallBackUrl" yaml:"applicationCallBackUrl"`
	// SAML attribute name and value.
	//
	// The name must always be `Email` and the value should be set to the attribute definition in which user email is set. For example, name would be `Email` and value `http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress` . Please check your SAML 2.0 compliant identity provider (IdP) documentation for details.
	AttributeMap interface{} `field:"optional" json:"attributeMap" yaml:"attributeMap"`
	// Name of the identity provider (IdP).
	FederationProviderName *string `field:"optional" json:"federationProviderName" yaml:"federationProviderName"`
	// The Uniform Resource Name (URN).
	//
	// Also referred as Service Provider URN or Audience URI or Service Provider Entity ID.
	FederationUrn *string `field:"optional" json:"federationUrn" yaml:"federationUrn"`
	// SAML 2.0 Metadata document from identity provider (IdP).
	SamlMetadataDocument *string `field:"optional" json:"samlMetadataDocument" yaml:"samlMetadataDocument"`
	// Provide the metadata URL from your SAML 2.0 compliant identity provider (IdP).
	SamlMetadataUrl *string `field:"optional" json:"samlMetadataUrl" yaml:"samlMetadataUrl"`
}

