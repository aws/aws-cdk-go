package awsec2


// Optional parameters used when creating a user.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   initUserOptions := &initUserOptions{
//   	groups: []*string{
//   		jsii.String("groups"),
//   	},
//   	homeDir: jsii.String("homeDir"),
//   	userId: jsii.Number(123),
//   }
//
type InitUserOptions struct {
	// A list of group names.
	//
	// The user will be added to each group in the list.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// The user's home directory.
	HomeDir *string `field:"optional" json:"homeDir" yaml:"homeDir"`
	// A user ID.
	//
	// The creation process fails if the user name exists with a different user ID.
	// If the user ID is already assigned to an existing user the operating system may
	// reject the creation request.
	UserId *float64 `field:"optional" json:"userId" yaml:"userId"`
}

