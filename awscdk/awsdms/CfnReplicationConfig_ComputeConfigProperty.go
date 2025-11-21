package awsdms


// Configuration parameters for provisioning an AWS DMS Serverless replication.
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
	// Specifies the maximum value of the AWS DMS capacity units (DCUs) for which a given AWS DMS Serverless replication can be provisioned.
	//
	// A single DCU is 2GB of RAM, with 1 DCU as the minimum value allowed. The list of valid DCU values includes 1, 2, 4, 8, 16, 32, 64, 128, 192, 256, and 384. So, the maximum value that you can specify for AWS DMS Serverless is 384. The `MaxCapacityUnits` parameter is the only DCU parameter you are required to specify.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationconfig-computeconfig.html#cfn-dms-replicationconfig-computeconfig-maxcapacityunits
	//
	MaxCapacityUnits *float64 `field:"required" json:"maxCapacityUnits" yaml:"maxCapacityUnits"`
	// The Availability Zone where the AWS DMS Serverless replication using this configuration will run.
	//
	// The default value is a random, system-chosen Availability Zone in the configuration's AWS Region , for example, `"us-west-2"` . You can't set this parameter if the `MultiAZ` parameter is set to `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationconfig-computeconfig.html#cfn-dms-replicationconfig-computeconfig-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// A list of custom DNS name servers supported for the AWS DMS Serverless replication to access your source or target database.
	//
	// This list overrides the default name servers supported by the AWS DMS Serverless replication. You can specify a comma-separated list of internet addresses for up to four DNS name servers. For example: `"1.1.1.1,2.2.2.2,3.3.3.3,4.4.4.4"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationconfig-computeconfig.html#cfn-dms-replicationconfig-computeconfig-dnsnameservers
	//
	DnsNameServers *string `field:"optional" json:"dnsNameServers" yaml:"dnsNameServers"`
	// An AWS Key Management Service ( AWS  ) key Amazon Resource Name (ARN) that is used to encrypt the data during AWS DMS Serverless replication.
	//
	// If you don't specify a value for the `KmsKeyId` parameter, AWS DMS uses your default encryption key.
	//
	// AWS  creates the default encryption key for your Amazon Web Services account. Your AWS account has a different default encryption key for each AWS Region .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationconfig-computeconfig.html#cfn-dms-replicationconfig-computeconfig-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the minimum value of the AWS DMS capacity units (DCUs) for which a given AWS DMS Serverless replication can be provisioned.
	//
	// A single DCU is 2GB of RAM, with 1 DCU as the minimum value allowed. The list of valid DCU values includes 1, 2, 4, 8, 16, 32, 64, 128, 192, 256, and 384. So, the minimum DCU value that you can specify for AWS DMS Serverless is 1. If you don't set this value, AWS DMS sets this parameter to the minimum DCU value allowed, 1. If there is no current source activity, AWS DMS scales down your replication until it reaches the value specified in `MinCapacityUnits` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationconfig-computeconfig.html#cfn-dms-replicationconfig-computeconfig-mincapacityunits
	//
	MinCapacityUnits *float64 `field:"optional" json:"minCapacityUnits" yaml:"minCapacityUnits"`
	// Specifies whether the AWS DMS Serverless replication is a Multi-AZ deployment.
	//
	// You can't set the `AvailabilityZone` parameter if the `MultiAZ` parameter is set to `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationconfig-computeconfig.html#cfn-dms-replicationconfig-computeconfig-multiaz
	//
	MultiAz interface{} `field:"optional" json:"multiAz" yaml:"multiAz"`
	// The weekly time range during which system maintenance can occur for the AWS DMS Serverless replication, in Universal Coordinated Time (UTC).
	//
	// The format is `ddd:hh24:mi-ddd:hh24:mi` .
	//
	// The default is a 30-minute window selected at random from an 8-hour block of time per AWS Region . This maintenance occurs on a random day of the week. Valid values for days of the week include `Mon` , `Tue` , `Wed` , `Thu` , `Fri` , `Sat` , and `Sun` .
	//
	// Constraints include a minimum 30-minute window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationconfig-computeconfig.html#cfn-dms-replicationconfig-computeconfig-preferredmaintenancewindow
	//
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Specifies a subnet group identifier to associate with the AWS DMS Serverless replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationconfig-computeconfig.html#cfn-dms-replicationconfig-computeconfig-replicationsubnetgroupid
	//
	ReplicationSubnetGroupId *string `field:"optional" json:"replicationSubnetGroupId" yaml:"replicationSubnetGroupId"`
	// Specifies the virtual private cloud (VPC) security group to use with the AWS DMS Serverless replication.
	//
	// The VPC security group must work with the VPC containing the replication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-replicationconfig-computeconfig.html#cfn-dms-replicationconfig-computeconfig-vpcsecuritygroupids
	//
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

