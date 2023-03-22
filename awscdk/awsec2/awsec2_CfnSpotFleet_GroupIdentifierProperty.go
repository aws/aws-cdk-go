package awsec2


// Describes a security group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   groupIdentifierProperty := &groupIdentifierProperty{
//   	groupId: jsii.String("groupId"),
//   }
//
type CfnSpotFleet_GroupIdentifierProperty struct {
	// The ID of the security group.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
}

