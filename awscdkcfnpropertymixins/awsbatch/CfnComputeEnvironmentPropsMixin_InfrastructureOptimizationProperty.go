package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   infrastructureOptimizationProperty := &InfrastructureOptimizationProperty{
//   	ScaleInAfter: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-infrastructureoptimization.html
//
type CfnComputeEnvironmentPropsMixin_InfrastructureOptimizationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-infrastructureoptimization.html#cfn-batch-computeenvironment-infrastructureoptimization-scaleinafter
	//
	ScaleInAfter *float64 `field:"optional" json:"scaleInAfter" yaml:"scaleInAfter"`
}

