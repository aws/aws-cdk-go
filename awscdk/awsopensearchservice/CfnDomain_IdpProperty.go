package awsopensearchservice


// The SAML Identity Provider's information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   idpProperty := &IdpProperty{
//   	EntityId: jsii.String("entityId"),
//   	MetadataContent: jsii.String("metadataContent"),
//   }
//
type CfnDomain_IdpProperty struct {
	// The unique entity ID of the application in the SAML identity provider.
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// The metadata of the SAML application, in XML format.
	MetadataContent *string `field:"required" json:"metadataContent" yaml:"metadataContent"`
}

