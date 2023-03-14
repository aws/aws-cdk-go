package awselasticache


// Properties for defining a `CfnUserGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserGroupProps := &cfnUserGroupProps{
//   	engine: jsii.String("engine"),
//   	userGroupId: jsii.String("userGroupId"),
//
//   	// the properties below are optional
//   	userIds: []*string{
//   		jsii.String("userIds"),
//   	},
//   }
//
type CfnUserGroupProps struct {
	// The current supported value is redis.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// The ID of the user group.
	UserGroupId *string `field:"required" json:"userGroupId" yaml:"userGroupId"`
	// The list of user IDs that belong to the user group.
	//
	// A user named `default` must be included.
	UserIds *[]*string `field:"optional" json:"userIds" yaml:"userIds"`
}

