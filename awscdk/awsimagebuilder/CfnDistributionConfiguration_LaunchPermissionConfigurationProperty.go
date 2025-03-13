package awsimagebuilder


// Describes the configuration for a launch permission.
//
// The launch permission modification request is sent to the [Amazon EC2 ModifyImageAttribute](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyImageAttribute.html) API on behalf of the user for each Region they have selected to distribute the AMI. To make an AMI public, set the launch permission authorized accounts to `all` . See the examples for making an AMI public at [Amazon EC2 ModifyImageAttribute](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyImageAttribute.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchPermissionConfigurationProperty := &LaunchPermissionConfigurationProperty{
//   	OrganizationalUnitArns: []*string{
//   		jsii.String("organizationalUnitArns"),
//   	},
//   	OrganizationArns: []*string{
//   		jsii.String("organizationArns"),
//   	},
//   	UserGroups: []*string{
//   		jsii.String("userGroups"),
//   	},
//   	UserIds: []*string{
//   		jsii.String("userIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-launchpermissionconfiguration.html
//
type CfnDistributionConfiguration_LaunchPermissionConfigurationProperty struct {
	// The ARN for an AWS Organizations organizational unit (OU) that you want to share your AMI with.
	//
	// For more information about key concepts for AWS Organizations , see [AWS Organizations terminology and concepts](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-launchpermissionconfiguration.html#cfn-imagebuilder-distributionconfiguration-launchpermissionconfiguration-organizationalunitarns
	//
	OrganizationalUnitArns *[]*string `field:"optional" json:"organizationalUnitArns" yaml:"organizationalUnitArns"`
	// The ARN for an AWS Organization that you want to share your AMI with.
	//
	// For more information, see [What is AWS Organizations ?](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_introduction.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-launchpermissionconfiguration.html#cfn-imagebuilder-distributionconfiguration-launchpermissionconfiguration-organizationarns
	//
	OrganizationArns *[]*string `field:"optional" json:"organizationArns" yaml:"organizationArns"`
	// The name of the group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-launchpermissionconfiguration.html#cfn-imagebuilder-distributionconfiguration-launchpermissionconfiguration-usergroups
	//
	UserGroups *[]*string `field:"optional" json:"userGroups" yaml:"userGroups"`
	// The AWS account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-launchpermissionconfiguration.html#cfn-imagebuilder-distributionconfiguration-launchpermissionconfiguration-userids
	//
	UserIds *[]*string `field:"optional" json:"userIds" yaml:"userIds"`
}

