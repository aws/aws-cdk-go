package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cpuConfigurationProperty := &CpuConfigurationProperty{
//   	Architecture: jsii.String("architecture"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-cpuconfiguration.html
//
type CfnMicrovmImage_CpuConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-cpuconfiguration.html#cfn-lambda-microvmimage-cpuconfiguration-architecture
	//
	Architecture *string `field:"required" json:"architecture" yaml:"architecture"`
}

