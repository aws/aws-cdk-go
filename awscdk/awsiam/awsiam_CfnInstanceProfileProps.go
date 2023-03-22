package awsiam


// Properties for defining a `CfnInstanceProfile`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInstanceProfileProps := &cfnInstanceProfileProps{
//   	roles: []*string{
//   		jsii.String("roles"),
//   	},
//
//   	// the properties below are optional
//   	instanceProfileName: jsii.String("instanceProfileName"),
//   	path: jsii.String("path"),
//   }
//
type CfnInstanceProfileProps struct {
	// The name of the role to associate with the instance profile.
	//
	// Only one role can be assigned to an EC2 instance at a time, and all applications on the instance share the same role and permissions.
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
	// The name of the instance profile to create.
	//
	// This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) ) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
	InstanceProfileName *string `field:"optional" json:"instanceProfileName" yaml:"instanceProfileName"`
	// The path to the instance profile.
	//
	// For more information about paths, see [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the *IAM User Guide* .
	//
	// This parameter is optional. If it is not included, it defaults to a slash (/).
	//
	// This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) ) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! ( `\ u0021` ) through the DEL character ( `\ u007F` ), including most punctuation characters, digits, and upper and lowercased letters.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

