package awsmediaconvert


// Properties for defining a `CfnJobTemplate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var settingsJson interface{}
//   var tags interface{}
//
//   cfnJobTemplateProps := &cfnJobTemplateProps{
//   	settingsJson: settingsJson,
//
//   	// the properties below are optional
//   	accelerationSettings: &accelerationSettingsProperty{
//   		mode: jsii.String("mode"),
//   	},
//   	category: jsii.String("category"),
//   	description: jsii.String("description"),
//   	hopDestinations: []interface{}{
//   		&hopDestinationProperty{
//   			priority: jsii.Number(123),
//   			queue: jsii.String("queue"),
//   			waitMinutes: jsii.Number(123),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	priority: jsii.Number(123),
//   	queue: jsii.String("queue"),
//   	statusUpdateInterval: jsii.String("statusUpdateInterval"),
//   	tags: tags,
//   }
//
type CfnJobTemplateProps struct {
	// Specify, in JSON format, the transcoding job settings for this job template.
	//
	// This specification must conform to the AWS Elemental MediaConvert job validation. For information about forming this specification, see the Remarks section later in this topic.
	//
	// For more information about MediaConvert job templates, see [Working with AWS Elemental MediaConvert Job Templates](https://docs.aws.amazon.com/mediaconvert/latest/ug/working-with-job-templates.html) in the ** .
	SettingsJson interface{} `field:"required" json:"settingsJson" yaml:"settingsJson"`
	// Accelerated transcoding can significantly speed up jobs with long, visually complex content.
	//
	// Outputs that use this feature incur pro-tier pricing. For information about feature limitations, For more information, see [Job Limitations for Accelerated Transcoding in AWS Elemental MediaConvert](https://docs.aws.amazon.com/mediaconvert/latest/ug/job-requirements.html) in the *AWS Elemental MediaConvert User Guide* .
	AccelerationSettings interface{} `field:"optional" json:"accelerationSettings" yaml:"accelerationSettings"`
	// Optional.
	//
	// A category for the job template you are creating.
	Category *string `field:"optional" json:"category" yaml:"category"`
	// Optional.
	//
	// A description of the job template you are creating.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional.
	//
	// Configuration for a destination queue to which the job can hop once a customer-defined minimum wait time has passed. For more information, see [Setting Up Queue Hopping to Avoid Long Waits](https://docs.aws.amazon.com/mediaconvert/latest/ug/setting-up-queue-hopping-to-avoid-long-waits.html) in the *AWS Elemental MediaConvert User Guide* .
	HopDestinations interface{} `field:"optional" json:"hopDestinations" yaml:"hopDestinations"`
	// The name of the job template you are creating.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify the relative priority for this job.
	//
	// In any given queue, the service begins processing the job with the highest value first. When more than one job has the same priority, the service begins processing the job that you submitted first. If you don't specify a priority, the service uses the default value 0. Minimum: -50 Maximum: 50
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Optional.
	//
	// The queue that jobs created from this template are assigned to. Specify the Amazon Resource Name (ARN) of the queue. For example, arn:aws:mediaconvert:us-west-2:505474453218:queues/Default. If you don't specify this, jobs will go to the default queue.
	Queue *string `field:"optional" json:"queue" yaml:"queue"`
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
	StatusUpdateInterval *string `field:"optional" json:"statusUpdateInterval" yaml:"statusUpdateInterval"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

