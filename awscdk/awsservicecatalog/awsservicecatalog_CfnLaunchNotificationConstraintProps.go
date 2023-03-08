package awsservicecatalog


// Properties for defining a `CfnLaunchNotificationConstraint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLaunchNotificationConstraintProps := &cfnLaunchNotificationConstraintProps{
//   	notificationArns: []*string{
//   		jsii.String("notificationArns"),
//   	},
//   	portfolioId: jsii.String("portfolioId"),
//   	productId: jsii.String("productId"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   }
//
type CfnLaunchNotificationConstraintProps struct {
	// The notification ARNs.
	NotificationArns *[]*string `field:"required" json:"notificationArns" yaml:"notificationArns"`
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
	// The description of the constraint.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

