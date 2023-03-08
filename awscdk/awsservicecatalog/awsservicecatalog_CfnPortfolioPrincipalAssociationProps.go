package awsservicecatalog


// Properties for defining a `CfnPortfolioPrincipalAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPortfolioPrincipalAssociationProps := &cfnPortfolioPrincipalAssociationProps{
//   	portfolioId: jsii.String("portfolioId"),
//   	principalArn: jsii.String("principalArn"),
//   	principalType: jsii.String("principalType"),
//
//   	// the properties below are optional
//   	acceptLanguage: jsii.String("acceptLanguage"),
//   }
//
type CfnPortfolioPrincipalAssociationProps struct {
	// The portfolio identifier.
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
	// The ARN of the principal (IAM user, role, or group).
	//
	// This field allows an ARN with no `accountID` if `PrincipalType` is `IAM_PATTERN` .
	//
	// You can associate multiple `IAM` patterns even if the account has no principal with that name. This is useful in Principal Name Sharing if you want to share a principal without creating it in the account that owns the portfolio.
	PrincipalArn *string `field:"required" json:"principalArn" yaml:"principalArn"`
	// The principal type.
	//
	// The supported value is `IAM` if you use a fully defined ARN, or `IAM_PATTERN` if you use an ARN with no `accountID` .
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
	// The language code.
	//
	// - `en` - English (default)
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
}

