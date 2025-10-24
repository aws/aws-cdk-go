package awsaps


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   randomCutForestConfigurationProperty := &RandomCutForestConfigurationProperty{
//   	Query: jsii.String("query"),
//
//   	// the properties below are optional
//   	IgnoreNearExpectedFromAbove: &IgnoreNearExpectedProperty{
//   		Amount: jsii.Number(123),
//   		Ratio: jsii.Number(123),
//   	},
//   	IgnoreNearExpectedFromBelow: &IgnoreNearExpectedProperty{
//   		Amount: jsii.Number(123),
//   		Ratio: jsii.Number(123),
//   	},
//   	SampleSize: jsii.Number(123),
//   	ShingleSize: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-randomcutforestconfiguration.html
//
type CfnAnomalyDetector_RandomCutForestConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-randomcutforestconfiguration.html#cfn-aps-anomalydetector-randomcutforestconfiguration-query
	//
	Query *string `field:"required" json:"query" yaml:"query"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-randomcutforestconfiguration.html#cfn-aps-anomalydetector-randomcutforestconfiguration-ignorenearexpectedfromabove
	//
	IgnoreNearExpectedFromAbove interface{} `field:"optional" json:"ignoreNearExpectedFromAbove" yaml:"ignoreNearExpectedFromAbove"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-randomcutforestconfiguration.html#cfn-aps-anomalydetector-randomcutforestconfiguration-ignorenearexpectedfrombelow
	//
	IgnoreNearExpectedFromBelow interface{} `field:"optional" json:"ignoreNearExpectedFromBelow" yaml:"ignoreNearExpectedFromBelow"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-randomcutforestconfiguration.html#cfn-aps-anomalydetector-randomcutforestconfiguration-samplesize
	//
	// Default: - 256.
	//
	SampleSize *float64 `field:"optional" json:"sampleSize" yaml:"sampleSize"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-randomcutforestconfiguration.html#cfn-aps-anomalydetector-randomcutforestconfiguration-shinglesize
	//
	// Default: - 8.
	//
	ShingleSize *float64 `field:"optional" json:"shingleSize" yaml:"shingleSize"`
}

