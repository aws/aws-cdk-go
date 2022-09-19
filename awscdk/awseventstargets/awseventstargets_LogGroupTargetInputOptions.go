package awseventstargets


// Options used when creating a target input template.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var message interface{}
//   var timestamp interface{}
//
//   logGroupTargetInputOptions := &logGroupTargetInputOptions{
//   	message: message,
//   	timestamp: timestamp,
//   }
//
type LogGroupTargetInputOptions struct {
	// The value provided here will be used in the Log "message" field.
	//
	// This field must be a string. If an object is passed (e.g. JSON data)
	// it will not throw an error, but the message that makes it to
	// CloudWatch logs will be incorrect. This is a likely scenario if
	// doing something like: EventField.fromPath('$.detail') since in most cases
	// the `detail` field contains JSON data.
	Message interface{} `field:"optional" json:"message" yaml:"message"`
	// The timestamp that will appear in the CloudWatch Logs record.
	Timestamp interface{} `field:"optional" json:"timestamp" yaml:"timestamp"`
}

