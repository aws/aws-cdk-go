package mixinsawsshield


// Properties for CfnDRTAccessPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDRTAccessMixinProps := &CfnDRTAccessMixinProps{
//   	LogBucketList: []*string{
//   		jsii.String("logBucketList"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-shield-drtaccess.html
//
type CfnDRTAccessMixinProps struct {
	// Authorizes the Shield Response Team (SRT) to access the specified Amazon S3 bucket containing log data such as Application Load Balancer access logs, CloudFront logs, or logs from third party sources.
	//
	// You can associate up to 10 Amazon S3 buckets with your subscription.
	//
	// Use this to share information with the SRT that's not available in AWS WAF logs.
	//
	// To use the services of the SRT, you must be subscribed to the [Business Support plan](https://docs.aws.amazon.com/premiumsupport/business-support/) or the [Enterprise Support plan](https://docs.aws.amazon.com/premiumsupport/enterprise-support/) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-shield-drtaccess.html#cfn-shield-drtaccess-logbucketlist
	//
	LogBucketList *[]*string `field:"optional" json:"logBucketList" yaml:"logBucketList"`
	// Authorizes the Shield Response Team (SRT) using the specified role, to access your AWS account to assist with DDoS attack mitigation during potential attacks.
	//
	// This enables the SRT to inspect your AWS WAF configuration and logs and to create or update AWS WAF rules and web ACLs.
	//
	// You can associate only one `RoleArn` with your subscription. If you submit this update for an account that already has an associated role, the new `RoleArn` will replace the existing `RoleArn` .
	//
	// This change requires the following:
	//
	// - You must be subscribed to the [Business Support plan](https://docs.aws.amazon.com/premiumsupport/business-support/) or the [Enterprise Support plan](https://docs.aws.amazon.com/premiumsupport/enterprise-support/) .
	// - The `AWSShieldDRTAccessPolicy` managed policy must be attached to the role that you specify in the request. You can access this policy in the IAM console at [AWSShieldDRTAccessPolicy](https://docs.aws.amazon.com/iam/home?#/policies/arn:aws:iam::aws:policy/service-role/AWSShieldDRTAccessPolicy) . For information, see [Adding and removing IAM identity permissions](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_manage-attach-detach.html) .
	// - The role must trust the service principal `drt.shield.amazonaws.com` . For information, see [IAM JSON policy elements: Principal](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html) .
	//
	// The SRT will have access only to your AWS WAF and Shield resources. By submitting this request, you provide permissions to the SRT to inspect your AWS WAF and Shield configuration and logs, and to create and update AWS WAF rules and web ACLs on your behalf. The SRT takes these actions only if explicitly authorized by you.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-shield-drtaccess.html#cfn-shield-drtaccess-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

