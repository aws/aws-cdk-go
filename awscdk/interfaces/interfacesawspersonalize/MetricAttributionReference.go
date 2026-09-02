package interfacesawspersonalize


// A reference to a MetricAttribution resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricAttributionReference := &MetricAttributionReference{
//   	MetricAttributionArn: jsii.String("metricAttributionArn"),
//   }
//
type MetricAttributionReference struct {
	// The MetricAttributionArn of the MetricAttribution resource.
	MetricAttributionArn *string `field:"required" json:"metricAttributionArn" yaml:"metricAttributionArn"`
}

