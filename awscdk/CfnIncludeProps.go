package awscdk


// Construction properties for {@link CfnInclude}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var template interface{}
//
//   cfnIncludeProps := &CfnIncludeProps{
//   	Template: template,
//   }
//
// Deprecated: use the CfnInclude class from the cloudformation-include module instead.
type CfnIncludeProps struct {
	// The CloudFormation template to include in the stack (as is).
	// Deprecated: use the CfnInclude class from the cloudformation-include module instead.
	Template *map[string]interface{} `field:"required" json:"template" yaml:"template"`
}

