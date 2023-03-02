package awsmedialive


// Settings that apply only if the input is an push input where the source is on Amazon VPC.
//
// The parent of this entity is Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputVpcRequestProperty := &inputVpcRequestProperty{
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   }
//
type CfnInput_InputVpcRequestProperty struct {
	// The list of up to five VPC security group IDs to attach to the input VPC network interfaces.
	//
	// The security groups require subnet IDs. If none are specified, MediaLive uses the VPC default security group.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The list of two VPC subnet IDs from the same VPC.
	//
	// You must associate subnet IDs to two unique Availability Zones.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

