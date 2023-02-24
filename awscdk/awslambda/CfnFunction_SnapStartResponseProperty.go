package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snapStartResponseProperty := &SnapStartResponseProperty{
//   	ApplyOn: jsii.String("applyOn"),
//   	OptimizationStatus: jsii.String("optimizationStatus"),
//   }
//
type CfnFunction_SnapStartResponseProperty struct {
	// `CfnFunction.SnapStartResponseProperty.ApplyOn`.
	ApplyOn *string `field:"optional" json:"applyOn" yaml:"applyOn"`
	// `CfnFunction.SnapStartResponseProperty.OptimizationStatus`.
	OptimizationStatus *string `field:"optional" json:"optimizationStatus" yaml:"optimizationStatus"`
}

