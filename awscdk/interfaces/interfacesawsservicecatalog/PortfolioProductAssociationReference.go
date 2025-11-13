package interfacesawsservicecatalog


// A reference to a PortfolioProductAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portfolioProductAssociationReference := &PortfolioProductAssociationReference{
//   	PortfolioId: jsii.String("portfolioId"),
//   	ProductId: jsii.String("productId"),
//   }
//
type PortfolioProductAssociationReference struct {
	// The PortfolioId of the PortfolioProductAssociation resource.
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
	// The ProductId of the PortfolioProductAssociation resource.
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
}

