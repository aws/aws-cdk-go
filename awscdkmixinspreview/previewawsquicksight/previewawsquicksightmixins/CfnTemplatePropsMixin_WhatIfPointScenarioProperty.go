package previewawsquicksightmixins


// Provides the forecast to meet the target for a particular date.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   whatIfPointScenarioProperty := &WhatIfPointScenarioProperty{
//   	Date: jsii.String("date"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-whatifpointscenario.html
//
type CfnTemplatePropsMixin_WhatIfPointScenarioProperty struct {
	// The date that you need the forecast results for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-whatifpointscenario.html#cfn-quicksight-template-whatifpointscenario-date
	//
	Date *string `field:"optional" json:"date" yaml:"date"`
	// The target value that you want to meet for the provided date.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-whatifpointscenario.html#cfn-quicksight-template-whatifpointscenario-value
	//
	// Default: - 0.
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

