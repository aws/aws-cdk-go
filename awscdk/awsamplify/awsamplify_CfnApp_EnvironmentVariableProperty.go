package awsamplify


// Environment variables are key-value pairs that are available at build time.
//
// Set environment variables for all branches in your app.
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
type CfnApp_EnvironmentVariableProperty struct {
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

