package awslogs


// Copy Value processor, copies values from source to target for each entry.
//
// This processor copies values within a log event.
// You can also use this processor to add metadata to log events by copying values from metadata keys.
// For more information about this processor including examples, see copyValue in the CloudWatch Logs User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   copyValueProperty := &CopyValueProperty{
//   	Entries: []CopyValueEntryProperty{
//   		&CopyValueEntryProperty{
//   			Source: jsii.String("source"),
//   			Target: jsii.String("target"),
//
//   			// the properties below are optional
//   			OverwriteIfExists: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CopyValueProperty struct {
	// List of sources and target to copy.
	//
	// An array of CopyValueEntry objects, where each object contains information about one field value to copy.
	Entries *[]*CopyValueEntryProperty `field:"required" json:"entries" yaml:"entries"`
}

