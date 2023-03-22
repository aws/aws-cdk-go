package awskendra


// Specifies a key-value pair of the search boost value for a document when the key is part of the metadata of a document.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   valueImportanceItemProperty := &ValueImportanceItemProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.Number(123),
//   }
//
type CfnIndex_ValueImportanceItemProperty struct {
	// The document metadata value used for the search boost.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The boost value for a document when the key is part of the metadata of a document.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

