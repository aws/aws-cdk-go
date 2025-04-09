package awsappsync


// Config for binding runtime to a function or resolver.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runtimeConfig := &RuntimeConfig{
//   	Name: jsii.String("name"),
//   	RuntimeVersion: jsii.String("runtimeVersion"),
//   }
//
type RuntimeConfig struct {
	// The name of the runtime.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The version string of the runtime.
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
}

