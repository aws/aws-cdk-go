package interfacesawsconnect


// A reference to a PredefinedAttribute resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   predefinedAttributeReference := &PredefinedAttributeReference{
//   	InstanceArn: jsii.String("instanceArn"),
//   	PredefinedAttributeName: jsii.String("predefinedAttributeName"),
//   }
//
type PredefinedAttributeReference struct {
	// The InstanceArn of the PredefinedAttribute resource.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The Name of the PredefinedAttribute resource.
	PredefinedAttributeName *string `field:"required" json:"predefinedAttributeName" yaml:"predefinedAttributeName"`
}

