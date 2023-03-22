package awsmedialive


// A DVB Service Description Table (SDT).
//
// The parent of this entity is M2tsSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dvbSdtSettingsProperty := &dvbSdtSettingsProperty{
//   	outputSdt: jsii.String("outputSdt"),
//   	repInterval: jsii.Number(123),
//   	serviceName: jsii.String("serviceName"),
//   	serviceProviderName: jsii.String("serviceProviderName"),
//   }
//
type CfnChannel_DvbSdtSettingsProperty struct {
	// Selects a method of inserting SDT information into an output stream.
	//
	// The sdtFollow setting copies SDT information from input stream to output stream. The sdtFollowIfPresent setting copies SDT information from input stream to output stream if SDT information is present in the input. Otherwise, it falls back on the user-defined values. The sdtManual setting means that the user will enter the SDT information. The sdtNone setting means that the output stream will not contain SDT information.
	OutputSdt *string `field:"optional" json:"outputSdt" yaml:"outputSdt"`
	// The number of milliseconds between instances of this table in the output transport stream.
	RepInterval *float64 `field:"optional" json:"repInterval" yaml:"repInterval"`
	// The service name placed in the serviceDescriptor in the Service Description Table (SDT).
	//
	// The maximum length is 256 characters.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// The service provider name placed in the serviceDescriptor in the Service Description Table (SDT).
	//
	// The maximum length is 256 characters.
	ServiceProviderName *string `field:"optional" json:"serviceProviderName" yaml:"serviceProviderName"`
}

