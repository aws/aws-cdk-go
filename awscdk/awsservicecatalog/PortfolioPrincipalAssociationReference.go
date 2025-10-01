package awsservicecatalog


// A reference to a PortfolioPrincipalAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portfolioPrincipalAssociationReference := &PortfolioPrincipalAssociationReference{
//   	PortfolioPrincipalAssociationId: jsii.String("portfolioPrincipalAssociationId"),
//   }
//
type PortfolioPrincipalAssociationReference struct {
	// The Id of the PortfolioPrincipalAssociation resource.
	PortfolioPrincipalAssociationId *string `field:"required" json:"portfolioPrincipalAssociationId" yaml:"portfolioPrincipalAssociationId"`
}

