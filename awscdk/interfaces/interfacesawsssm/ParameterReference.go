package interfacesawsssm


// A reference to a Parameter resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterReference := &ParameterReference{
//   	ParameterName: jsii.String("parameterName"),
//   }
//
type ParameterReference struct {
	// The Name of the Parameter resource.
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
}

