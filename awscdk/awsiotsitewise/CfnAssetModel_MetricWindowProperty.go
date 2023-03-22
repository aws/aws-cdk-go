package awsiotsitewise


// Contains a time interval window used for data aggregate computations (for example, average, sum, count, and so on).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricWindowProperty := &MetricWindowProperty{
//   	Tumbling: &TumblingWindowProperty{
//   		Interval: jsii.String("interval"),
//
//   		// the properties below are optional
//   		Offset: jsii.String("offset"),
//   	},
//   }
//
type CfnAssetModel_MetricWindowProperty struct {
	// The tumbling time interval window.
	Tumbling interface{} `field:"optional" json:"tumbling" yaml:"tumbling"`
}

