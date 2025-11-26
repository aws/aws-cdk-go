package previewawsdatabrewmixins


// A key-value pair to associate an expression's substitution variable names with their values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   substitutionValueProperty := &SubstitutionValueProperty{
//   	Value: jsii.String("value"),
//   	ValueReference: jsii.String("valueReference"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-ruleset-substitutionvalue.html
//
type CfnRulesetPropsMixin_SubstitutionValueProperty struct {
	// Value or column name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-ruleset-substitutionvalue.html#cfn-databrew-ruleset-substitutionvalue-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
	// Variable name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-ruleset-substitutionvalue.html#cfn-databrew-ruleset-substitutionvalue-valuereference
	//
	ValueReference *string `field:"optional" json:"valueReference" yaml:"valueReference"`
}

