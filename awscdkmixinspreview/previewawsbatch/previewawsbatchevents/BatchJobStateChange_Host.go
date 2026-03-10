package previewawsbatchevents


// Type definition for Host.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   host := &Host{
//   	SourcePath: []*string{
//   		jsii.String("sourcePath"),
//   	},
//   }
//
// Experimental.
type BatchJobStateChange_Host struct {
	// sourcePath property.
	//
	// Specify an array of string values to match this event if the actual value of sourcePath is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourcePath *[]*string `field:"optional" json:"sourcePath" yaml:"sourcePath"`
}

