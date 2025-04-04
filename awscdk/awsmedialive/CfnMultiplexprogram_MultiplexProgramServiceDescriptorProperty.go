package awsmedialive


// Transport stream service descriptor configuration for the Multiplex program.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiplexProgramServiceDescriptorProperty := &MultiplexProgramServiceDescriptorProperty{
//   	ProviderName: jsii.String("providerName"),
//   	ServiceName: jsii.String("serviceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogramservicedescriptor.html
//
type CfnMultiplexprogram_MultiplexProgramServiceDescriptorProperty struct {
	// Name of the provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogramservicedescriptor.html#cfn-medialive-multiplexprogram-multiplexprogramservicedescriptor-providername
	//
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// Name of the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogramservicedescriptor.html#cfn-medialive-multiplexprogram-multiplexprogramservicedescriptor-servicename
	//
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
}

