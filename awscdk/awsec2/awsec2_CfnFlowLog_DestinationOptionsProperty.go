package awsec2


// Describes the destination options for a flow log.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationOptionsProperty := &destinationOptionsProperty{
//   	fileFormat: jsii.String("fileFormat"),
//   	hiveCompatiblePartitions: jsii.Boolean(false),
//   	perHourPartition: jsii.Boolean(false),
//   }
//
type CfnFlowLog_DestinationOptionsProperty struct {
	// The format for the flow log.
	//
	// The default is `plain-text` .
	FileFormat *string `field:"required" json:"fileFormat" yaml:"fileFormat"`
	// Indicates whether to use Hive-compatible prefixes for flow logs stored in Amazon S3.
	//
	// The default is `false` .
	HiveCompatiblePartitions interface{} `field:"required" json:"hiveCompatiblePartitions" yaml:"hiveCompatiblePartitions"`
	// Indicates whether to partition the flow log per hour.
	//
	// This reduces the cost and response time for queries. The default is `false` .
	PerHourPartition interface{} `field:"required" json:"perHourPartition" yaml:"perHourPartition"`
}

