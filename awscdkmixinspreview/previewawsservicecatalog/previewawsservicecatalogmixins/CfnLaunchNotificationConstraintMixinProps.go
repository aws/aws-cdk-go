package previewawsservicecatalogmixins


// Properties for CfnLaunchNotificationConstraintPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLaunchNotificationConstraintMixinProps := &CfnLaunchNotificationConstraintMixinProps{
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   	Description: jsii.String("description"),
//   	NotificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	PortfolioId: jsii.String("portfolioId"),
//   	ProductId: jsii.String("productId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchnotificationconstraint.html
//
type CfnLaunchNotificationConstraintMixinProps struct {
	// The language code.
	//
	// - `jp` - Japanese
	// - `zh` - Chinese.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchnotificationconstraint.html#cfn-servicecatalog-launchnotificationconstraint-acceptlanguage
	//
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The description of the constraint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchnotificationconstraint.html#cfn-servicecatalog-launchnotificationconstraint-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The notification ARNs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchnotificationconstraint.html#cfn-servicecatalog-launchnotificationconstraint-notificationarns
	//
	NotificationArns *[]*string `field:"optional" json:"notificationArns" yaml:"notificationArns"`
	// The portfolio identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchnotificationconstraint.html#cfn-servicecatalog-launchnotificationconstraint-portfolioid
	//
	PortfolioId *string `field:"optional" json:"portfolioId" yaml:"portfolioId"`
	// The product identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchnotificationconstraint.html#cfn-servicecatalog-launchnotificationconstraint-productid
	//
	ProductId *string `field:"optional" json:"productId" yaml:"productId"`
}

