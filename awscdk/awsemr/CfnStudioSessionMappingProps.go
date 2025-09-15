package awsemr


// Properties for defining a `CfnStudioSessionMapping`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnStudioSessionMappingProps := &CfnStudioSessionMappingProps{
//   	IdentityName: jsii.String("identityName"),
//   	IdentityType: jsii.String("identityType"),
//   	SessionPolicyArn: jsii.String("sessionPolicyArn"),
//   	StudioId: jsii.String("studioId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studiosessionmapping.html
//
type CfnStudioSessionMappingProps struct {
	// The name of the user or group.
	//
	// For more information, see [UserName](https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_User.html#singlesignon-Type-User-UserName) and [DisplayName](https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_Group.html#singlesignon-Type-Group-DisplayName) in the *IAM Identity Center Identity Store API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studiosessionmapping.html#cfn-emr-studiosessionmapping-identityname
	//
	IdentityName *string `field:"required" json:"identityName" yaml:"identityName"`
	// Specifies whether the identity to map to the Amazon EMR Studio is a user or a group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studiosessionmapping.html#cfn-emr-studiosessionmapping-identitytype
	//
	IdentityType *string `field:"required" json:"identityType" yaml:"identityType"`
	// The Amazon Resource Name (ARN) for the session policy that will be applied to the user or group.
	//
	// Session policies refine Studio user permissions without the need to use multiple IAM user roles. For more information, see [Create an EMR Studio user role with session policies](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-studio-user-role.html) in the *Amazon EMR Management Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studiosessionmapping.html#cfn-emr-studiosessionmapping-sessionpolicyarn
	//
	SessionPolicyArn *string `field:"required" json:"sessionPolicyArn" yaml:"sessionPolicyArn"`
	// The ID of the Amazon EMR Studio to which the user or group will be mapped.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-studiosessionmapping.html#cfn-emr-studiosessionmapping-studioid
	//
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
}

