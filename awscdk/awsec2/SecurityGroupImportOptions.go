package awsec2


// Additional options for imported security groups.
//
// Example:
//   securityGroup := ec2.SecurityGroup_FromSecurityGroupId(this, jsii.String("SG"), jsii.String("sg-12345"), &SecurityGroupImportOptions{
//   	Mutable: jsii.Boolean(false),
//   })
//
type SecurityGroupImportOptions struct {
	// Mark the SecurityGroup as having been created allowing all outbound ipv6 traffic.
	//
	// Only if this is set to false will egress rules for ipv6 be added to this security
	// group. Be aware, this would undo any potential "all outbound traffic"
	// default.
	// Default: false.
	//
	AllowAllIpv6Outbound *bool `field:"optional" json:"allowAllIpv6Outbound" yaml:"allowAllIpv6Outbound"`
	// Mark the SecurityGroup as having been created allowing all outbound traffic.
	//
	// Only if this is set to false will egress rules be added to this security
	// group. Be aware, this would undo any potential "all outbound traffic"
	// default.
	// Default: true.
	//
	AllowAllOutbound *bool `field:"optional" json:"allowAllOutbound" yaml:"allowAllOutbound"`
	// If a SecurityGroup is mutable CDK can add rules to existing groups.
	//
	// Beware that making a SecurityGroup immutable might lead to issue
	// due to missing ingress/egress rules for new resources.
	// Default: true.
	//
	Mutable *bool `field:"optional" json:"mutable" yaml:"mutable"`
}

