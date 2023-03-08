package awsamplify


// The EnvironmentVariable property type sets environment variables for a specific branch.
//
// Environment variables are key-value pairs that are available at build time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentVariableProperty := &environmentVariableProperty{
//   	name: jsii.String("name"),
//   	value: jsii.String("value"),
//   }
//
type CfnBranch_EnvironmentVariableProperty struct {
	// The environment variable name.
	//
	// *Length Constraints:* Maximum length of 255.
	//
	// *Pattern:* (?s).*
	Name *string `field:"required" json:"name" yaml:"name"`
	// The environment variable value.
	//
	// *Length Constraints:* Maximum length of 5500.
	//
	// *Pattern:* (?s).*
	Value *string `field:"required" json:"value" yaml:"value"`
}

