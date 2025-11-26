package previewawsec2mixins


// Describes the destination options for a flow log.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   destinationOptionsProperty := &DestinationOptionsProperty{
//   	FileFormat: jsii.String("fileFormat"),
//   	HiveCompatiblePartitions: jsii.Boolean(false),
//   	PerHourPartition: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-flowlog-destinationoptions.html
//
type CfnFlowLogPropsMixin_DestinationOptionsProperty struct {
	// The format for the flow log.
	//
	// The default is `plain-text` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-flowlog-destinationoptions.html#cfn-ec2-flowlog-destinationoptions-fileformat
	//
	FileFormat *string `field:"optional" json:"fileFormat" yaml:"fileFormat"`
	// Indicates whether to use Hive-compatible prefixes for flow logs stored in Amazon S3.
	//
	// The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-flowlog-destinationoptions.html#cfn-ec2-flowlog-destinationoptions-hivecompatiblepartitions
	//
	HiveCompatiblePartitions interface{} `field:"optional" json:"hiveCompatiblePartitions" yaml:"hiveCompatiblePartitions"`
	// Indicates whether to partition the flow log per hour.
	//
	// This reduces the cost and response time for queries. The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-flowlog-destinationoptions.html#cfn-ec2-flowlog-destinationoptions-perhourpartition
	//
	PerHourPartition interface{} `field:"optional" json:"perHourPartition" yaml:"perHourPartition"`
}

