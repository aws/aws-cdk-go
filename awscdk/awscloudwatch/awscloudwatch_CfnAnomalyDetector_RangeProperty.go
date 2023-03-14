package awscloudwatch


// Each `Range` specifies one range of days or times to exclude from use for training or updating an anomaly detection model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rangeProperty := &rangeProperty{
//   	endTime: jsii.String("endTime"),
//   	startTime: jsii.String("startTime"),
//   }
//
type CfnAnomalyDetector_RangeProperty struct {
	// The end time of the range to exclude.
	//
	// The format is `yyyy-MM-dd'T'HH:mm:ss` . For example, `2019-07-01T23:59:59` .
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// The start time of the range to exclude.
	//
	// The format is `yyyy-MM-dd'T'HH:mm:ss` . For example, `2019-07-01T23:59:59` .
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

