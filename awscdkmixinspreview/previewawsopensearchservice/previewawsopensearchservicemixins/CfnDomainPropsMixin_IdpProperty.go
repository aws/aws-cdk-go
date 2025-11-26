package previewawsopensearchservicemixins


// The SAML Identity Provider's information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   idpProperty := &IdpProperty{
//   	EntityId: jsii.String("entityId"),
//   	MetadataContent: jsii.String("metadataContent"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-idp.html
//
type CfnDomainPropsMixin_IdpProperty struct {
	// The unique entity ID of the application in the SAML identity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-idp.html#cfn-opensearchservice-domain-idp-entityid
	//
	EntityId *string `field:"optional" json:"entityId" yaml:"entityId"`
	// The metadata of the SAML application, in XML format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-idp.html#cfn-opensearchservice-domain-idp-metadatacontent
	//
	MetadataContent *string `field:"optional" json:"metadataContent" yaml:"metadataContent"`
}

