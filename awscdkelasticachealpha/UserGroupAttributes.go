package awscdkelasticachealpha


// Attributes that can be specified when importing a UserGroup.
//
// Example:
//   stack := awscdk.Newstack()
//
//   importedIamUser := elasticache.IamUser_FromUserId(this, jsii.String("ImportedIamUser"), jsii.String("my-iam-user-id"))
//
//   importedPasswordUser := elasticache.PasswordUser_FromUserAttributes(stack, jsii.String("ImportedPasswordUser"), &UserBaseAttributes{
//   	UserId: jsii.String("my-password-user-id"),
//   })
//
//   importedNoPasswordUser := elasticache.NoPasswordUser_FromUserAttributes(stack, jsii.String("ImportedNoPasswordUser"), &UserBaseAttributes{
//   	UserId: jsii.String("my-no-password-user-id"),
//   })
//
//   importedUserGroup := elasticache.UserGroup_FromUserGroupAttributes(this, jsii.String("ImportedUserGroup"), &UserGroupAttributes{
//   	UserGroupName: jsii.String("my-user-group-name"),
//   })
//
// Experimental.
type UserGroupAttributes struct {
	// The engine type for the user group.
	// Default: - engine type is unknown.
	//
	// Experimental.
	Engine UserEngine `field:"optional" json:"engine" yaml:"engine"`
	// The ARN of the user group.
	//
	// One of `userGroupName` or `userGroupArn` is required.
	// Default: - derived from userGroupName.
	//
	// Experimental.
	UserGroupArn *string `field:"optional" json:"userGroupArn" yaml:"userGroupArn"`
	// The name of the user group.
	//
	// One of `userGroupName` or `userGroupArn` is required.
	// Default: - derived from userGroupArn.
	//
	// Experimental.
	UserGroupName *string `field:"optional" json:"userGroupName" yaml:"userGroupName"`
	// List of users in the user group.
	// Default: - users are unknown.
	//
	// Experimental.
	Users *[]IUser `field:"optional" json:"users" yaml:"users"`
}

