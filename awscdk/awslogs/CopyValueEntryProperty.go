package awslogs


// This object defines one value to be copied with the copyValue processor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   copyValueEntryProperty := &CopyValueEntryProperty{
//   	Source: jsii.String("source"),
//   	Target: jsii.String("target"),
//
//   	// the properties below are optional
//   	OverwriteIfExists: jsii.Boolean(false),
//   }
//
type CopyValueEntryProperty struct {
	// The key to copy.
	Source *string `field:"required" json:"source" yaml:"source"`
	// The key of the field to copy the value to.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Specifies whether to overwrite the value if the target key already exists.
	// Default: false.
	//
	OverwriteIfExists *bool `field:"optional" json:"overwriteIfExists" yaml:"overwriteIfExists"`
}

