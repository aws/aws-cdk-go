package previewawsec2events


// Type definition for StateReason.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   stateReason := &StateReason{
//   	Code: []*string{
//   		jsii.String("code"),
//   	},
//   	Message: []*string{
//   		jsii.String("message"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_StateReason struct {
	// code property.
	//
	// Specify an array of string values to match this event if the actual value of code is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Code *[]*string `field:"optional" json:"code" yaml:"code"`
	// message property.
	//
	// Specify an array of string values to match this event if the actual value of message is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Message *[]*string `field:"optional" json:"message" yaml:"message"`
}

