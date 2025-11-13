package interfacesawslogs


// A reference to a LogAnomalyDetector resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logAnomalyDetectorReference := &LogAnomalyDetectorReference{
//   	AnomalyDetectorArn: jsii.String("anomalyDetectorArn"),
//   }
//
type LogAnomalyDetectorReference struct {
	// The AnomalyDetectorArn of the LogAnomalyDetector resource.
	AnomalyDetectorArn *string `field:"required" json:"anomalyDetectorArn" yaml:"anomalyDetectorArn"`
}

