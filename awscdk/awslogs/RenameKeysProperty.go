package awslogs


// Use this processor to rename keys in a log event.
//
// For more information about this processor including examples, see renameKeys in the CloudWatch Logs User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   renameKeysProperty := &RenameKeysProperty{
//   	Entries: []RenameKeyEntryProperty{
//   		&RenameKeyEntryProperty{
//   			Key: jsii.String("key"),
//   			RenameTo: jsii.String("renameTo"),
//
//   			// the properties below are optional
//   			OverwriteIfExists: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type RenameKeysProperty struct {
	// An array of RenameKeyEntry objects, where each object contains information about one key to rename.
	Entries *[]*RenameKeyEntryProperty `field:"required" json:"entries" yaml:"entries"`
}

