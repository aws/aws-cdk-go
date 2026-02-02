package awsemrcontainers


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerLogRotationConfigurationProperty := &ContainerLogRotationConfigurationProperty{
//   	MaxFilesToKeep: jsii.Number(123),
//   	RotationSize: jsii.String("rotationSize"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-containerlogrotationconfiguration.html
//
type CfnEndpoint_ContainerLogRotationConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-containerlogrotationconfiguration.html#cfn-emrcontainers-endpoint-containerlogrotationconfiguration-maxfilestokeep
	//
	MaxFilesToKeep *float64 `field:"required" json:"maxFilesToKeep" yaml:"maxFilesToKeep"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-containerlogrotationconfiguration.html#cfn-emrcontainers-endpoint-containerlogrotationconfiguration-rotationsize
	//
	RotationSize *string `field:"required" json:"rotationSize" yaml:"rotationSize"`
}

