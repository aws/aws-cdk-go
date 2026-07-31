package awsiam


// Options for resource-based policy validation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourcePolicyValidationOptions := &ResourcePolicyValidationOptions{
//   	SkipResourceValidation: jsii.Boolean(false),
//   }
//
type ResourcePolicyValidationOptions struct {
	// Whether to skip resource validation for policies where resources are implicit (e.g., ECR repository policies where the resource is the repository itself).
	// Default: false.
	//
	SkipResourceValidation *bool `field:"optional" json:"skipResourceValidation" yaml:"skipResourceValidation"`
}

