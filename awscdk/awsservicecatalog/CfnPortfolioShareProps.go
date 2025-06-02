package awsservicecatalog


// Properties for defining a `CfnPortfolioShare`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPortfolioShareProps := &CfnPortfolioShareProps{
//   	AccountId: jsii.String("accountId"),
//   	PortfolioId: jsii.String("portfolioId"),
//
//   	// the properties below are optional
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   	ShareTagOptions: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioshare.html
//
type CfnPortfolioShareProps struct {
	// The AWS account ID.
	//
	// For example, `123456789012` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioshare.html#cfn-servicecatalog-portfolioshare-accountid
	//
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The portfolio identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioshare.html#cfn-servicecatalog-portfolioshare-portfolioid
	//
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
	// The language code.
	//
	// - `jp` - Japanese
	// - `zh` - Chinese.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioshare.html#cfn-servicecatalog-portfolioshare-acceptlanguage
	//
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// Indicates whether TagOptions sharing is enabled or disabled for the portfolio share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioshare.html#cfn-servicecatalog-portfolioshare-sharetagoptions
	//
	ShareTagOptions interface{} `field:"optional" json:"shareTagOptions" yaml:"shareTagOptions"`
}

