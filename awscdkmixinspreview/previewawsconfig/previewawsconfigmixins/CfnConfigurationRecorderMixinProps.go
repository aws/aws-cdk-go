package previewawsconfigmixins


// Properties for CfnConfigurationRecorderPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfigurationRecorderMixinProps := &CfnConfigurationRecorderMixinProps{
//   	Name: jsii.String("name"),
//   	RecordingGroup: &RecordingGroupProperty{
//   		AllSupported: jsii.Boolean(false),
//   		ExclusionByResourceTypes: &ExclusionByResourceTypesProperty{
//   			ResourceTypes: []*string{
//   				jsii.String("resourceTypes"),
//   			},
//   		},
//   		IncludeGlobalResourceTypes: jsii.Boolean(false),
//   		RecordingStrategy: &RecordingStrategyProperty{
//   			UseOnly: jsii.String("useOnly"),
//   		},
//   		ResourceTypes: []*string{
//   			jsii.String("resourceTypes"),
//   		},
//   	},
//   	RecordingMode: &RecordingModeProperty{
//   		RecordingFrequency: jsii.String("recordingFrequency"),
//   		RecordingModeOverrides: []interface{}{
//   			&RecordingModeOverrideProperty{
//   				Description: jsii.String("description"),
//   				RecordingFrequency: jsii.String("recordingFrequency"),
//   				ResourceTypes: []*string{
//   					jsii.String("resourceTypes"),
//   				},
//   			},
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html
//
type CfnConfigurationRecorderMixinProps struct {
	// The name of the configuration recorder. AWS Config automatically assigns the name of "default" when creating the configuration recorder.
	//
	// You cannot change the name of the configuration recorder after it has been created. To change the configuration recorder name, you must delete it and create a new configuration recorder with a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html#cfn-config-configurationrecorder-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies which resource types AWS Config records for configuration changes.
	//
	// > *High Number of AWS Config Evaluations*
	// >
	// > You may notice increased activity in your account during your initial month recording with AWS Config when compared to subsequent months. During the initial bootstrapping process, AWS Config runs evaluations on all the resources in your account that you have selected for AWS Config to record.
	// >
	// > If you are running ephemeral workloads, you may see increased activity from AWS Config as it records configuration changes associated with creating and deleting these temporary resources. An *ephemeral workload* is a temporary use of computing resources that are loaded and run when needed. Examples include Amazon Elastic Compute Cloud ( Amazon EC2 ) Spot Instances, Amazon EMR jobs, and AWS Auto Scaling . If you want to avoid the increased activity from running ephemeral workloads, you can run these types of workloads in a separate account with AWS Config turned off to avoid increased configuration recording and rule evaluations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html#cfn-config-configurationrecorder-recordinggroup
	//
	RecordingGroup interface{} `field:"optional" json:"recordingGroup" yaml:"recordingGroup"`
	// Specifies the default recording frequency for the configuration recorder. AWS Config supports *Continuous recording* and *Daily recording* .
	//
	// - Continuous recording allows you to record configuration changes continuously whenever a change occurs.
	// - Daily recording allows you to receive a configuration item (CI) representing the most recent state of your resources over the last 24-hour period, only if it’s different from the previous CI recorded.
	//
	// > *Some resource types require continuous recording*
	// >
	// > AWS Firewall Manager depends on continuous recording to monitor your resources. If you are using Firewall Manager, it is recommended that you set the recording frequency to Continuous.
	//
	// You can also override the recording frequency for specific resource types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html#cfn-config-configurationrecorder-recordingmode
	//
	RecordingMode interface{} `field:"optional" json:"recordingMode" yaml:"recordingMode"`
	// Amazon Resource Name (ARN) of the IAM role assumed by AWS Config and used by the configuration recorder.
	//
	// For more information, see [Permissions for the IAM Role Assigned](https://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) to AWS Config in the AWS Config Developer Guide.
	//
	// > *Pre-existing AWS Config role*
	// >
	// > If you have used an AWS service that uses AWS Config , such as Security Hub or AWS Control Tower , and an AWS Config role has already been created, make sure that the IAM role that you use when setting up AWS Config keeps the same minimum permissions as the already created AWS Config role. You must do this so that the other AWS service continues to run as expected.
	// >
	// > For example, if AWS Control Tower has an IAM role that allows AWS Config to read Amazon Simple Storage Service ( Amazon S3 ) objects, make sure that the same permissions are granted within the IAM role you use when setting up AWS Config . Otherwise, it may interfere with how AWS Control Tower operates. For more information about IAM roles for AWS Config , see [*Identity and Access Management for AWS Config*](https://docs.aws.amazon.com/config/latest/developerguide/security-iam.html) in the *AWS Config Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html#cfn-config-configurationrecorder-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

