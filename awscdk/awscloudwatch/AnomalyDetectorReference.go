package awscloudwatch


// A reference to a AnomalyDetector resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   anomalyDetectorReference := &AnomalyDetectorReference{
//   	AnomalyDetectorId: jsii.String("anomalyDetectorId"),
//   }
//
type AnomalyDetectorReference struct {
	// The Id of the AnomalyDetector resource.
	AnomalyDetectorId *string `field:"required" json:"anomalyDetectorId" yaml:"anomalyDetectorId"`
}

