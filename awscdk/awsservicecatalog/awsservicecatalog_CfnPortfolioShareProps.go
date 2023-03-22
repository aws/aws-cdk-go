package awsservicecatalog


// Properties for defining a `CfnPortfolioShare`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPortfolioShareProps := &cfnPortfolioShareProps{
//   	accountId: jsii.String("accountId"),
//   	portfolioId: jsii.String("portfolioId"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	shareTagOptions: jsii.Boolean(false),
//   }
//
type CfnPortfolioShareProps struct {
	// The AWS account ID.
	//
	// For example, `123456789012` .
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The portfolio identifier.
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// Indicates whether TagOptions sharing is enabled or disabled for the portfolio share.
	ShareTagOptions interface{} `field:"optional" json:"shareTagOptions" yaml:"shareTagOptions"`
}

