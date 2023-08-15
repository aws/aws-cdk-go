package awsiam


// Attributes of an Instance Profile.
//
// Example:
//   role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("ec2.amazonaws.com")),
//   })
//
//   instanceProfile := iam.InstanceProfile_FromInstanceProfileAttributes(this, jsii.String("ImportedInstanceProfile"), &InstanceProfileAttributes{
//   	InstanceProfileArn: jsii.String("arn:aws:iam::account-id:instance-profile/MyInstanceProfile"),
//   	Role: Role,
//   })
//
type InstanceProfileAttributes struct {
	// The ARN of the InstanceProfile.
	//
	// Format: arn:<partition>:iam::<account-id>:instance-profile/<instance-profile-name-with-path>.
	InstanceProfileArn *string `field:"required" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// The role associated with the InstanceProfile.
	// Default: - no role.
	//
	Role IRole `field:"optional" json:"role" yaml:"role"`
}

