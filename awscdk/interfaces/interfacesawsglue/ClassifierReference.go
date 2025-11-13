package interfacesawsglue


// A reference to a Classifier resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   classifierReference := &ClassifierReference{
//   	ClassifierId: jsii.String("classifierId"),
//   }
//
type ClassifierReference struct {
	// The Id of the Classifier resource.
	ClassifierId *string `field:"required" json:"classifierId" yaml:"classifierId"`
}

