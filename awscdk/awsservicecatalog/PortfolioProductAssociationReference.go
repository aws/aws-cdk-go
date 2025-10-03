package awsservicecatalog


// A reference to a PortfolioProductAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portfolioProductAssociationReference := &PortfolioProductAssociationReference{
//   	PortfolioProductAssociationId: jsii.String("portfolioProductAssociationId"),
//   }
//
type PortfolioProductAssociationReference struct {
	// The Id of the PortfolioProductAssociation resource.
	PortfolioProductAssociationId *string `field:"required" json:"portfolioProductAssociationId" yaml:"portfolioProductAssociationId"`
}

