package interfacesawsservicecatalog


// A reference to a AcceptedPortfolioShare resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acceptedPortfolioShareReference := &AcceptedPortfolioShareReference{
//   	AcceptedPortfolioShareId: jsii.String("acceptedPortfolioShareId"),
//   }
//
type AcceptedPortfolioShareReference struct {
	// The Id of the AcceptedPortfolioShare resource.
	AcceptedPortfolioShareId *string `field:"required" json:"acceptedPortfolioShareId" yaml:"acceptedPortfolioShareId"`
}

