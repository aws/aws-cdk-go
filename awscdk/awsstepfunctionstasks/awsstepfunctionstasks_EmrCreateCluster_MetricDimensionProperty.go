package awsstepfunctionstasks


// A CloudWatch dimension, which is specified using a Key (known as a Name in CloudWatch), Value pair.
//
// By default, Amazon EMR uses
// one dimension whose Key is JobFlowID and Value is a variable representing the cluster ID, which is ${emr.clusterId}. This enables
// the rule to bootstrap when the cluster ID becomes available.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDimensionProperty := &metricDimensionProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_MetricDimension.html
//
type EmrCreateCluster_MetricDimensionProperty struct {
	// The dimension name.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The dimension value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

