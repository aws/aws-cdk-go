package awscognito


// Properties for defining a `CfnUserPoolGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserPoolGroupProps := &cfnUserPoolGroupProps{
//   	userPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	groupName: jsii.String("groupName"),
//   	precedence: jsii.Number(123),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnUserPoolGroupProps struct {
	// The user pool ID for the user pool.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// A string containing the description of the group.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the group.
	//
	// Must be unique.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// A non-negative integer value that specifies the precedence of this group relative to the other groups that a user can belong to in the user pool.
	//
	// Zero is the highest precedence value. Groups with lower `Precedence` values take precedence over groups with higher or null `Precedence` values. If a user belongs to two or more groups, it is the group with the lowest precedence value whose role ARN is given in the user's tokens for the `cognito:roles` and `cognito:preferred_role` claims.
	//
	// Two groups can have the same `Precedence` value. If this happens, neither group takes precedence over the other. If two groups with the same `Precedence` have the same role ARN, that role is used in the `cognito:preferred_role` claim in tokens for users in each group. If the two groups have different role ARNs, the `cognito:preferred_role` claim isn't set in users' tokens.
	//
	// The default `Precedence` value is null. The maximum `Precedence` value is `2^31-1` .
	Precedence *float64 `field:"optional" json:"precedence" yaml:"precedence"`
	// The role Amazon Resource Name (ARN) for the group.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

