package awsaps


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   missingDataActionProperty := &MissingDataActionProperty{
//   	MarkAsAnomaly: jsii.Boolean(false),
//   	Skip: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-missingdataaction.html
//
type CfnAnomalyDetector_MissingDataActionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-missingdataaction.html#cfn-aps-anomalydetector-missingdataaction-markasanomaly
	//
	MarkAsAnomaly interface{} `field:"optional" json:"markAsAnomaly" yaml:"markAsAnomaly"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-missingdataaction.html#cfn-aps-anomalydetector-missingdataaction-skip
	//
	Skip interface{} `field:"optional" json:"skip" yaml:"skip"`
}

