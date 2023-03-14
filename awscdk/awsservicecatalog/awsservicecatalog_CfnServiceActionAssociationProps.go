package awsservicecatalog


// Properties for defining a `CfnServiceActionAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServiceActionAssociationProps := &cfnServiceActionAssociationProps{
//   	productId: jsii.String("productId"),
//   	provisioningArtifactId: jsii.String("provisioningArtifactId"),
//   	serviceActionId: jsii.String("serviceActionId"),
//   }
//
type CfnServiceActionAssociationProps struct {
	// The product identifier.
	//
	// For example, `prod-abcdzk7xy33qa` .
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// The identifier of the provisioning artifact.
	//
	// For example, `pa-4abcdjnxjj6ne` .
	ProvisioningArtifactId *string `field:"required" json:"provisioningArtifactId" yaml:"provisioningArtifactId"`
	// The self-service action identifier.
	//
	// For example, `act-fs7abcd89wxyz` .
	ServiceActionId *string `field:"required" json:"serviceActionId" yaml:"serviceActionId"`
}

