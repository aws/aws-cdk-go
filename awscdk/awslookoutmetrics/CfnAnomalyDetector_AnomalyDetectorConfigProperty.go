package awslookoutmetrics


// Contains information about a detector's configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   anomalyDetectorConfigProperty := &AnomalyDetectorConfigProperty{
//   	AnomalyDetectorFrequency: jsii.String("anomalyDetectorFrequency"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-anomalydetectorconfig.html
//
type CfnAnomalyDetector_AnomalyDetectorConfigProperty struct {
	// The frequency at which the detector analyzes its source data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lookoutmetrics-anomalydetector-anomalydetectorconfig.html#cfn-lookoutmetrics-anomalydetector-anomalydetectorconfig-anomalydetectorfrequency
	//
	AnomalyDetectorFrequency *string `field:"required" json:"anomalyDetectorFrequency" yaml:"anomalyDetectorFrequency"`
}

