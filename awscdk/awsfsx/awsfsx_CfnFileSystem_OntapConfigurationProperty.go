package awsfsx


// The configuration for this Amazon FSx for NetApp ONTAP file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ontapConfigurationProperty := &ontapConfigurationProperty{
//   	deploymentType: jsii.String("deploymentType"),
//
//   	// the properties below are optional
//   	automaticBackupRetentionDays: jsii.Number(123),
//   	dailyAutomaticBackupStartTime: jsii.String("dailyAutomaticBackupStartTime"),
//   	diskIopsConfiguration: &diskIopsConfigurationProperty{
//   		iops: jsii.Number(123),
//   		mode: jsii.String("mode"),
//   	},
//   	endpointIpAddressRange: jsii.String("endpointIpAddressRange"),
//   	fsxAdminPassword: jsii.String("fsxAdminPassword"),
//   	preferredSubnetId: jsii.String("preferredSubnetId"),
//   	routeTableIds: []*string{
//   		jsii.String("routeTableIds"),
//   	},
//   	throughputCapacity: jsii.Number(123),
//   	weeklyMaintenanceStartTime: jsii.String("weeklyMaintenanceStartTime"),
//   }
//
type CfnFileSystem_OntapConfigurationProperty struct {
	// Specifies the FSx for ONTAP file system deployment type to use in creating the file system.
	//
	// - `MULTI_AZ_1` - (Default) A high availability file system configured for Multi-AZ redundancy to tolerate temporary Availability Zone (AZ) unavailability.
	// - `SINGLE_AZ_1` - A file system configured for Single-AZ redundancy.
	//
	// For information about the use cases for Multi-AZ and Single-AZ deployments, refer to [Choosing a file system deployment type](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/high-availability-AZ.html) .
	DeploymentType *string `field:"required" json:"deploymentType" yaml:"deploymentType"`
	// The number of days to retain automatic backups.
	//
	// Setting this property to `0` disables automatic backups. You can retain automatic backups for a maximum of 90 days. The default is `0` .
	AutomaticBackupRetentionDays *float64 `field:"optional" json:"automaticBackupRetentionDays" yaml:"automaticBackupRetentionDays"`
	// A recurring daily time, in the format `HH:MM` .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour. For example, `05:00` specifies 5 AM daily.
	DailyAutomaticBackupStartTime *string `field:"optional" json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// The SSD IOPS configuration for the FSx for ONTAP file system.
	DiskIopsConfiguration interface{} `field:"optional" json:"diskIopsConfiguration" yaml:"diskIopsConfiguration"`
	// (Multi-AZ only) Specifies the IP address range in which the endpoints to access your file system will be created.
	//
	// By default, Amazon FSx selects an unused IP address range for you from the 198.19.* range.
	//
	// > The Endpoint IP address range you select for your file system must exist outside the VPC's CIDR range and must be at least /30 or larger.
	EndpointIpAddressRange *string `field:"optional" json:"endpointIpAddressRange" yaml:"endpointIpAddressRange"`
	// The ONTAP administrative password for the `fsxadmin` user with which you administer your file system using the NetApp ONTAP CLI and REST API.
	FsxAdminPassword *string `field:"optional" json:"fsxAdminPassword" yaml:"fsxAdminPassword"`
	// Required when `DeploymentType` is set to `MULTI_AZ_1` .
	//
	// This specifies the subnet in which you want the preferred file server to be located.
	PreferredSubnetId *string `field:"optional" json:"preferredSubnetId" yaml:"preferredSubnetId"`
	// (Multi-AZ only) Specifies the virtual private cloud (VPC) route tables in which your file system's endpoints will be created.
	//
	// You should specify all VPC route tables associated with the subnets in which your clients are located. By default, Amazon FSx selects your VPC's default route table.
	RouteTableIds *[]*string `field:"optional" json:"routeTableIds" yaml:"routeTableIds"`
	// Sets the throughput capacity for the file system that you're creating.
	//
	// Valid values are 128, 256, 512, 1024, and 2048 MBps.
	ThroughputCapacity *float64 `field:"optional" json:"throughputCapacity" yaml:"throughputCapacity"`
	// A recurring weekly time, in the format `D:HH:MM` .
	//
	// `D` is the day of the week, for which 1 represents Monday and 7 represents Sunday. For further details, see [the ISO-8601 spec as described on Wikipedia](https://docs.aws.amazon.com/https://en.wikipedia.org/wiki/ISO_week_date) .
	//
	// `HH` is the zero-padded hour of the day (0-23), and `MM` is the zero-padded minute of the hour.
	//
	// For example, `1:05:00` specifies maintenance at 5 AM Monday.
	WeeklyMaintenanceStartTime *string `field:"optional" json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

