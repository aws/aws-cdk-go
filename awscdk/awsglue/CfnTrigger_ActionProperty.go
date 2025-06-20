package awsglue


// Defines an action to be initiated by a trigger.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var arguments_ interface{}
//
//   actionProperty := &ActionProperty{
//   	Arguments: arguments_,
//   	CrawlerName: jsii.String("crawlerName"),
//   	JobName: jsii.String("jobName"),
//   	NotificationProperty: &NotificationPropertyProperty{
//   		NotifyDelayAfter: jsii.Number(123),
//   	},
//   	SecurityConfiguration: jsii.String("securityConfiguration"),
//   	Timeout: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html
//
type CfnTrigger_ActionProperty struct {
	// The job arguments used when this trigger fires.
	//
	// For this job run, they replace the default arguments set in the job definition itself.
	//
	// You can specify arguments here that your own job-execution script consumes, in addition to arguments that AWS Glue itself consumes.
	//
	// For information about how to specify and consume your own job arguments, see [Calling AWS Glue APIs in Python](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) in the *AWS Glue Developer Guide* .
	//
	// For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html) topic in the developer guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html#cfn-glue-trigger-action-arguments
	//
	Arguments interface{} `field:"optional" json:"arguments" yaml:"arguments"`
	// The name of the crawler to be used with this action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html#cfn-glue-trigger-action-crawlername
	//
	CrawlerName *string `field:"optional" json:"crawlerName" yaml:"crawlerName"`
	// The name of a job to be executed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html#cfn-glue-trigger-action-jobname
	//
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// Specifies configuration properties of a job run notification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html#cfn-glue-trigger-action-notificationproperty
	//
	NotificationProperty interface{} `field:"optional" json:"notificationProperty" yaml:"notificationProperty"`
	// The name of the `SecurityConfiguration` structure to be used with this action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html#cfn-glue-trigger-action-securityconfiguration
	//
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The `JobRun` timeout in minutes.
	//
	// This is the maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status. The default is 2,880 minutes (48 hours). This overrides the timeout value set in the parent job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-trigger-action.html#cfn-glue-trigger-action-timeout
	//
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

