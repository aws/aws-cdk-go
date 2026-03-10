package previewawsbatchevents


// Type definition for ULimit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   uLimit := &ULimit{
//   	HardLimit: []*string{
//   		jsii.String("hardLimit"),
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	SoftLimit: []*string{
//   		jsii.String("softLimit"),
//   	},
//   }
//
// Experimental.
type BatchJobStateChange_ULimit struct {
	// hardLimit property.
	//
	// Specify an array of string values to match this event if the actual value of hardLimit is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HardLimit *[]*string `field:"optional" json:"hardLimit" yaml:"hardLimit"`
	// name property.
	//
	// Specify an array of string values to match this event if the actual value of name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// softLimit property.
	//
	// Specify an array of string values to match this event if the actual value of softLimit is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SoftLimit *[]*string `field:"optional" json:"softLimit" yaml:"softLimit"`
}

