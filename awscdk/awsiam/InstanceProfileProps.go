package awsiam


// Properties of an Instance Profile.
//
// Example:
//   role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("ec2.amazonaws.com")),
//   })
//
//   instanceProfile := iam.NewInstanceProfile(this, jsii.String("InstanceProfile"), &InstanceProfileProps{
//   	Role: Role,
//   	InstanceProfileName: jsii.String("MyInstanceProfile"),
//   	Path: jsii.String("/sample/path/"),
//   })
//
type InstanceProfileProps struct {
	// The name of the InstanceProfile to create.
	// Default: - generated by CloudFormation.
	//
	InstanceProfileName *string `field:"optional" json:"instanceProfileName" yaml:"instanceProfileName"`
	// The path to the InstanceProfile.
	// Default: /.
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
	// An IAM role to associate with the instance profile that is used by EC2 instances.
	//
	// The role must be assumable by the service principal `ec2.amazonaws.com`:
	//
	// Example:
	//   role := iam.NewRole(this, jsii.String("MyRole"), &RoleProps{
	//   	AssumedBy: iam.NewServicePrincipal(jsii.String("ec2.amazonaws.com")),
	//   })
	//
	// Default: - a role will be automatically created, it can be accessed via the `role` property.
	//
	Role IRole `field:"optional" json:"role" yaml:"role"`
}

