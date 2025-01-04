package awscognito


// Properties for defining a `CfnUserPoolGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnUserPoolGroupProps := &CfnUserPoolGroupProps{
//   	UserPoolId: jsii.String("userPoolId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	GroupName: jsii.String("groupName"),
//   	Precedence: jsii.Number(123),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html
//
type CfnUserPoolGroupProps struct {
	// The ID of the user pool where you want to create a user group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html#cfn-cognito-userpoolgroup-userpoolid
	//
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// A description of the group that you're creating.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html#cfn-cognito-userpoolgroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the group.
	//
	// This name must be unique in your user pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html#cfn-cognito-userpoolgroup-groupname
	//
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// A non-negative integer value that specifies the precedence of this group relative to the other groups that a user can belong to in the user pool.
	//
	// Zero is the highest precedence value. Groups with lower `Precedence` values take precedence over groups with higher or null `Precedence` values. If a user belongs to two or more groups, it is the group with the lowest precedence value whose role ARN is given in the user's tokens for the `cognito:roles` and `cognito:preferred_role` claims.
	//
	// Two groups can have the same `Precedence` value. If this happens, neither group takes precedence over the other. If two groups with the same `Precedence` have the same role ARN, that role is used in the `cognito:preferred_role` claim in tokens for users in each group. If the two groups have different role ARNs, the `cognito:preferred_role` claim isn't set in users' tokens.
	//
	// The default `Precedence` value is null. The maximum `Precedence` value is `2^31-1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html#cfn-cognito-userpoolgroup-precedence
	//
	Precedence *float64 `field:"optional" json:"precedence" yaml:"precedence"`
	// The Amazon Resource Name (ARN) for the IAM role that you want to associate with the group.
	//
	// A group role primarily declares a preferred role for the credentials that you get from an identity pool. Amazon Cognito ID tokens have a `cognito:preferred_role` claim that presents the highest-precedence group that a user belongs to. Both ID and access tokens also contain a `cognito:groups` claim that list all the groups that a user is a member of.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolgroup.html#cfn-cognito-userpoolgroup-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

