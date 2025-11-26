package previewawsglueevents


// Type definition for NotificationCondition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   notificationCondition := &NotificationCondition{
//   	NotifyDelayAfter: []*string{
//   		jsii.String("notifyDelayAfter"),
//   	},
//   }
//
// Experimental.
type JobEvents_GlueJobRunStatus_NotificationCondition struct {
	// NotifyDelayAfter property.
	//
	// Specify an array of string values to match this event if the actual value of NotifyDelayAfter is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NotifyDelayAfter *[]*string `field:"optional" json:"notifyDelayAfter" yaml:"notifyDelayAfter"`
}

