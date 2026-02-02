package previewawsemrcontainersmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   containerLogRotationConfigurationProperty := &ContainerLogRotationConfigurationProperty{
//   	MaxFilesToKeep: jsii.Number(123),
//   	RotationSize: jsii.String("rotationSize"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-containerlogrotationconfiguration.html
//
type CfnEndpointPropsMixin_ContainerLogRotationConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-containerlogrotationconfiguration.html#cfn-emrcontainers-endpoint-containerlogrotationconfiguration-maxfilestokeep
	//
	MaxFilesToKeep *float64 `field:"optional" json:"maxFilesToKeep" yaml:"maxFilesToKeep"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-endpoint-containerlogrotationconfiguration.html#cfn-emrcontainers-endpoint-containerlogrotationconfiguration-rotationsize
	//
	RotationSize *string `field:"optional" json:"rotationSize" yaml:"rotationSize"`
}

