package awsdms


// Configuration parameters for provisioning a AWS DMS Serverless replication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   computeConfigProperty := &ComputeConfigProperty{
//   	MaxCapacityUnits: jsii.Number(123),
//
//   	// the properties below are optional
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	DnsNameServers: jsii.String("dnsNameServers"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	MinCapacityUnits: jsii.Number(123),
//   	MultiAz: jsii.Boolean(false),
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	ReplicationSubnetGroupId: jsii.String("replicationSubnetGroupId"),
//   	VpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationconfig-computeconfig.html
//
type CfnReplicationConfig_ComputeConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationconfig-computeconfig.html#cfn-dms-replicationconfig-computeconfig-maxcapacityunits
	//
	MaxCapacityUnits *float64 `field:"required" json:"maxCapacityUnits" yaml:"maxCapacityUnits"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationconfig-computeconfig.html#cfn-dms-replicationconfig-computeconfig-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationconfig-computeconfig.html#cfn-dms-replicationconfig-computeconfig-dnsnameservers
	//
	DnsNameServers *string `field:"optional" json:"dnsNameServers" yaml:"dnsNameServers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationconfig-computeconfig.html#cfn-dms-replicationconfig-computeconfig-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationconfig-computeconfig.html#cfn-dms-replicationconfig-computeconfig-mincapacityunits
	//
	MinCapacityUnits *float64 `field:"optional" json:"minCapacityUnits" yaml:"minCapacityUnits"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationconfig-computeconfig.html#cfn-dms-replicationconfig-computeconfig-multiaz
	//
	MultiAz interface{} `field:"optional" json:"multiAz" yaml:"multiAz"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationconfig-computeconfig.html#cfn-dms-replicationconfig-computeconfig-preferredmaintenancewindow
	//
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationconfig-computeconfig.html#cfn-dms-replicationconfig-computeconfig-replicationsubnetgroupid
	//
	ReplicationSubnetGroupId *string `field:"optional" json:"replicationSubnetGroupId" yaml:"replicationSubnetGroupId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationconfig-computeconfig.html#cfn-dms-replicationconfig-computeconfig-vpcsecuritygroupids
	//
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

