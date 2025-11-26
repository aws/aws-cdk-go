package previewawsemrmixins


// `MetricDimension` is a subproperty of the `CloudWatchAlarmDefinition` property type.
//
// `MetricDimension` specifies a CloudWatch dimension, which is specified with a `Key` `Value` pair. The key is known as a `Name` in CloudWatch. By default, Amazon EMR uses one dimension whose `Key` is `JobFlowID` and `Value` is a variable representing the cluster ID, which is `${emr.clusterId}` . This enables the automatic scaling rule for EMR to bootstrap when the cluster ID becomes available during cluster creation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   metricDimensionProperty := &MetricDimensionProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-metricdimension.html
//
type CfnClusterPropsMixin_MetricDimensionProperty struct {
	// The dimension name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-metricdimension.html#cfn-emr-cluster-metricdimension-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The dimension value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-metricdimension.html#cfn-emr-cluster-metricdimension-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

