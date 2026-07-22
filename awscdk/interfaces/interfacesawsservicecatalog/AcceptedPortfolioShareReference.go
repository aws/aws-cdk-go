package interfacesawsservicecatalog


// A reference to a AcceptedPortfolioShare resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acceptedPortfolioShareReference := &AcceptedPortfolioShareReference{
//   	PortfolioId: jsii.String("portfolioId"),
//   }
//
type AcceptedPortfolioShareReference struct {
	// The PortfolioId of the AcceptedPortfolioShare resource.
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
}

