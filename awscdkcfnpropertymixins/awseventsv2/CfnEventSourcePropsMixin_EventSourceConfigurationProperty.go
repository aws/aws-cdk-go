package awseventsv2


// Event source configuration.
//
// Exactly one of the two variants must be set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   eventSourceConfigurationProperty := &EventSourceConfigurationProperty{
//   	AwsServiceEventsConfiguration: &AwsServiceEventsConfigurationProperty{
//   		AwsService: jsii.String("awsService"),
//   		OnFailureConfiguration: &OnFailureConfigurationProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   		Pattern: jsii.String("pattern"),
//   	},
//   	PartnerEventsConfiguration: &PartnerEventsConfigurationProperty{
//   		OnFailureConfiguration: &OnFailureConfigurationProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   		PartnerBusKmsKeyIdentifier: jsii.String("partnerBusKmsKeyIdentifier"),
//   		PartnerEventSourceArn: jsii.String("partnerEventSourceArn"),
//   		Pattern: jsii.String("pattern"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-eventsource-eventsourceconfiguration.html
//
type CfnEventSourcePropsMixin_EventSourceConfigurationProperty struct {
	// Forward a single AWS service's events from the account's default event bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-eventsource-eventsourceconfiguration.html#cfn-eventsv2-eventsource-eventsourceconfiguration-awsserviceeventsconfiguration
	//
	AwsServiceEventsConfiguration interface{} `field:"optional" json:"awsServiceEventsConfiguration" yaml:"awsServiceEventsConfiguration"`
	// Forward a partner event source's events through a managed partner event bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-eventsource-eventsourceconfiguration.html#cfn-eventsv2-eventsource-eventsourceconfiguration-partnereventsconfiguration
	//
	PartnerEventsConfiguration interface{} `field:"optional" json:"partnerEventsConfiguration" yaml:"partnerEventsConfiguration"`
}

