package previewawsautoscalingevents


// Type definition for ResponseElements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var alarms interface{}
//   var failedScheduledActions interface{}
//   var failedScheduledUpdateGroupActions interface{}
//
//   responseElements := &ResponseElements{
//   	Alarms: []interface{}{
//   		alarms,
//   	},
//   	FailedScheduledActions: []interface{}{
//   		failedScheduledActions,
//   	},
//   	FailedScheduledUpdateGroupActions: []interface{}{
//   		failedScheduledUpdateGroupActions,
//   	},
//   	PolicyArn: []*string{
//   		jsii.String("policyArn"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_ResponseElements struct {
	// alarms property.
	//
	// Specify an array of string values to match this event if the actual value of alarms is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Alarms *[]interface{} `field:"optional" json:"alarms" yaml:"alarms"`
	// failedScheduledActions property.
	//
	// Specify an array of string values to match this event if the actual value of failedScheduledActions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FailedScheduledActions *[]interface{} `field:"optional" json:"failedScheduledActions" yaml:"failedScheduledActions"`
	// failedScheduledUpdateGroupActions property.
	//
	// Specify an array of string values to match this event if the actual value of failedScheduledUpdateGroupActions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FailedScheduledUpdateGroupActions *[]interface{} `field:"optional" json:"failedScheduledUpdateGroupActions" yaml:"failedScheduledUpdateGroupActions"`
	// policyARN property.
	//
	// Specify an array of string values to match this event if the actual value of policyARN is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PolicyArn *[]*string `field:"optional" json:"policyArn" yaml:"policyArn"`
}

