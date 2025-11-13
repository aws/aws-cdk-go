package awsaps


// The configuration for the anomaly detection algorithm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   anomalyDetectorConfigurationProperty := &AnomalyDetectorConfigurationProperty{
//   	RandomCutForest: &RandomCutForestConfigurationProperty{
//   		Query: jsii.String("query"),
//
//   		// the properties below are optional
//   		IgnoreNearExpectedFromAbove: &IgnoreNearExpectedProperty{
//   			Amount: jsii.Number(123),
//   			Ratio: jsii.Number(123),
//   		},
//   		IgnoreNearExpectedFromBelow: &IgnoreNearExpectedProperty{
//   			Amount: jsii.Number(123),
//   			Ratio: jsii.Number(123),
//   		},
//   		SampleSize: jsii.Number(123),
//   		ShingleSize: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-anomalydetectorconfiguration.html
//
type CfnAnomalyDetector_AnomalyDetectorConfigurationProperty struct {
	// The Random Cut Forest algorithm configuration for anomaly detection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-anomalydetector-anomalydetectorconfiguration.html#cfn-aps-anomalydetector-anomalydetectorconfiguration-randomcutforest
	//
	RandomCutForest interface{} `field:"required" json:"randomCutForest" yaml:"randomCutForest"`
}

