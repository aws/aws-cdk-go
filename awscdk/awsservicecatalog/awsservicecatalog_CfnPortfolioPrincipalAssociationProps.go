package awsservicecatalog


// Properties for defining a `CfnPortfolioPrincipalAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPortfolioPrincipalAssociationProps := &CfnPortfolioPrincipalAssociationProps{
//   	PortfolioId: jsii.String("portfolioId"),
//   	PrincipalArn: jsii.String("principalArn"),
//   	PrincipalType: jsii.String("principalType"),
//
//   	// the properties below are optional
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   }
//
type CfnPortfolioPrincipalAssociationProps struct {
	// The portfolio identifier.
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
	// The ARN of the principal ( IAM user, role, or group).
	PrincipalArn *string `field:"required" json:"principalArn" yaml:"principalArn"`
	// The principal type. The supported value is `IAM` .
	//
	// *Allowed Values* : `IAM`.
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
	// The language code.
	//
	// - `jp` - Japanese
	// - `zh` - Chinese.
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
}

