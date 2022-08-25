package awsecs


// The ulimit settings to pass to the container.
//
// NOTE: Does not work for Windows containers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ulimit := &ulimit{
//   	hardLimit: jsii.Number(123),
//   	name: awscdk.Aws_ecs.ulimitName_CORE,
//   	softLimit: jsii.Number(123),
//   }
//
type Ulimit struct {
	// The hard limit for the ulimit type.
	HardLimit *float64 `field:"required" json:"hardLimit" yaml:"hardLimit"`
	// The type of the ulimit.
	//
	// For more information, see [UlimitName](https://docs.aws.amazon.com/cdk/api/latest/typescript/api/aws-ecs/ulimitname.html#aws_ecs_UlimitName).
	Name UlimitName `field:"required" json:"name" yaml:"name"`
	// The soft limit for the ulimit type.
	SoftLimit *float64 `field:"required" json:"softLimit" yaml:"softLimit"`
}

