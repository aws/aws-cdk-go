package previewawsecsmixins


// The minimum and maximum amount of memory in mebibytes (MiB) for instance type selection.
//
// This ensures that selected instance types have adequate memory for your workloads.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   memoryMiBRequestProperty := &MemoryMiBRequestProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-memorymibrequest.html
//
type CfnCapacityProviderPropsMixin_MemoryMiBRequestProperty struct {
	// The maximum amount of memory in MiB.
	//
	// Instance types with more memory than this value are excluded from selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-memorymibrequest.html#cfn-ecs-capacityprovider-memorymibrequest-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum amount of memory in MiB.
	//
	// Instance types with less memory than this value are excluded from selection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-capacityprovider-memorymibrequest.html#cfn-ecs-capacityprovider-memorymibrequest-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

