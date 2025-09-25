package awsbatch


// Sets limits for a resource with `ulimit` on linux systems.
//
// Used by the Docker daemon.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ulimit := &Ulimit{
//   	HardLimit: jsii.Number(123),
//   	Name: awscdk.Aws_batch.UlimitName_CORE,
//   	SoftLimit: jsii.Number(123),
//   }
//
type Ulimit struct {
	// The hard limit for this resource.
	//
	// The container will
	// be terminated if it exceeds this limit.
	HardLimit *float64 `field:"required" json:"hardLimit" yaml:"hardLimit"`
	// The resource to limit.
	Name UlimitName `field:"required" json:"name" yaml:"name"`
	// The reservation for this resource.
	//
	// The container will
	// not be terminated if it exceeds this limit.
	SoftLimit *float64 `field:"required" json:"softLimit" yaml:"softLimit"`
}

