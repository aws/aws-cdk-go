package cxapi


// Properties of a discovered SecurityGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityGroupContextResponse := &securityGroupContextResponse{
//   	allowAllOutbound: jsii.Boolean(false),
//   	securityGroupId: jsii.String("securityGroupId"),
//   }
//
type SecurityGroupContextResponse struct {
	// Whether the security group allows all outbound traffic.
	//
	// This will be true
	// when the security group has all-protocol egress permissions to access both
	// `0.0.0.0/0` and `::/0`.
	AllowAllOutbound *bool `field:"required" json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// The security group's id.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
}

