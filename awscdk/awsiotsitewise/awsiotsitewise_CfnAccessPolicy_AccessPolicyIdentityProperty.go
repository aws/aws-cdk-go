package awsiotsitewise


// The identity ( AWS SSO user, AWS SSO group, or IAM user) to which this access policy applies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessPolicyIdentityProperty := &accessPolicyIdentityProperty{
//   	iamRole: &iamRoleProperty{
//   		arn: jsii.String("arn"),
//   	},
//   	iamUser: &iamUserProperty{
//   		arn: jsii.String("arn"),
//   	},
//   	user: &userProperty{
//   		id: jsii.String("id"),
//   	},
//   }
//
type CfnAccessPolicy_AccessPolicyIdentityProperty struct {
	// An IAM role identity.
	IamRole interface{} `field:"optional" json:"iamRole" yaml:"iamRole"`
	// An IAM user identity.
	IamUser interface{} `field:"optional" json:"iamUser" yaml:"iamUser"`
	// The AWS SSO user to which this access policy maps.
	User interface{} `field:"optional" json:"user" yaml:"user"`
}

