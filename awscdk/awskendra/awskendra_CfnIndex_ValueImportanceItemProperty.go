package awskendra


// Specifies a key-value pair that determines the search boost value that a document receives when the key is part of the metadata of a document.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   valueImportanceItemProperty := &valueImportanceItemProperty{
//   	key: jsii.String("key"),
//   	value: jsii.Number(123),
//   }
//
type CfnIndex_ValueImportanceItemProperty struct {
	// The document metadata value that receives the search boost.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The boost value that a document receives when the key is part of the metadata of a document.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

