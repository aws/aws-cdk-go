package awsram


// The configuration for a resource share.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceShareConfigurationProperty := &ResourceShareConfigurationProperty{
//   	ExclusiveAccountAccess: jsii.Boolean(false),
//   	RetainSharingOnAccountLeaveOrganization: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ram-resourceshare-resourceshareconfiguration.html
//
type CfnResourceShare_ResourceShareConfigurationProperty struct {
	// The resource share restricts access to an account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ram-resourceshare-resourceshareconfiguration.html#cfn-ram-resourceshare-resourceshareconfiguration-exclusiveaccountaccess
	//
	ExclusiveAccountAccess interface{} `field:"optional" json:"exclusiveAccountAccess" yaml:"exclusiveAccountAccess"`
	// Specifies whether the consumer account retains access to the resource share after leaving the organization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ram-resourceshare-resourceshareconfiguration.html#cfn-ram-resourceshare-resourceshareconfiguration-retainsharingonaccountleaveorganization
	//
	RetainSharingOnAccountLeaveOrganization interface{} `field:"optional" json:"retainSharingOnAccountLeaveOrganization" yaml:"retainSharingOnAccountLeaveOrganization"`
}

