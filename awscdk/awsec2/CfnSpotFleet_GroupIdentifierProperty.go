package awsec2


// Describes a security group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupIdentifierProperty := &GroupIdentifierProperty{
//   	GroupId: jsii.String("groupId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-groupidentifier.html
//
type CfnSpotFleet_GroupIdentifierProperty struct {
	// The ID of the security group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-groupidentifier.html#cfn-ec2-spotfleet-groupidentifier-groupid
	//
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
}

