package awstransfer


// Represents an object that contains entries and targets for `HomeDirectoryMappings` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   homeDirectoryMapEntryProperty := &homeDirectoryMapEntryProperty{
//   	entry: jsii.String("entry"),
//   	target: jsii.String("target"),
//   }
//
type CfnUser_HomeDirectoryMapEntryProperty struct {
	// Represents an entry for `HomeDirectoryMappings` .
	Entry *string `field:"required" json:"entry" yaml:"entry"`
	// Represents the map target that is used in a `HomeDirectorymapEntry` .
	Target *string `field:"required" json:"target" yaml:"target"`
}

