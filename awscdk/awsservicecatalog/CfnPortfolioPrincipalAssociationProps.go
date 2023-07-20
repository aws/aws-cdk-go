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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioprincipalassociation.html
//
type CfnPortfolioPrincipalAssociationProps struct {
	// The portfolio identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioprincipalassociation.html#cfn-servicecatalog-portfolioprincipalassociation-portfolioid
	//
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
	// The ARN of the principal ( IAM user, role, or group).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioprincipalassociation.html#cfn-servicecatalog-portfolioprincipalassociation-principalarn
	//
	PrincipalArn *string `field:"required" json:"principalArn" yaml:"principalArn"`
	// The principal type. The supported value is `IAM` .
	//
	// *Allowed Values* : `IAM`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioprincipalassociation.html#cfn-servicecatalog-portfolioprincipalassociation-principaltype
	//
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
	// The language code.
	//
	// - `jp` - Japanese
	// - `zh` - Chinese.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioprincipalassociation.html#cfn-servicecatalog-portfolioprincipalassociation-acceptlanguage
	//
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
}

