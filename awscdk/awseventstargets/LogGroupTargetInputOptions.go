package awseventstargets


// Options used when creating a target input template.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//   var logGroup logGroup
//   var rule rule
//
//
//   rule.AddTarget(targets.NewCloudWatchLogGroup(logGroup, &LogGroupProps{
//   	LogEvent: targets.LogGroupTargetInput_FromObject(&LogGroupTargetInputOptions{
//   		Timestamp: events.EventField_FromPath(jsii.String("$.time")),
//   		Message: events.EventField_*FromPath(jsii.String("$.detail-type")),
//   	}),
//   }))
//
type LogGroupTargetInputOptions struct {
	// The value provided here will be used in the Log "message" field.
	//
	// This field must be a string. If an object is passed (e.g. JSON data)
	// it will not throw an error, but the message that makes it to
	// CloudWatch logs will be incorrect. This is a likely scenario if
	// doing something like: EventField.fromPath('$.detail') since in most cases
	// the `detail` field contains JSON data.
	// Default: EventField.detailType
	//
	Message interface{} `field:"optional" json:"message" yaml:"message"`
	// The timestamp that will appear in the CloudWatch Logs record.
	// Default: EventField.time
	//
	Timestamp interface{} `field:"optional" json:"timestamp" yaml:"timestamp"`
}

