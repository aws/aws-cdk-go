package awscdkelasticachealpha


// Properties for defining an ElastiCache UserGroup.
//
// Example:
//   newDefaultUser := elasticache.NewNoPasswordUser(this, jsii.String("NoPasswordUser"), &NoPasswordUserProps{
//   	UserId: jsii.String("default"),
//   	AccessControl: elasticache.AccessControl_FromAccessString(jsii.String("on ~* +@all")),
//   })
//
//   userGroup := elasticache.NewUserGroup(this, jsii.String("UserGroup"), &UserGroupProps{
//   	Users: []IUser{
//   		newDefaultUser,
//   	},
//   })
//
// Experimental.
type UserGroupProps struct {
	// The engine type for the user group Enum options: UserEngine.VALKEY, UserEngine.REDIS.
	// Default: UserEngine.VALKEY
	//
	// Experimental.
	Engine UserEngine `field:"optional" json:"engine" yaml:"engine"`
	// Enforces a particular physical user group name.
	// Default: <generated>.
	//
	// Experimental.
	UserGroupName *string `field:"optional" json:"userGroupName" yaml:"userGroupName"`
	// List of users inside the user group.
	// Default: - no users.
	//
	// Experimental.
	Users *[]IUser `field:"optional" json:"users" yaml:"users"`
}

