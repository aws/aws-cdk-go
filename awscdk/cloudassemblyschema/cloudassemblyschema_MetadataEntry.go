package cloudassemblyschema


// A metadata entry in a cloud assembly artifact.
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
type MetadataEntry struct {
	// The type of the metadata entry.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The data.
	Data interface{} `field:"optional" json:"data" yaml:"data"`
	// A stack trace for when the entry was created.
	Trace *[]*string `field:"optional" json:"trace" yaml:"trace"`
}

