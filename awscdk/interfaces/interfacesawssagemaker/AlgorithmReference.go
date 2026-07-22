package interfacesawssagemaker


// A reference to a Algorithm resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   algorithmReference := &AlgorithmReference{
//   	AlgorithmArn: jsii.String("algorithmArn"),
//   }
//
type AlgorithmReference struct {
	// The AlgorithmArn of the Algorithm resource.
	AlgorithmArn *string `field:"required" json:"algorithmArn" yaml:"algorithmArn"`
}

