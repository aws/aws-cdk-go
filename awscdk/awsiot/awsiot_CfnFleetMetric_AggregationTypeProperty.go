package awsiot


// The type of aggregation queries.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aggregationTypeProperty := &aggregationTypeProperty{
//   	name: jsii.String("name"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnFleetMetric_AggregationTypeProperty struct {
	// The name of the aggregation type.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of the values of aggregation types.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

