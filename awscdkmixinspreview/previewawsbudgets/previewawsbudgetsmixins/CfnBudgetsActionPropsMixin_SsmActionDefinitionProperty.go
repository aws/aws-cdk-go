package previewawsbudgetsmixins


// The Amazon EC2 Systems Manager ( SSM ) action definition details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ssmActionDefinitionProperty := &SsmActionDefinitionProperty{
//   	InstanceIds: []*string{
//   		jsii.String("instanceIds"),
//   	},
//   	Region: jsii.String("region"),
//   	Subtype: jsii.String("subtype"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budgetsaction-ssmactiondefinition.html
//
type CfnBudgetsActionPropsMixin_SsmActionDefinitionProperty struct {
	// The EC2 and RDS instance IDs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budgetsaction-ssmactiondefinition.html#cfn-budgets-budgetsaction-ssmactiondefinition-instanceids
	//
	InstanceIds *[]*string `field:"optional" json:"instanceIds" yaml:"instanceIds"`
	// The Region to run the ( SSM ) document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budgetsaction-ssmactiondefinition.html#cfn-budgets-budgetsaction-ssmactiondefinition-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The action subType.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budgetsaction-ssmactiondefinition.html#cfn-budgets-budgetsaction-ssmactiondefinition-subtype
	//
	Subtype *string `field:"optional" json:"subtype" yaml:"subtype"`
}

