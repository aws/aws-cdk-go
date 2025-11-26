package previewawsec2events


// Type definition for InstanceState.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   instanceState := &InstanceState{
//   	Code: []*string{
//   		jsii.String("code"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   }
//
// Experimental.
type InstanceEvents_AWSAPICallViaCloudTrail_InstanceState struct {
	// code property.
	//
	// Specify an array of string values to match this event if the actual value of code is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Code *[]*string `field:"optional" json:"code" yaml:"code"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
}

