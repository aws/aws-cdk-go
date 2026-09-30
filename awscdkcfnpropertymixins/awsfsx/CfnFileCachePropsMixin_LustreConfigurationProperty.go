package awsfsx


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   lustreConfigurationProperty := &LustreConfigurationProperty{
//   	DeploymentType: jsii.String("deploymentType"),
//   	LogConfiguration: &LogConfigurationProperty{
//   		Destination: jsii.String("destination"),
//   		Level: jsii.String("level"),
//   	},
//   	MetadataConfiguration: &MetadataConfigurationProperty{
//   		StorageCapacity: jsii.Number(123),
//   	},
//   	MountName: jsii.String("mountName"),
//   	PerUnitStorageThroughput: jsii.Number(123),
//   	WeeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-lustreconfiguration.html
//
type CfnFileCachePropsMixin_LustreConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-lustreconfiguration.html#cfn-fsx-filecache-lustreconfiguration-deploymenttype
	//
	DeploymentType *string `field:"optional" json:"deploymentType" yaml:"deploymentType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-lustreconfiguration.html#cfn-fsx-filecache-lustreconfiguration-logconfiguration
	//
	LogConfiguration interface{} `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-lustreconfiguration.html#cfn-fsx-filecache-lustreconfiguration-metadataconfiguration
	//
	MetadataConfiguration interface{} `field:"optional" json:"metadataConfiguration" yaml:"metadataConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-lustreconfiguration.html#cfn-fsx-filecache-lustreconfiguration-mountname
	//
	MountName *string `field:"optional" json:"mountName" yaml:"mountName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-lustreconfiguration.html#cfn-fsx-filecache-lustreconfiguration-perunitstoragethroughput
	//
	PerUnitStorageThroughput *float64 `field:"optional" json:"perUnitStorageThroughput" yaml:"perUnitStorageThroughput"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filecache-lustreconfiguration.html#cfn-fsx-filecache-lustreconfiguration-weeklymaintenancestarttime
	//
	WeeklyMaintenanceStartTime *string `field:"optional" json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

