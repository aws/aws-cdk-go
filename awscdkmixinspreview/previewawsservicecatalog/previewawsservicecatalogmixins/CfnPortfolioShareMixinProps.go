package previewawsservicecatalogmixins


// Properties for CfnPortfolioSharePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPortfolioShareMixinProps := &CfnPortfolioShareMixinProps{
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   	AccountId: jsii.String("accountId"),
//   	PortfolioId: jsii.String("portfolioId"),
//   	ShareTagOptions: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioshare.html
//
type CfnPortfolioShareMixinProps struct {
	// The language code.
	//
	// - `jp` - Japanese
	// - `zh` - Chinese.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioshare.html#cfn-servicecatalog-portfolioshare-acceptlanguage
	//
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The AWS account ID.
	//
	// For example, `123456789012` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioshare.html#cfn-servicecatalog-portfolioshare-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The portfolio identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioshare.html#cfn-servicecatalog-portfolioshare-portfolioid
	//
	PortfolioId *string `field:"optional" json:"portfolioId" yaml:"portfolioId"`
	// Indicates whether TagOptions sharing is enabled or disabled for the portfolio share.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioshare.html#cfn-servicecatalog-portfolioshare-sharetagoptions
	//
	ShareTagOptions interface{} `field:"optional" json:"shareTagOptions" yaml:"shareTagOptions"`
}

