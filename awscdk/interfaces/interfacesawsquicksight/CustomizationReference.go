package interfacesawsquicksight


// A reference to a Customization resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customizationReference := &CustomizationReference{
//   	CustomizationArn: jsii.String("customizationArn"),
//   }
//
type CustomizationReference struct {
	// The Arn of the Customization resource.
	CustomizationArn *string `field:"required" json:"customizationArn" yaml:"customizationArn"`
}

