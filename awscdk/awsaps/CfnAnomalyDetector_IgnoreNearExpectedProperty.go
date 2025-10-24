package awsaps


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ignoreNearExpectedProperty := &IgnoreNearExpectedProperty{
//   	Amount: jsii.Number(123),
//   	Ratio: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-ignorenearexpected.html
//
type CfnAnomalyDetector_IgnoreNearExpectedProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-ignorenearexpected.html#cfn-aps-anomalydetector-ignorenearexpected-amount
	//
	Amount *float64 `field:"optional" json:"amount" yaml:"amount"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-ignorenearexpected.html#cfn-aps-anomalydetector-ignorenearexpected-ratio
	//
	Ratio *float64 `field:"optional" json:"ratio" yaml:"ratio"`
}

