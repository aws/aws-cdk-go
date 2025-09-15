package awslogs


// This object defines one key that will be renamed with the renameKey processor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   renameKeyEntryProperty := &RenameKeyEntryProperty{
//   	Key: jsii.String("key"),
//   	RenameTo: jsii.String("renameTo"),
//
//   	// the properties below are optional
//   	OverwriteIfExists: jsii.Boolean(false),
//   }
//
type RenameKeyEntryProperty struct {
	// The key to rename.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The string to use for the new key name.
	RenameTo *string `field:"required" json:"renameTo" yaml:"renameTo"`
	// Whether to overwrite the target key if it already exists.
	// Default: false.
	//
	OverwriteIfExists *bool `field:"optional" json:"overwriteIfExists" yaml:"overwriteIfExists"`
}

