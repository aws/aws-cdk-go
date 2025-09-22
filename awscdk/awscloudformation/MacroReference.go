package awscloudformation


// A reference to a Macro resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   macroReference := &MacroReference{
//   	MacroId: jsii.String("macroId"),
//   }
//
type MacroReference struct {
	// The Id of the Macro resource.
	MacroId *string `field:"required" json:"macroId" yaml:"macroId"`
}

