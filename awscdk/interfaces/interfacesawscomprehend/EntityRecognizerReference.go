package interfacesawscomprehend


// A reference to a EntityRecognizer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entityRecognizerReference := &EntityRecognizerReference{
//   	EntityRecognizerArn: jsii.String("entityRecognizerArn"),
//   }
//
type EntityRecognizerReference struct {
	// The Arn of the EntityRecognizer resource.
	EntityRecognizerArn *string `field:"required" json:"entityRecognizerArn" yaml:"entityRecognizerArn"`
}

