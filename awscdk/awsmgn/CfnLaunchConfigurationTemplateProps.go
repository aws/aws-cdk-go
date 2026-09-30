package awsmgn

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLaunchConfigurationTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLaunchConfigurationTemplateProps := &CfnLaunchConfigurationTemplateProps{
//   	AssociatePublicIpAddress: jsii.Boolean(false),
//   	BootMode: jsii.String("bootMode"),
//   	CopyPrivateIp: jsii.Boolean(false),
//   	CopyTags: jsii.Boolean(false),
//   	EnableMapAutoTagging: jsii.Boolean(false),
//   	EnableParametersEncryption: jsii.Boolean(false),
//   	LargeVolumeConf: &LaunchTemplateDiskConfProperty{
//   		Iops: jsii.Number(123),
//   		Throughput: jsii.Number(123),
//   		VolumeType: jsii.String("volumeType"),
//   	},
//   	LaunchDisposition: jsii.String("launchDisposition"),
//   	Licensing: &LicensingProperty{
//   		OsByol: jsii.Boolean(false),
//   	},
//   	MapAutoTaggingMpeId: jsii.String("mapAutoTaggingMpeId"),
//   	ParametersEncryptionKey: jsii.String("parametersEncryptionKey"),
//   	PostLaunchActions: &PostLaunchActionsProperty{
//   		CloudWatchLogGroupName: jsii.String("cloudWatchLogGroupName"),
//   		Deployment: jsii.String("deployment"),
//   		S3LogBucket: jsii.String("s3LogBucket"),
//   		S3OutputKeyPrefix: jsii.String("s3OutputKeyPrefix"),
//   		SsmDocuments: []interface{}{
//   			&SsmDocumentProperty{
//   				ActionName: jsii.String("actionName"),
//   				SsmDocumentName: jsii.String("ssmDocumentName"),
//
//   				// the properties below are optional
//   				ExternalParameters: map[string]interface{}{
//   					"externalParametersKey": &SsmExternalParameterProperty{
//   						"dynamicPath": jsii.String("dynamicPath"),
//   					},
//   				},
//   				MustSucceedForCutover: jsii.Boolean(false),
//   				Parameters: map[string]interface{}{
//   					"parametersKey": []interface{}{
//   						&SsmParameterStoreParameterProperty{
//   							"parameterName": jsii.String("parameterName"),
//   							"parameterType": jsii.String("parameterType"),
//   						},
//   					},
//   				},
//   				TimeoutSeconds: jsii.Number(123),
//   			},
//   		},
//   	},
//   	SmallVolumeConf: &LaunchTemplateDiskConfProperty{
//   		Iops: jsii.Number(123),
//   		Throughput: jsii.Number(123),
//   		VolumeType: jsii.String("volumeType"),
//   	},
//   	SmallVolumeMaxSize: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetInstanceTypeRightSizingMethod: jsii.String("targetInstanceTypeRightSizingMethod"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-launchconfigurationtemplate.html
//
type CfnLaunchConfigurationTemplateProps struct {
	// Whether to associate a public IP address with the launched instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-launchconfigurationtemplate.html#cfn-mgn-launchconfigurationtemplate-associatepublicipaddress
	//
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// Launch configuration template boot mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-launchconfigurationtemplate.html#cfn-mgn-launchconfigurationtemplate-bootmode
	//
	BootMode *string `field:"optional" json:"bootMode" yaml:"bootMode"`
	// Whether to copy the private IP of the source server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-launchconfigurationtemplate.html#cfn-mgn-launchconfigurationtemplate-copyprivateip
	//
	CopyPrivateIp interface{} `field:"optional" json:"copyPrivateIp" yaml:"copyPrivateIp"`
	// Whether to copy the tags of the source server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-launchconfigurationtemplate.html#cfn-mgn-launchconfigurationtemplate-copytags
	//
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// Whether to enable map auto tagging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-launchconfigurationtemplate.html#cfn-mgn-launchconfigurationtemplate-enablemapautotagging
	//
	EnableMapAutoTagging interface{} `field:"optional" json:"enableMapAutoTagging" yaml:"enableMapAutoTagging"`
	// Whether to enable encryption of the AWS Systems Manager parameters used by post launch actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-launchconfigurationtemplate.html#cfn-mgn-launchconfigurationtemplate-enableparametersencryption
	//
	EnableParametersEncryption interface{} `field:"optional" json:"enableParametersEncryption" yaml:"enableParametersEncryption"`
	// Launch template disk configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-launchconfigurationtemplate.html#cfn-mgn-launchconfigurationtemplate-largevolumeconf
	//
	LargeVolumeConf interface{} `field:"optional" json:"largeVolumeConf" yaml:"largeVolumeConf"`
	// Launch disposition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-launchconfigurationtemplate.html#cfn-mgn-launchconfigurationtemplate-launchdisposition
	//
	LaunchDisposition *string `field:"optional" json:"launchDisposition" yaml:"launchDisposition"`
	// Configuration of a machine's license.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-launchconfigurationtemplate.html#cfn-mgn-launchconfigurationtemplate-licensing
	//
	Licensing interface{} `field:"optional" json:"licensing" yaml:"licensing"`
	// Launch configuration template map auto tagging MPE ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-launchconfigurationtemplate.html#cfn-mgn-launchconfigurationtemplate-mapautotaggingmpeid
	//
	MapAutoTaggingMpeId *string `field:"optional" json:"mapAutoTaggingMpeId" yaml:"mapAutoTaggingMpeId"`
	// ARN of the KMS key used to encrypt the AWS Systems Manager parameters used by post launch actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-launchconfigurationtemplate.html#cfn-mgn-launchconfigurationtemplate-parametersencryptionkey
	//
	ParametersEncryptionKey *string `field:"optional" json:"parametersEncryptionKey" yaml:"parametersEncryptionKey"`
	// Post launch actions to execute on the Test or Cutover instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-launchconfigurationtemplate.html#cfn-mgn-launchconfigurationtemplate-postlaunchactions
	//
	PostLaunchActions interface{} `field:"optional" json:"postLaunchActions" yaml:"postLaunchActions"`
	// Launch template disk configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-launchconfigurationtemplate.html#cfn-mgn-launchconfigurationtemplate-smallvolumeconf
	//
	SmallVolumeConf interface{} `field:"optional" json:"smallVolumeConf" yaml:"smallVolumeConf"`
	// Small volume maximum size, in GiB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-launchconfigurationtemplate.html#cfn-mgn-launchconfigurationtemplate-smallvolumemaxsize
	//
	SmallVolumeMaxSize *float64 `field:"optional" json:"smallVolumeMaxSize" yaml:"smallVolumeMaxSize"`
	// A set of tags to be associated with the Launch Configuration Template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-launchconfigurationtemplate.html#cfn-mgn-launchconfigurationtemplate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Target instance type right-sizing method.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mgn-launchconfigurationtemplate.html#cfn-mgn-launchconfigurationtemplate-targetinstancetyperightsizingmethod
	//
	TargetInstanceTypeRightSizingMethod *string `field:"optional" json:"targetInstanceTypeRightSizingMethod" yaml:"targetInstanceTypeRightSizingMethod"`
}

