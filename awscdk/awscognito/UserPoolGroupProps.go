package awscognito

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam"
)

// Props for UserPoolGroup construct.
//
// Example:
//   var userPool UserPool
//   var role Role
//
//
//   cognito.NewUserPoolGroup(this, jsii.String("UserPoolGroup"), &UserPoolGroupProps{
//   	UserPool: UserPool,
//   	GroupName: jsii.String("my-group-name"),
//   	Precedence: jsii.Number(1),
//   	Role: Role,
//   })
//
//   // You can also add a group by using addGroup method.
//   userPool.addGroup(jsii.String("AnotherUserPoolGroup"), &UserPoolGroupOptions{
//   	GroupName: jsii.String("another-group-name"),
//   })
//
type UserPoolGroupProps struct {
	// A string containing the description of the group.
	// Default: - no description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the group.
	//
	// Must be unique.
	// Default: - auto generate a name.
	//
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// A non-negative integer value that specifies the precedence of this group relative to the other groups that a user can belong to in the user pool.
	//
	// Zero is the highest precedence value.
	//
	// Groups with lower Precedence values take precedence over groups with higher or null Precedence values.
	// If a user belongs to two or more groups, it is the group with the lowest precedence value
	// whose role ARN is given in the user's tokens for the cognito:roles and cognito:preferred_role claims.
	//
	// Two groups can have the same Precedence value. If this happens, neither group takes precedence over the other.
	// If two groups with the same Precedence have the same role ARN, that role is used in the cognito:preferred_role
	// claim in tokens for users in each group.
	// If the two groups have different role ARNs, the cognito:preferred_role claim isn't set in users' tokens.
	// Default: - null.
	//
	Precedence *float64 `field:"optional" json:"precedence" yaml:"precedence"`
	// The role for the group.
	// Default: - no description.
	//
	Role interfacesawsiam.IRoleRef `field:"optional" json:"role" yaml:"role"`
	// The user pool to which this group is associated.
	UserPool interfacesawscognito.IUserPoolRef `field:"required" json:"userPool" yaml:"userPool"`
}

