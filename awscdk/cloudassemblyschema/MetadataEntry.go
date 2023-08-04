package cloudassemblyschema


// A metadata entry in a cloud assembly artifact.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataEntry := &MetadataEntry{
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Data: jsii.String("data"),
//   	Trace: []*string{
//   		jsii.String("trace"),
//   	},
//   }
//
type MetadataEntry struct {
	// The type of the metadata entry.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The data.
	// Default: - no data.
	//
	Data interface{} `field:"optional" json:"data" yaml:"data"`
	// A stack trace for when the entry was created.
	// Default: - no trace.
	//
	Trace *[]*string `field:"optional" json:"trace" yaml:"trace"`
}

