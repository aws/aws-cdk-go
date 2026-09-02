package interfacesawsglue


// A reference to a Classifier resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   classifierReference := &ClassifierReference{
//   	ClassifierName: jsii.String("classifierName"),
//   }
//
type ClassifierReference struct {
	// The Name of the Classifier resource.
	ClassifierName *string `field:"required" json:"classifierName" yaml:"classifierName"`
}

