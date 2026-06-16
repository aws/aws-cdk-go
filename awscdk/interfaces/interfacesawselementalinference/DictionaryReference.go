package interfacesawselementalinference


// A reference to a Dictionary resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dictionaryReference := &DictionaryReference{
//   	DictionaryArn: jsii.String("dictionaryArn"),
//   	DictionaryId: jsii.String("dictionaryId"),
//   }
//
type DictionaryReference struct {
	// The ARN of the Dictionary resource.
	DictionaryArn *string `field:"required" json:"dictionaryArn" yaml:"dictionaryArn"`
	// The Id of the Dictionary resource.
	DictionaryId *string `field:"required" json:"dictionaryId" yaml:"dictionaryId"`
}

