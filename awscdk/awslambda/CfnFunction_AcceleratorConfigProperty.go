package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acceleratorConfigProperty := &AcceleratorConfigProperty{
//   	AcceleratorMemorySize: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-acceleratorconfig.html
//
type CfnFunction_AcceleratorConfigProperty struct {
	// The minimum amount of memory (in GB) of each accelerator device exposed to the execution environment.
	//
	// Lambda might allocate more, depending on the accelerator instance types available. Valid values: 3, 6, 12, 16, 24, 48.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-acceleratorconfig.html#cfn-lambda-function-acceleratorconfig-acceleratormemorysize
	//
	AcceleratorMemorySize *float64 `field:"required" json:"acceleratorMemorySize" yaml:"acceleratorMemorySize"`
}

