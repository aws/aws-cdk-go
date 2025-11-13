package interfacesawsservicecatalog


// A reference to a Portfolio resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portfolioReference := &PortfolioReference{
//   	PortfolioId: jsii.String("portfolioId"),
//   }
//
type PortfolioReference struct {
	// The Id of the Portfolio resource.
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
}

