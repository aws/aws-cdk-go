package awsecs


// The `Ulimit` property specifies the `ulimit` settings to pass to the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ulimitProperty := &ulimitProperty{
//   	hardLimit: jsii.Number(123),
//   	name: jsii.String("name"),
//   	softLimit: jsii.Number(123),
//   }
//
type CfnTaskDefinition_UlimitProperty struct {
	// The hard limit for the ulimit type.
	HardLimit *float64 `field:"required" json:"hardLimit" yaml:"hardLimit"`
	// The `type` of the `ulimit` .
	Name *string `field:"required" json:"name" yaml:"name"`
	// The soft limit for the ulimit type.
	SoftLimit *float64 `field:"required" json:"softLimit" yaml:"softLimit"`
}

