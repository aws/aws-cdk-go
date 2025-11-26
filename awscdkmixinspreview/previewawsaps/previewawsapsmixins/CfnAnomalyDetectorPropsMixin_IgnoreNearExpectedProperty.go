package previewawsapsmixins


// Configuration for threshold settings that determine when values near expected values should be ignored during anomaly detection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ignoreNearExpectedProperty := &IgnoreNearExpectedProperty{
//   	Amount: jsii.Number(123),
//   	Ratio: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-ignorenearexpected.html
//
type CfnAnomalyDetectorPropsMixin_IgnoreNearExpectedProperty struct {
	// The absolute amount by which values can differ from expected values before being considered anomalous.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-ignorenearexpected.html#cfn-aps-anomalydetector-ignorenearexpected-amount
	//
	Amount *float64 `field:"optional" json:"amount" yaml:"amount"`
	// The ratio by which values can differ from expected values before being considered anomalous.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-ignorenearexpected.html#cfn-aps-anomalydetector-ignorenearexpected-ratio
	//
	Ratio *float64 `field:"optional" json:"ratio" yaml:"ratio"`
}

