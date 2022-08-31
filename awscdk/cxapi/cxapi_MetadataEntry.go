package cxapi


// Backwards compatibility for when `MetadataEntry` was defined here.
//
// This is necessary because its used as an input in the stable.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataEntry := &metadataEntry{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	data: jsii.String("data"),
//   	trace: []*string{
//   		jsii.String("trace"),
//   	},
//   }
//
// See: core.ConstructNode.metadata
//
// Deprecated: moved to package 'cloud-assembly-schema'.
type MetadataEntry struct {
	// The type of the metadata entry.
	// Deprecated: moved to package 'cloud-assembly-schema'.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The data.
	// Deprecated: moved to package 'cloud-assembly-schema'.
	Data interface{} `field:"optional" json:"data" yaml:"data"`
	// A stack trace for when the entry was created.
	// Deprecated: moved to package 'cloud-assembly-schema'.
	Trace *[]*string `field:"optional" json:"trace" yaml:"trace"`
}

