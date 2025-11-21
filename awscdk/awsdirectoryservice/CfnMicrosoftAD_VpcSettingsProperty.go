package awsdirectoryservice


// Contains VPC information for the [CreateDirectory](https://docs.aws.amazon.com/directoryservice/latest/devguide/API_CreateDirectory.html) or [CreateMicrosoftAD](https://docs.aws.amazon.com/directoryservice/latest/devguide/API_CreateMicrosoftAD.html) operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcSettingsProperty := &VpcSettingsProperty{
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-microsoftad-vpcsettings.html
//
type CfnMicrosoftAD_VpcSettingsProperty struct {
	// The identifiers of the subnets for the directory servers.
	//
	// The two subnets must be in different Availability Zones. Directory Service specifies a directory server and a DNS server in each of these subnets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-microsoftad-vpcsettings.html#cfn-directoryservice-microsoftad-vpcsettings-subnetids
	//
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The identifier of the VPC in which to create the directory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-directoryservice-microsoftad-vpcsettings.html#cfn-directoryservice-microsoftad-vpcsettings-vpcid
	//
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

