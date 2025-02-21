package awsiam


// Properties for defining a `CfnUserToGroupAddition`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserToGroupAdditionProps := &CfnUserToGroupAdditionProps{
//   	GroupName: jsii.String("groupName"),
//   	Users: []*string{
//   		jsii.String("users"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-usertogroupaddition.html
//
type CfnUserToGroupAdditionProps struct {
	// The name of the group to update.
	//
	// This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) ) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-usertogroupaddition.html#cfn-iam-usertogroupaddition-groupname
	//
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// A list of the names of the users that you want to add to the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-usertogroupaddition.html#cfn-iam-usertogroupaddition-users
	//
	Users *[]*string `field:"required" json:"users" yaml:"users"`
}

