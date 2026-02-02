package previewawsmwaaserverlessmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scheduleConfigurationProperty := &ScheduleConfigurationProperty{
//   	CronExpression: jsii.String("cronExpression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaaserverless-workflow-scheduleconfiguration.html
//
type CfnWorkflowPropsMixin_ScheduleConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mwaaserverless-workflow-scheduleconfiguration.html#cfn-mwaaserverless-workflow-scheduleconfiguration-cronexpression
	//
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
}

