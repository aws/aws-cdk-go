package awslogs


// This processor moves a key from one field to another.
//
// The original key is deleted.
// For more information about this processor including examples, see moveKeys in the CloudWatch Logs User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   moveKeysProperty := &MoveKeysProperty{
//   	Entries: []moveKeyEntryProperty{
//   		&moveKeyEntryProperty{
//   			Source: jsii.String("source"),
//   			Target: jsii.String("target"),
//
//   			// the properties below are optional
//   			OverwriteIfExists: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type MoveKeysProperty struct {
	// An array of objects, where each object contains information about one key to move.
	Entries *[]*MoveKeyEntryProperty `field:"required" json:"entries" yaml:"entries"`
}

