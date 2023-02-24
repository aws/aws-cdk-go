package awsservicecatalog


// Properties for defining a `CfnPortfolioProductAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPortfolioProductAssociationProps := &CfnPortfolioProductAssociationProps{
//   	PortfolioId: jsii.String("portfolioId"),
//   	ProductId: jsii.String("productId"),
//
//   	// the properties below are optional
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   	SourcePortfolioId: jsii.String("sourcePortfolioId"),
//   }
//
type CfnPortfolioProductAssociationProps struct {
	// The portfolio identifier.
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
	// The product identifier.
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The identifier of the source portfolio.
	SourcePortfolioId *string `field:"optional" json:"sourcePortfolioId" yaml:"sourcePortfolioId"`
}

