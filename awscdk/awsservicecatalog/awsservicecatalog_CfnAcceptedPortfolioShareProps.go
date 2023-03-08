package awsservicecatalog


// Properties for defining a `CfnAcceptedPortfolioShare`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAcceptedPortfolioShareProps := &cfnAcceptedPortfolioShareProps{
//   	portfolioId: jsii.String("portfolioId"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   }
//
type CfnAcceptedPortfolioShareProps struct {
	// The portfolio identifier.
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
}

