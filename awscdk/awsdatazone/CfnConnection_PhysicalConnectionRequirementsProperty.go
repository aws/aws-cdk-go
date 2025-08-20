package awsdatazone


// Physical connection requirements of a connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   physicalConnectionRequirementsProperty := &PhysicalConnectionRequirementsProperty{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	SecurityGroupIdList: []*string{
//   		jsii.String("securityGroupIdList"),
//   	},
//   	SubnetId: jsii.String("subnetId"),
//   	SubnetIdList: []*string{
//   		jsii.String("subnetIdList"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-physicalconnectionrequirements.html
//
type CfnConnection_PhysicalConnectionRequirementsProperty struct {
	// The availability zone of the physical connection requirements of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-physicalconnectionrequirements.html#cfn-datazone-connection-physicalconnectionrequirements-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The group ID list of the physical connection requirements of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-physicalconnectionrequirements.html#cfn-datazone-connection-physicalconnectionrequirements-securitygroupidlist
	//
	SecurityGroupIdList *[]*string `field:"optional" json:"securityGroupIdList" yaml:"securityGroupIdList"`
	// The subnet ID of the physical connection requirements of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-physicalconnectionrequirements.html#cfn-datazone-connection-physicalconnectionrequirements-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// The subnet ID list of the physical connection requirements of a connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-physicalconnectionrequirements.html#cfn-datazone-connection-physicalconnectionrequirements-subnetidlist
	//
	SubnetIdList *[]*string `field:"optional" json:"subnetIdList" yaml:"subnetIdList"`
}

