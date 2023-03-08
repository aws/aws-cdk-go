package awslookoutmetrics


// Contains information about a detector's configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   anomalyDetectorConfigProperty := &anomalyDetectorConfigProperty{
//   	anomalyDetectorFrequency: jsii.String("anomalyDetectorFrequency"),
//   }
//
type CfnAnomalyDetector_AnomalyDetectorConfigProperty struct {
	// The frequency at which the detector analyzes its source data.
	AnomalyDetectorFrequency *string `field:"required" json:"anomalyDetectorFrequency" yaml:"anomalyDetectorFrequency"`
}

