package awsservicecatalog


// Properties for defining a `CfnLaunchRoleConstraint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLaunchRoleConstraintProps := &cfnLaunchRoleConstraintProps{
//   	portfolioId: jsii.String("portfolioId"),
//   	productId: jsii.String("productId"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   	description: jsii.String("description"),
//   	localRoleName: jsii.String("localRoleName"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnLaunchRoleConstraintProps struct {
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
	// You are required to specify either the `RoleArn` or the `LocalRoleName` but can't use both.
	//
	// If you specify the `LocalRoleName` property, when an account uses the launch constraint, the IAM role with that name in the account will be used. This allows launch-role constraints to be account-agnostic so the administrator can create fewer resources per shared account.
	//
	// The given role name must exist in the account used to create the launch constraint and the account of the user who launches a product with this launch constraint.
	LocalRoleName *string `field:"optional" json:"localRoleName" yaml:"localRoleName"`
	// The ARN of the launch role.
	//
	// You are required to specify `RoleArn` or `LocalRoleName` but can't use both.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

