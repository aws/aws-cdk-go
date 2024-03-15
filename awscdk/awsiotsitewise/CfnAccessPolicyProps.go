package awsiotsitewise


// Properties for defining a `CfnAccessPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAccessPolicyProps := &CfnAccessPolicyProps{
//   	AccessPolicyIdentity: &AccessPolicyIdentityProperty{
//   		IamRole: &IamRoleProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   		IamUser: &IamUserProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   		User: &UserProperty{
//   			Id: jsii.String("id"),
//   		},
//   	},
//   	AccessPolicyPermission: jsii.String("accessPolicyPermission"),
//   	AccessPolicyResource: &AccessPolicyResourceProperty{
//   		Portal: &PortalProperty{
//   			Id: jsii.String("id"),
//   		},
//   		Project: &ProjectProperty{
//   			Id: jsii.String("id"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-accesspolicy.html
//
type CfnAccessPolicyProps struct {
	// The identity for this access policy.
	//
	// Choose an IAM Identity Center user, an IAM Identity Center group, or an IAM user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-accesspolicy.html#cfn-iotsitewise-accesspolicy-accesspolicyidentity
	//
	AccessPolicyIdentity interface{} `field:"required" json:"accessPolicyIdentity" yaml:"accessPolicyIdentity"`
	// The permission level for this access policy.
	//
	// Note that a project `ADMINISTRATOR` is also known as a project owner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-accesspolicy.html#cfn-iotsitewise-accesspolicy-accesspolicypermission
	//
	AccessPolicyPermission *string `field:"required" json:"accessPolicyPermission" yaml:"accessPolicyPermission"`
	// The AWS IoT SiteWise Monitor resource for this access policy.
	//
	// Choose either a portal or a project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-accesspolicy.html#cfn-iotsitewise-accesspolicy-accesspolicyresource
	//
	AccessPolicyResource interface{} `field:"required" json:"accessPolicyResource" yaml:"accessPolicyResource"`
}

