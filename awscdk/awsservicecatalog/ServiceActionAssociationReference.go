package awsservicecatalog


// A reference to a ServiceActionAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceActionAssociationReference := &ServiceActionAssociationReference{
//   	ProductId: jsii.String("productId"),
//   	ProvisioningArtifactId: jsii.String("provisioningArtifactId"),
//   	ServiceActionId: jsii.String("serviceActionId"),
//   }
//
type ServiceActionAssociationReference struct {
	// The ProductId of the ServiceActionAssociation resource.
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// The ProvisioningArtifactId of the ServiceActionAssociation resource.
	ProvisioningArtifactId *string `field:"required" json:"provisioningArtifactId" yaml:"provisioningArtifactId"`
	// The ServiceActionId of the ServiceActionAssociation resource.
	ServiceActionId *string `field:"required" json:"serviceActionId" yaml:"serviceActionId"`
}

