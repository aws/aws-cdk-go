package mixinsawsmediaconvert


// Properties for CfnJobTemplatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var settingsJson interface{}
//   var tags interface{}
//
//   cfnJobTemplateMixinProps := &CfnJobTemplateMixinProps{
//   	AccelerationSettings: &AccelerationSettingsProperty{
//   		Mode: jsii.String("mode"),
//   	},
//   	Category: jsii.String("category"),
//   	Description: jsii.String("description"),
//   	HopDestinations: []interface{}{
//   		&HopDestinationProperty{
//   			Priority: jsii.Number(123),
//   			Queue: jsii.String("queue"),
//   			WaitMinutes: jsii.Number(123),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Priority: jsii.Number(123),
//   	Queue: jsii.String("queue"),
//   	SettingsJson: settingsJson,
//   	StatusUpdateInterval: jsii.String("statusUpdateInterval"),
//   	Tags: tags,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html
//
type CfnJobTemplateMixinProps struct {
	// Accelerated transcoding can significantly speed up jobs with long, visually complex content.
	//
	// Outputs that use this feature incur pro-tier pricing. For information about feature limitations, For more information, see [Job Limitations for Accelerated Transcoding in AWS Elemental MediaConvert](https://docs.aws.amazon.com/mediaconvert/latest/ug/job-requirements.html) in the *AWS Elemental MediaConvert User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-accelerationsettings
	//
	AccelerationSettings interface{} `field:"optional" json:"accelerationSettings" yaml:"accelerationSettings"`
	// Optional.
	//
	// A category for the job template you are creating.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-category
	//
	Category *string `field:"optional" json:"category" yaml:"category"`
	// Optional.
	//
	// A description of the job template you are creating.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional.
	//
	// Configuration for a destination queue to which the job can hop once a customer-defined minimum wait time has passed. For more information, see [Setting Up Queue Hopping to Avoid Long Waits](https://docs.aws.amazon.com/mediaconvert/latest/ug/setting-up-queue-hopping-to-avoid-long-waits.html) in the *AWS Elemental MediaConvert User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-hopdestinations
	//
	HopDestinations interface{} `field:"optional" json:"hopDestinations" yaml:"hopDestinations"`
	// Name of the output group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify the relative priority for this job.
	//
	// In any given queue, the service begins processing the job with the highest value first. When more than one job has the same priority, the service begins processing the job that you submitted first. If you don't specify a priority, the service uses the default value 0. Minimum: -50 Maximum: 50
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-priority
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Optional.
	//
	// The queue that jobs created from this template are assigned to. Specify the Amazon Resource Name (ARN) of the queue. For example, arn:aws:mediaconvert:us-west-2:505474453218:queues/Default. If you don't specify this, jobs will go to the default queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-queue
	//
	Queue *string `field:"optional" json:"queue" yaml:"queue"`
	// Specify, in JSON format, the transcoding job settings for this job template.
	//
	// This specification must conform to the AWS Elemental MediaConvert job validation. For information about forming this specification, see the Remarks section later in this topic.
	//
	// For more information about MediaConvert job templates, see [Working with AWS Elemental MediaConvert Job Templates](https://docs.aws.amazon.com/mediaconvert/latest/ug/working-with-job-templates.html) in the ** .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-settingsjson
	//
	SettingsJson interface{} `field:"optional" json:"settingsJson" yaml:"settingsJson"`
	// Specify how often MediaConvert sends STATUS_UPDATE events to Amazon CloudWatch Events.
	//
	// Set the interval, in seconds, between status updates. MediaConvert sends an update at this interval from the time the service begins processing your job to the time it completes the transcode or encounters an error.
	//
	// Specify one of the following enums:
	//
	// SECONDS_10
	//
	// SECONDS_12
	//
	// SECONDS_15
	//
	// SECONDS_20
	//
	// SECONDS_30
	//
	// SECONDS_60
	//
	// SECONDS_120
	//
	// SECONDS_180
	//
	// SECONDS_240
	//
	// SECONDS_300
	//
	// SECONDS_360
	//
	// SECONDS_420
	//
	// SECONDS_480
	//
	// SECONDS_540
	//
	// SECONDS_600.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-statusupdateinterval
	//
	StatusUpdateInterval *string `field:"optional" json:"statusUpdateInterval" yaml:"statusUpdateInterval"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html#cfn-mediaconvert-jobtemplate-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

