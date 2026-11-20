package awsdsql


// An Advanced monitoring metric group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricGroupProperty := &MetricGroupProperty{
//   	MetricGroup: jsii.String("metricGroup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dsql-cluster-metricgroup.html
//
type CfnCluster_MetricGroupProperty struct {
	// The name of the metric group.
	//
	// Valid values: STATEMENT_STATISTICS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dsql-cluster-metricgroup.html#cfn-dsql-cluster-metricgroup-metricgroup
	//
	MetricGroup *string `field:"required" json:"metricGroup" yaml:"metricGroup"`
}

