package awssagemaker


// Used to set feature group throughput configuration.
//
// There are two modes: `ON_DEMAND` and `PROVISIONED` . With on-demand mode, you are charged for data reads and writes that your application performs on your feature group. You do not need to specify read and write throughput because Feature Store accommodates your workloads as they ramp up and down. You can switch a feature group to on-demand only once in a 24 hour period. With provisioned throughput mode, you specify the read and write capacity per second that you expect your application to require, and you are billed based on those limits. Exceeding provisioned throughput will result in your requests being throttled.
//
// Note: `PROVISIONED` throughput mode is supported only for feature groups that are offline-only, or use the [`Standard`](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_OnlineStoreConfig.html#sagemaker-Type-OnlineStoreConfig-StorageType) tier online store.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   throughputConfigProperty := &ThroughputConfigProperty{
//   	ThroughputMode: jsii.String("throughputMode"),
//
//   	// the properties below are optional
//   	ProvisionedReadCapacityUnits: jsii.Number(123),
//   	ProvisionedWriteCapacityUnits: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-throughputconfig.html
//
type CfnFeatureGroup_ThroughputConfigProperty struct {
	// The mode used for your feature group throughput: `ON_DEMAND` or `PROVISIONED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-throughputconfig.html#cfn-sagemaker-featuregroup-throughputconfig-throughputmode
	//
	ThroughputMode *string `field:"required" json:"throughputMode" yaml:"throughputMode"`
	// For provisioned feature groups with online store enabled, this indicates the read throughput you are billed for and can consume without throttling.
	//
	// This field is not applicable for on-demand feature groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-throughputconfig.html#cfn-sagemaker-featuregroup-throughputconfig-provisionedreadcapacityunits
	//
	ProvisionedReadCapacityUnits *float64 `field:"optional" json:"provisionedReadCapacityUnits" yaml:"provisionedReadCapacityUnits"`
	// For provisioned feature groups, this indicates the write throughput you are billed for and can consume without throttling.
	//
	// This field is not applicable for on-demand feature groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-featuregroup-throughputconfig.html#cfn-sagemaker-featuregroup-throughputconfig-provisionedwritecapacityunits
	//
	ProvisionedWriteCapacityUnits *float64 `field:"optional" json:"provisionedWriteCapacityUnits" yaml:"provisionedWriteCapacityUnits"`
}

