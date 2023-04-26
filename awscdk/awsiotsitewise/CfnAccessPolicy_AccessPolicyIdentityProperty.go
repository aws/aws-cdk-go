package awsiotsitewise


// The identity (IAM Identity Center user, IAM Identity Center group, or IAM user) to which this access policy applies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnAccessPolicy_AccessPolicyIdentityProperty struct {
	// An IAM role identity.
	IamRole interface{} `field:"optional" json:"iamRole" yaml:"iamRole"`
	// An IAM user identity.
	IamUser interface{} `field:"optional" json:"iamUser" yaml:"iamUser"`
	// The IAM Identity Center user to which this access policy maps.
	User interface{} `field:"optional" json:"user" yaml:"user"`
}

