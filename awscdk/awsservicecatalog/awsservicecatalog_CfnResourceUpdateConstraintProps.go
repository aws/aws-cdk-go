package awsservicecatalog


// Properties for defining a `CfnResourceUpdateConstraint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceUpdateConstraintProps := &cfnResourceUpdateConstraintProps{
//   	portfolioId: jsii.String("portfolioId"),
//   	productId: jsii.String("productId"),
//   	tagUpdateOnProvisionedProduct: jsii.String("tagUpdateOnProvisionedProduct"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   }
//
type CfnResourceUpdateConstraintProps struct {
	// The portfolio identifier.
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
	// The product identifier.
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// If set to `ALLOWED` , lets users change tags in a [CloudFormationProvisionedProduct](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html) resource.
	//
	// If set to `NOT_ALLOWED` , prevents users from changing tags in a [CloudFormationProvisionedProduct](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html) resource.
	TagUpdateOnProvisionedProduct *string `field:"required" json:"tagUpdateOnProvisionedProduct" yaml:"tagUpdateOnProvisionedProduct"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The description of the constraint.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

