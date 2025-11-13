package interfacesawsfis


// A reference to a ExperimentTemplate resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   experimentTemplateReference := &ExperimentTemplateReference{
//   	ExperimentTemplateId: jsii.String("experimentTemplateId"),
//   }
//
type ExperimentTemplateReference struct {
	// The Id of the ExperimentTemplate resource.
	ExperimentTemplateId *string `field:"required" json:"experimentTemplateId" yaml:"experimentTemplateId"`
}

