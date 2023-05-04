package awsquicksight


// The custom narrative options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customNarrativeOptionsProperty := &CustomNarrativeOptionsProperty{
//   	Narrative: jsii.String("narrative"),
//   }
//
type CfnTemplate_CustomNarrativeOptionsProperty struct {
	// The string input of custom narrative.
	Narrative *string `field:"required" json:"narrative" yaml:"narrative"`
}

