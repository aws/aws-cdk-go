package interfacesawscontrolcatalog


// A reference to a Control resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   controlReference := &ControlReference{
//   	ControlArn: jsii.String("controlArn"),
//   	ControlId: jsii.String("controlId"),
//   }
//
type ControlReference struct {
	// The ARN of the Control resource.
	ControlArn *string `field:"required" json:"controlArn" yaml:"controlArn"`
	// The ControlId of the Control resource.
	ControlId *string `field:"required" json:"controlId" yaml:"controlId"`
}

