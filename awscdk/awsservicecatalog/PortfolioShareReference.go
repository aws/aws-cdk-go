package awsservicecatalog


// A reference to a PortfolioShare resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portfolioShareReference := &PortfolioShareReference{
//   	PortfolioShareId: jsii.String("portfolioShareId"),
//   }
//
type PortfolioShareReference struct {
	// The Id of the PortfolioShare resource.
	PortfolioShareId *string `field:"required" json:"portfolioShareId" yaml:"portfolioShareId"`
}

