package mixinsawsservicecatalog


// Properties for CfnPortfolioPrincipalAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPortfolioPrincipalAssociationMixinProps := &CfnPortfolioPrincipalAssociationMixinProps{
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   	PortfolioId: jsii.String("portfolioId"),
//   	PrincipalArn: jsii.String("principalArn"),
//   	PrincipalType: jsii.String("principalType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioprincipalassociation.html
//
type CfnPortfolioPrincipalAssociationMixinProps struct {
	// The language code.
	//
	// - `jp` - Japanese
	// - `zh` - Chinese.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioprincipalassociation.html#cfn-servicecatalog-portfolioprincipalassociation-acceptlanguage
	//
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The portfolio identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioprincipalassociation.html#cfn-servicecatalog-portfolioprincipalassociation-portfolioid
	//
	PortfolioId *string `field:"optional" json:"portfolioId" yaml:"portfolioId"`
	// The ARN of the principal ( IAM user, role, or group).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioprincipalassociation.html#cfn-servicecatalog-portfolioprincipalassociation-principalarn
	//
	PrincipalArn *string `field:"optional" json:"principalArn" yaml:"principalArn"`
	// The principal type.
	//
	// The supported values are `IAM` and `IAM_PATTERN` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioprincipalassociation.html#cfn-servicecatalog-portfolioprincipalassociation-principaltype
	//
	PrincipalType *string `field:"optional" json:"principalType" yaml:"principalType"`
}

