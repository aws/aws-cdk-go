package previewawsemrserverlessmixins


// The scheduler configuration for batch and streaming jobs running on this application.
//
// Supported with release labels emr-7.0.0 and above.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   schedulerConfigurationProperty := &SchedulerConfigurationProperty{
//   	MaxConcurrentRuns: jsii.Number(123),
//   	QueueTimeoutMinutes: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-schedulerconfiguration.html
//
type CfnApplicationPropsMixin_SchedulerConfigurationProperty struct {
	// The maximum concurrent job runs on this application.
	//
	// If scheduler configuration is enabled on your application, the default value is 15. The valid range is 1 to 1000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-schedulerconfiguration.html#cfn-emrserverless-application-schedulerconfiguration-maxconcurrentruns
	//
	MaxConcurrentRuns *float64 `field:"optional" json:"maxConcurrentRuns" yaml:"maxConcurrentRuns"`
	// The maximum duration in minutes for the job in QUEUED state.
	//
	// If scheduler configuration is enabled on your application, the default value is 360 minutes (6 hours). The valid range is from 15 to 720.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-schedulerconfiguration.html#cfn-emrserverless-application-schedulerconfiguration-queuetimeoutminutes
	//
	QueueTimeoutMinutes *float64 `field:"optional" json:"queueTimeoutMinutes" yaml:"queueTimeoutMinutes"`
}

