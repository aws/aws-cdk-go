package awsodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCloudAutonomousVmCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCloudAutonomousVmClusterProps := &CfnCloudAutonomousVmClusterProps{
//   	AutonomousDataStorageSizeInTBs: jsii.Number(123),
//   	CloudExadataInfrastructureId: jsii.String("cloudExadataInfrastructureId"),
//   	CpuCoreCountPerNode: jsii.Number(123),
//   	DbServers: []*string{
//   		jsii.String("dbServers"),
//   	},
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	IsMtlsEnabledVmCluster: jsii.Boolean(false),
//   	LicenseModel: jsii.String("licenseModel"),
//   	MaintenanceWindow: &MaintenanceWindowProperty{
//   		DaysOfWeek: []*string{
//   			jsii.String("daysOfWeek"),
//   		},
//   		HoursOfDay: []interface{}{
//   			jsii.Number(123),
//   		},
//   		LeadTimeInWeeks: jsii.Number(123),
//   		Months: []*string{
//   			jsii.String("months"),
//   		},
//   		Preference: jsii.String("preference"),
//   		WeeksOfMonth: []interface{}{
//   			jsii.Number(123),
//   		},
//   	},
//   	MemoryPerOracleComputeUnitInGBs: jsii.Number(123),
//   	OdbNetworkId: jsii.String("odbNetworkId"),
//   	ScanListenerPortNonTls: jsii.Number(123),
//   	ScanListenerPortTls: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeZone: jsii.String("timeZone"),
//   	TotalContainerDatabases: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudautonomousvmcluster.html
//
type CfnCloudAutonomousVmClusterProps struct {
	// The data storage size allocated for Autonomous Databases in the Autonomous VM cluster, in TB.
	//
	// Required when creating an Autonomous VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudautonomousvmcluster.html#cfn-odb-cloudautonomousvmcluster-autonomousdatastoragesizeintbs
	//
	AutonomousDataStorageSizeInTBs *float64 `field:"optional" json:"autonomousDataStorageSizeInTBs" yaml:"autonomousDataStorageSizeInTBs"`
	// The unique identifier of the Cloud Exadata Infrastructure containing this Autonomous VM cluster.
	//
	// Required when creating an Autonomous VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudautonomousvmcluster.html#cfn-odb-cloudautonomousvmcluster-cloudexadatainfrastructureid
	//
	CloudExadataInfrastructureId *string `field:"optional" json:"cloudExadataInfrastructureId" yaml:"cloudExadataInfrastructureId"`
	// The number of CPU cores enabled per node in the Autonomous VM cluster.
	//
	// Required when creating an Autonomous VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudautonomousvmcluster.html#cfn-odb-cloudautonomousvmcluster-cpucorecountpernode
	//
	CpuCoreCountPerNode *float64 `field:"optional" json:"cpuCoreCountPerNode" yaml:"cpuCoreCountPerNode"`
	// The list of database servers associated with the Autonomous VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudautonomousvmcluster.html#cfn-odb-cloudautonomousvmcluster-dbservers
	//
	DbServers *[]*string `field:"optional" json:"dbServers" yaml:"dbServers"`
	// The user-provided description of the Autonomous VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudautonomousvmcluster.html#cfn-odb-cloudautonomousvmcluster-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the Autonomous VM cluster.
	//
	// Required when creating an Autonomous VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudautonomousvmcluster.html#cfn-odb-cloudautonomousvmcluster-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Specifies whether mutual TLS (mTLS) authentication is enabled for the Autonomous VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudautonomousvmcluster.html#cfn-odb-cloudautonomousvmcluster-ismtlsenabledvmcluster
	//
	IsMtlsEnabledVmCluster interface{} `field:"optional" json:"isMtlsEnabledVmCluster" yaml:"isMtlsEnabledVmCluster"`
	// The Oracle license model that applies to the Autonomous VM cluster.
	//
	// Valid values are `LICENSE_INCLUDED` or `BRING_YOUR_OWN_LICENSE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudautonomousvmcluster.html#cfn-odb-cloudautonomousvmcluster-licensemodel
	//
	LicenseModel *string `field:"optional" json:"licenseModel" yaml:"licenseModel"`
	// The scheduling details for the maintenance window.
	//
	// Patching and system updates take place during the maintenance window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudautonomousvmcluster.html#cfn-odb-cloudautonomousvmcluster-maintenancewindow
	//
	MaintenanceWindow interface{} `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// The amount of memory allocated per Oracle Compute Unit, in GB.
	//
	// Required when creating an Autonomous VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudautonomousvmcluster.html#cfn-odb-cloudautonomousvmcluster-memoryperoraclecomputeunitingbs
	//
	MemoryPerOracleComputeUnitInGBs *float64 `field:"optional" json:"memoryPerOracleComputeUnitInGBs" yaml:"memoryPerOracleComputeUnitInGBs"`
	// The unique identifier of the ODB network associated with this Autonomous VM cluster.
	//
	// Required when creating an Autonomous VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudautonomousvmcluster.html#cfn-odb-cloudautonomousvmcluster-odbnetworkid
	//
	OdbNetworkId *string `field:"optional" json:"odbNetworkId" yaml:"odbNetworkId"`
	// The SCAN listener port for non-TLS (TCP) protocol.
	//
	// The default is 1521.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudautonomousvmcluster.html#cfn-odb-cloudautonomousvmcluster-scanlistenerportnontls
	//
	ScanListenerPortNonTls *float64 `field:"optional" json:"scanListenerPortNonTls" yaml:"scanListenerPortNonTls"`
	// The SCAN listener port for TLS (TCP) protocol.
	//
	// The default is 2484.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudautonomousvmcluster.html#cfn-odb-cloudautonomousvmcluster-scanlistenerporttls
	//
	ScanListenerPortTls *float64 `field:"optional" json:"scanListenerPortTls" yaml:"scanListenerPortTls"`
	// Tags to assign to the Autonomous Vm Cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudautonomousvmcluster.html#cfn-odb-cloudautonomousvmcluster-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The time zone of the Autonomous VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudautonomousvmcluster.html#cfn-odb-cloudautonomousvmcluster-timezone
	//
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
	// The total number of Autonomous Container Databases that can be created with the allocated local storage.
	//
	// Required when creating an Autonomous VM cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-cloudautonomousvmcluster.html#cfn-odb-cloudautonomousvmcluster-totalcontainerdatabases
	//
	TotalContainerDatabases *float64 `field:"optional" json:"totalContainerDatabases" yaml:"totalContainerDatabases"`
}

