package awsservicecatalog


// Properties for defining a `CfnStackSetConstraint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStackSetConstraintProps := &cfnStackSetConstraintProps{
//   	accountList: []*string{
//   		jsii.String("accountList"),
//   	},
//   	adminRole: jsii.String("adminRole"),
//   	description: jsii.String("description"),
//   	executionRole: jsii.String("executionRole"),
//   	portfolioId: jsii.String("portfolioId"),
//   	productId: jsii.String("productId"),
//   	regionList: []*string{
//   		jsii.String("regionList"),
//   	},
//   	stackInstanceControl: jsii.String("stackInstanceControl"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   }
//
type CfnStackSetConstraintProps struct {
	// One or more AWS accounts that will have access to the provisioned product.
	AccountList *[]*string `field:"required" json:"accountList" yaml:"accountList"`
	// AdminRole ARN.
	AdminRole *string `field:"required" json:"adminRole" yaml:"adminRole"`
	// The description of the constraint.
	Description *string `field:"required" json:"description" yaml:"description"`
	// ExecutionRole name.
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// The portfolio identifier.
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
	// The product identifier.
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// One or more AWS Regions where the provisioned product will be available.
	//
	// Applicable only to a `CFN_STACKSET` provisioned product type.
	//
	// The specified Regions should be within the list of Regions from the `STACKSET` constraint. To get the list of Regions in the `STACKSET` constraint, use the `DescribeProvisioningParameters` operation.
	//
	// If no values are specified, the default value is all Regions from the `STACKSET` constraint.
	RegionList *[]*string `field:"required" json:"regionList" yaml:"regionList"`
	// Permission to create, update, and delete stack instances.
	//
	// Choose from ALLOWED and NOT_ALLOWED.
	StackInstanceControl *string `field:"required" json:"stackInstanceControl" yaml:"stackInstanceControl"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
}

