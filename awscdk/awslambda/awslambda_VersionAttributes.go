package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ function
//
//   versionAttributes := &VersionAttributes{
//   	Lambda: function_,
//   	Version: jsii.String("version"),
//   }
//
// Experimental.
type VersionAttributes struct {
	// The lambda function.
	// Experimental.
	Lambda IFunction `field:"required" json:"lambda" yaml:"lambda"`
	// The version.
	// Experimental.
	Version *string `field:"required" json:"version" yaml:"version"`
}

