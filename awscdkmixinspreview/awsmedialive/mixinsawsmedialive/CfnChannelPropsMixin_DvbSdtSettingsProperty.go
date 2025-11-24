package mixinsawsmedialive


// A DVB Service Description Table (SDT).
//
// The parent of this entity is M2tsSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dvbSdtSettingsProperty := &DvbSdtSettingsProperty{
//   	OutputSdt: jsii.String("outputSdt"),
//   	RepInterval: jsii.Number(123),
//   	ServiceName: jsii.String("serviceName"),
//   	ServiceProviderName: jsii.String("serviceProviderName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsdtsettings.html
//
type CfnChannelPropsMixin_DvbSdtSettingsProperty struct {
	// Selects a method of inserting SDT information into an output stream.
	//
	// The sdtFollow setting copies SDT information from input stream to output stream. The sdtFollowIfPresent setting copies SDT information from input stream to output stream if SDT information is present in the input. Otherwise, it falls back on the user-defined values. The sdtManual setting means that the user will enter the SDT information. The sdtNone setting means that the output stream will not contain SDT information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsdtsettings.html#cfn-medialive-channel-dvbsdtsettings-outputsdt
	//
	OutputSdt *string `field:"optional" json:"outputSdt" yaml:"outputSdt"`
	// The number of milliseconds between instances of this table in the output transport stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsdtsettings.html#cfn-medialive-channel-dvbsdtsettings-repinterval
	//
	RepInterval *float64 `field:"optional" json:"repInterval" yaml:"repInterval"`
	// The service name placed in the serviceDescriptor in the Service Description Table (SDT).
	//
	// The maximum length is 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsdtsettings.html#cfn-medialive-channel-dvbsdtsettings-servicename
	//
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The service provider name placed in the serviceDescriptor in the Service Description Table (SDT).
	//
	// The maximum length is 256 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-dvbsdtsettings.html#cfn-medialive-channel-dvbsdtsettings-serviceprovidername
	//
	ServiceProviderName *string `field:"optional" json:"serviceProviderName" yaml:"serviceProviderName"`
}

