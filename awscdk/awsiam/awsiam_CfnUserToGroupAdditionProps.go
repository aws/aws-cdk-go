package awsiam


// Properties for defining a `CfnUserToGroupAddition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserToGroupAdditionProps := &cfnUserToGroupAdditionProps{
//   	groupName: jsii.String("groupName"),
//   	users: []*string{
//   		jsii.String("users"),
//   	},
//   }
//
type CfnUserToGroupAdditionProps struct {
	// The name of the group to update.
	//
	// This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) ) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// A list of the names of the users that you want to add to the group.
	Users *[]*string `field:"required" json:"users" yaml:"users"`
}

