package awseventsv2


// Delivery logging configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   logConfigurationProperty := &LogConfigurationProperty{
//   	IncludePayload: jsii.String("includePayload"),
//   	Level: jsii.String("level"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-logconfiguration.html
//
type CfnSubscriberPropsMixin_LogConfigurationProperty struct {
	// Whether the event payload is included in emitted log records: FULL includes it in every emitted record, and ON_ERROR_ONLY includes it only in error records.
	//
	// The default is ON_ERROR_ONLY.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-logconfiguration.html#cfn-eventsv2-subscriber-logconfiguration-includepayload
	//
	IncludePayload *string `field:"optional" json:"includePayload" yaml:"includePayload"`
	// The minimum log level: OFF (no logging), ERROR, or INFO.
	//
	// Records below this level are not emitted. The default is OFF.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-logconfiguration.html#cfn-eventsv2-subscriber-logconfiguration-level
	//
	Level *string `field:"optional" json:"level" yaml:"level"`
}

