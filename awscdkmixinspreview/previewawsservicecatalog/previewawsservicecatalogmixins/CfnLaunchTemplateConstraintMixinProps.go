package previewawsservicecatalogmixins


// Properties for CfnLaunchTemplateConstraintPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLaunchTemplateConstraintMixinProps := &CfnLaunchTemplateConstraintMixinProps{
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   	Description: jsii.String("description"),
//   	PortfolioId: jsii.String("portfolioId"),
//   	ProductId: jsii.String("productId"),
//   	Rules: jsii.String("rules"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html
//
type CfnLaunchTemplateConstraintMixinProps struct {
	// The language code.
	//
	// - `jp` - Japanese
	// - `zh` - Chinese.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-acceptlanguage
	//
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The description of the constraint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The portfolio identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-portfolioid
	//
	PortfolioId *string `field:"optional" json:"portfolioId" yaml:"portfolioId"`
	// The product identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-productid
	//
	ProductId *string `field:"optional" json:"productId" yaml:"productId"`
	// The constraint rules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchtemplateconstraint.html#cfn-servicecatalog-launchtemplateconstraint-rules
	//
	Rules *string `field:"optional" json:"rules" yaml:"rules"`
}

