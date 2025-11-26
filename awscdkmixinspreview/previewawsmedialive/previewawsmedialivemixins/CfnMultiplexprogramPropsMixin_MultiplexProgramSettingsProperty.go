package previewawsmedialivemixins


// Multiplex Program settings configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   multiplexProgramSettingsProperty := &MultiplexProgramSettingsProperty{
//   	PreferredChannelPipeline: jsii.String("preferredChannelPipeline"),
//   	ProgramNumber: jsii.Number(123),
//   	ServiceDescriptor: &MultiplexProgramServiceDescriptorProperty{
//   		ProviderName: jsii.String("providerName"),
//   		ServiceName: jsii.String("serviceName"),
//   	},
//   	VideoSettings: &MultiplexVideoSettingsProperty{
//   		ConstantBitrate: jsii.Number(123),
//   		StatmuxSettings: &MultiplexStatmuxVideoSettingsProperty{
//   			MaximumBitrate: jsii.Number(123),
//   			MinimumBitrate: jsii.Number(123),
//   			Priority: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogramsettings.html
//
type CfnMultiplexprogramPropsMixin_MultiplexProgramSettingsProperty struct {
	// Indicates which pipeline is preferred by the multiplex for program ingest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogramsettings.html#cfn-medialive-multiplexprogram-multiplexprogramsettings-preferredchannelpipeline
	//
	PreferredChannelPipeline *string `field:"optional" json:"preferredChannelPipeline" yaml:"preferredChannelPipeline"`
	// Unique program number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogramsettings.html#cfn-medialive-multiplexprogram-multiplexprogramsettings-programnumber
	//
	ProgramNumber *float64 `field:"optional" json:"programNumber" yaml:"programNumber"`
	// Transport stream service descriptor configuration for the Multiplex program.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogramsettings.html#cfn-medialive-multiplexprogram-multiplexprogramsettings-servicedescriptor
	//
	ServiceDescriptor interface{} `field:"optional" json:"serviceDescriptor" yaml:"serviceDescriptor"`
	// Program video settings configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogramsettings.html#cfn-medialive-multiplexprogram-multiplexprogramsettings-videosettings
	//
	VideoSettings interface{} `field:"optional" json:"videoSettings" yaml:"videoSettings"`
}

