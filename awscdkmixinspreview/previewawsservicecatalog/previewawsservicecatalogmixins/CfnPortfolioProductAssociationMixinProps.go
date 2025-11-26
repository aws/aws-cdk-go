package previewawsservicecatalogmixins


// Properties for CfnPortfolioProductAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPortfolioProductAssociationMixinProps := &CfnPortfolioProductAssociationMixinProps{
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   	PortfolioId: jsii.String("portfolioId"),
//   	ProductId: jsii.String("productId"),
//   	SourcePortfolioId: jsii.String("sourcePortfolioId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioproductassociation.html
//
type CfnPortfolioProductAssociationMixinProps struct {
	// The language code.
	//
	// - `jp` - Japanese
	// - `zh` - Chinese.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioproductassociation.html#cfn-servicecatalog-portfolioproductassociation-acceptlanguage
	//
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The portfolio identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioproductassociation.html#cfn-servicecatalog-portfolioproductassociation-portfolioid
	//
	PortfolioId *string `field:"optional" json:"portfolioId" yaml:"portfolioId"`
	// The product identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioproductassociation.html#cfn-servicecatalog-portfolioproductassociation-productid
	//
	ProductId *string `field:"optional" json:"productId" yaml:"productId"`
	// The identifier of the source portfolio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioproductassociation.html#cfn-servicecatalog-portfolioproductassociation-sourceportfolioid
	//
	SourcePortfolioId *string `field:"optional" json:"sourcePortfolioId" yaml:"sourcePortfolioId"`
}

