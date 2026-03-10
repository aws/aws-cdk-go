package previewawscloudformationevents


// Type definition for Result.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   result := &Result{
//   	Data: []*string{
//   		jsii.String("data"),
//   	},
//   }
//
// Experimental.
type CloudFormationHookInvocationProgress_Result struct {
	// data property.
	//
	// Specify an array of string values to match this event if the actual value of data is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Data *[]*string `field:"optional" json:"data" yaml:"data"`
}

