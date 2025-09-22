package awsec2


// A reference to a SecurityGroupVpcAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityGroupVpcAssociationReference := &SecurityGroupVpcAssociationReference{
//   	GroupId: jsii.String("groupId"),
//   	VpcId: jsii.String("vpcId"),
//   }
//
type SecurityGroupVpcAssociationReference struct {
	// The GroupId of the SecurityGroupVpcAssociation resource.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// The VpcId of the SecurityGroupVpcAssociation resource.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

