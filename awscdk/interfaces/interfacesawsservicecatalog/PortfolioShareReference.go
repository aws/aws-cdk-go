package interfacesawsservicecatalog


// A reference to a PortfolioShare resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portfolioShareReference := &PortfolioShareReference{
//   	AccountId: jsii.String("accountId"),
//   	PortfolioId: jsii.String("portfolioId"),
//   }
//
type PortfolioShareReference struct {
	// The AccountId of the PortfolioShare resource.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The PortfolioId of the PortfolioShare resource.
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
}

