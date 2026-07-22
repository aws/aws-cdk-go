package awscdk


// Information about a single stack that is being validated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   policyValidationStack := &PolicyValidationStack{
//   	StackConstructPath: jsii.String("stackConstructPath"),
//   	TemplatePath: jsii.String("templatePath"),
//   }
//
type PolicyValidationStack struct {
	// The Stack's construct path.
	StackConstructPath *string `field:"required" json:"stackConstructPath" yaml:"stackConstructPath"`
	// The path to the template file on disk.
	TemplatePath *string `field:"required" json:"templatePath" yaml:"templatePath"`
}

