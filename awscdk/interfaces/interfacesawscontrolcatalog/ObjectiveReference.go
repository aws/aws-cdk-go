package interfacesawscontrolcatalog


// A reference to a Objective resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   objectiveReference := &ObjectiveReference{
//   	ObjectiveArn: jsii.String("objectiveArn"),
//   }
//
type ObjectiveReference struct {
	// The Arn of the Objective resource.
	ObjectiveArn *string `field:"required" json:"objectiveArn" yaml:"objectiveArn"`
}

