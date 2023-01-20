package awsec2


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
	// `CfnFlowLog.DestinationOptionsProperty.FileFormat`.
	FileFormat *string `field:"required" json:"fileFormat" yaml:"fileFormat"`
	// `CfnFlowLog.DestinationOptionsProperty.HiveCompatiblePartitions`.
	HiveCompatiblePartitions interface{} `field:"required" json:"hiveCompatiblePartitions" yaml:"hiveCompatiblePartitions"`
	// `CfnFlowLog.DestinationOptionsProperty.PerHourPartition`.
	PerHourPartition interface{} `field:"required" json:"perHourPartition" yaml:"perHourPartition"`
}

