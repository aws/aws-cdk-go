package awsfsx


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lustreConfigurationProperty := &LustreConfigurationProperty{
//   	DeploymentType: jsii.String("deploymentType"),
//   	MetadataConfiguration: &MetadataConfigurationProperty{
//   		StorageCapacity: jsii.Number(123),
//   	},
//   	PerUnitStorageThroughput: jsii.Number(123),
//
//   	// the properties below are optional
//   	LogConfiguration: &LogConfigurationProperty{
//   		Level: jsii.String("level"),
//
//   		// the properties below are optional
//   		Destination: jsii.String("destination"),
//   	},
//   	MountName: jsii.String("mountName"),
//   	WeeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-lustreconfiguration.html
//
type CfnFileCache_LustreConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-lustreconfiguration.html#cfn-fsx-filecache-lustreconfiguration-deploymenttype
	//
	DeploymentType *string `field:"required" json:"deploymentType" yaml:"deploymentType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-lustreconfiguration.html#cfn-fsx-filecache-lustreconfiguration-metadataconfiguration
	//
	MetadataConfiguration interface{} `field:"required" json:"metadataConfiguration" yaml:"metadataConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-lustreconfiguration.html#cfn-fsx-filecache-lustreconfiguration-perunitstoragethroughput
	//
	PerUnitStorageThroughput *float64 `field:"required" json:"perUnitStorageThroughput" yaml:"perUnitStorageThroughput"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-lustreconfiguration.html#cfn-fsx-filecache-lustreconfiguration-logconfiguration
	//
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-lustreconfiguration.html#cfn-fsx-filecache-lustreconfiguration-mountname
	//
	MountName *string `field:"optional" json:"mountName" yaml:"mountName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-lustreconfiguration.html#cfn-fsx-filecache-lustreconfiguration-weeklymaintenancestarttime
	//
	WeeklyMaintenanceStartTime *string `field:"optional" json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

