package interfacesawssagemaker


// A reference to a HumanTaskUi resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   humanTaskUiReference := &HumanTaskUiReference{
//   	HumanTaskUiArn: jsii.String("humanTaskUiArn"),
//   }
//
type HumanTaskUiReference struct {
	// The HumanTaskUiArn of the HumanTaskUi resource.
	HumanTaskUiArn *string `field:"required" json:"humanTaskUiArn" yaml:"humanTaskUiArn"`
}

