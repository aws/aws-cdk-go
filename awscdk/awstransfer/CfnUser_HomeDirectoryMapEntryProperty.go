package awstransfer


// Represents an object that contains entries and targets for `HomeDirectoryMappings` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   homeDirectoryMapEntryProperty := &HomeDirectoryMapEntryProperty{
//   	Entry: jsii.String("entry"),
//   	Target: jsii.String("target"),
//
//   	// the properties below are optional
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-user-homedirectorymapentry.html
//
type CfnUser_HomeDirectoryMapEntryProperty struct {
	// Represents an entry for `HomeDirectoryMappings` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-user-homedirectorymapentry.html#cfn-transfer-user-homedirectorymapentry-entry
	//
	Entry *string `field:"required" json:"entry" yaml:"entry"`
	// Represents the map target that is used in a `HomeDirectoryMapEntry` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-user-homedirectorymapentry.html#cfn-transfer-user-homedirectorymapentry-target
	//
	Target *string `field:"required" json:"target" yaml:"target"`
	// Specifies the type of mapping.
	//
	// Set the type to `FILE` if you want the mapping to point to a file, or `DIRECTORY` for the directory to point to a directory.
	//
	// > By default, home directory mappings have a `Type` of `DIRECTORY` when you create a Transfer Family server. You would need to explicitly set `Type` to `FILE` if you want a mapping to have a file target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-user-homedirectorymapentry.html#cfn-transfer-user-homedirectorymapentry-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

