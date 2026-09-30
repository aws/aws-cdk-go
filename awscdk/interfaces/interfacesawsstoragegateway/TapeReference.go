package interfacesawsstoragegateway


// A reference to a Tape resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tapeReference := &TapeReference{
//   	TapeArn: jsii.String("tapeArn"),
//   }
//
type TapeReference struct {
	// The TapeARN of the Tape resource.
	TapeArn *string `field:"required" json:"tapeArn" yaml:"tapeArn"`
}

