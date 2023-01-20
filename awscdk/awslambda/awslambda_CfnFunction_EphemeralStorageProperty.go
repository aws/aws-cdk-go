package awslambda


// The size of the function’s /tmp directory in MB.
//
// The default value is 512, but can be any whole number between 512 and 10240 MB.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ephemeralStorageProperty := &ephemeralStorageProperty{
//   	size: jsii.Number(123),
//   }
//
type CfnFunction_EphemeralStorageProperty struct {
	// The size of the function’s /tmp directory.
	Size *float64 `field:"required" json:"size" yaml:"size"`
}

