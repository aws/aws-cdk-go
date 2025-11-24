package mixinsawsaps


// The Amazon Managed Service for Prometheus metric labels associated with the anomaly detector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   labelProperty := &LabelProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-label.html
//
type CfnAnomalyDetectorPropsMixin_LabelProperty struct {
	// The key of the label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-label.html#cfn-aps-anomalydetector-label-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for this label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-label.html#cfn-aps-anomalydetector-label-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

