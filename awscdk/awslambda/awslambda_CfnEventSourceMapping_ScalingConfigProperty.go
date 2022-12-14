package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scalingConfigProperty := &scalingConfigProperty{
//   	maximumConcurrency: jsii.Number(123),
//   }
//
type CfnEventSourceMapping_ScalingConfigProperty struct {
	// `CfnEventSourceMapping.ScalingConfigProperty.MaximumConcurrency`.
	MaximumConcurrency *float64 `field:"optional" json:"maximumConcurrency" yaml:"maximumConcurrency"`
}

