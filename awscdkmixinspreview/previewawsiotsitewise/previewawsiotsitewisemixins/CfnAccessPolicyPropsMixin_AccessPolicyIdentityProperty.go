package previewawsiotsitewisemixins


// The identity (IAM Identity Center user, IAM Identity Center group, or IAM user) to which this access policy applies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   accessPolicyIdentityProperty := &AccessPolicyIdentityProperty{
//   	IamRole: &IamRoleProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	IamUser: &IamUserProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	User: &UserProperty{
//   		Id: jsii.String("id"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-accesspolicy-accesspolicyidentity.html
//
type CfnAccessPolicyPropsMixin_AccessPolicyIdentityProperty struct {
	// An IAM role identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-accesspolicy-accesspolicyidentity.html#cfn-iotsitewise-accesspolicy-accesspolicyidentity-iamrole
	//
	IamRole interface{} `field:"optional" json:"iamRole" yaml:"iamRole"`
	// An IAM user identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-accesspolicy-accesspolicyidentity.html#cfn-iotsitewise-accesspolicy-accesspolicyidentity-iamuser
	//
	IamUser interface{} `field:"optional" json:"iamUser" yaml:"iamUser"`
	// An IAM Identity Center user identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-accesspolicy-accesspolicyidentity.html#cfn-iotsitewise-accesspolicy-accesspolicyidentity-user
	//
	User interface{} `field:"optional" json:"user" yaml:"user"`
}

