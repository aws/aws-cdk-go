package awsmedialive


// Settings to enable VPC mode in the channel, so that the endpoints for all outputs are in your VPC.
//
// This entity is at the top level in the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcOutputSettingsProperty := &vpcOutputSettingsProperty{
//   	publicAddressAllocationIds: []*string{
//   		jsii.String("publicAddressAllocationIds"),
//   	},
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnChannel_VpcOutputSettingsProperty struct {
	// List of public address allocation IDs to associate with ENIs that will be created in Output VPC.
	//
	// Must specify one for SINGLE_PIPELINE, two for STANDARD channels.
	PublicAddressAllocationIds *[]*string `field:"optional" json:"publicAddressAllocationIds" yaml:"publicAddressAllocationIds"`
	// A list of up to 5 EC2 VPC security group IDs to attach to the Output VPC network interfaces.
	//
	// If none are specified then the VPC default security group will be used.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of VPC subnet IDs from the same VPC.
	//
	// If STANDARD channel, subnet IDs must be mapped to two unique availability zones (AZ).
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

