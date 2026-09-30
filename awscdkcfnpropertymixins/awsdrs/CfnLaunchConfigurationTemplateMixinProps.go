package awsdrs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnLaunchConfigurationTemplatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnLaunchConfigurationTemplateMixinProps := &CfnLaunchConfigurationTemplateMixinProps{
//   	CopyPrivateIp: jsii.Boolean(false),
//   	CopyTags: jsii.Boolean(false),
//   	ExportBucketArn: jsii.String("exportBucketArn"),
//   	LaunchDisposition: jsii.String("launchDisposition"),
//   	LaunchIntoSourceInstance: jsii.Boolean(false),
//   	Licensing: &LicensingProperty{
//   		OsByol: jsii.Boolean(false),
//   	},
//   	PostLaunchEnabled: jsii.Boolean(false),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetInstanceTypeRightSizingMethod: jsii.String("targetInstanceTypeRightSizingMethod"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-launchconfigurationtemplate.html
//
type CfnLaunchConfigurationTemplateMixinProps struct {
	// Copy private IP.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-launchconfigurationtemplate.html#cfn-drs-launchconfigurationtemplate-copyprivateip
	//
	CopyPrivateIp interface{} `field:"optional" json:"copyPrivateIp" yaml:"copyPrivateIp"`
	// Copy tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-launchconfigurationtemplate.html#cfn-drs-launchconfigurationtemplate-copytags
	//
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// S3 bucket ARN to export Source Network templates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-launchconfigurationtemplate.html#cfn-drs-launchconfigurationtemplate-exportbucketarn
	//
	ExportBucketArn *string `field:"optional" json:"exportBucketArn" yaml:"exportBucketArn"`
	// Launch disposition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-launchconfigurationtemplate.html#cfn-drs-launchconfigurationtemplate-launchdisposition
	//
	LaunchDisposition *string `field:"optional" json:"launchDisposition" yaml:"launchDisposition"`
	// DRS will set the 'launch into instance ID' of any source server when performing a drill, recovery or failback to the previous region or availability zone, using the instance ID of the source instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-launchconfigurationtemplate.html#cfn-drs-launchconfigurationtemplate-launchintosourceinstance
	//
	LaunchIntoSourceInstance interface{} `field:"optional" json:"launchIntoSourceInstance" yaml:"launchIntoSourceInstance"`
	// Configuration of a machine's license.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-launchconfigurationtemplate.html#cfn-drs-launchconfigurationtemplate-licensing
	//
	Licensing interface{} `field:"optional" json:"licensing" yaml:"licensing"`
	// Whether we want to activate post-launch actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-launchconfigurationtemplate.html#cfn-drs-launchconfigurationtemplate-postlaunchenabled
	//
	PostLaunchEnabled interface{} `field:"optional" json:"postLaunchEnabled" yaml:"postLaunchEnabled"`
	// A set of tags associated with the Launch Configuration Template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-launchconfigurationtemplate.html#cfn-drs-launchconfigurationtemplate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Target instance type right-sizing method.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-drs-launchconfigurationtemplate.html#cfn-drs-launchconfigurationtemplate-targetinstancetyperightsizingmethod
	//
	TargetInstanceTypeRightSizingMethod *string `field:"optional" json:"targetInstanceTypeRightSizingMethod" yaml:"targetInstanceTypeRightSizingMethod"`
}

