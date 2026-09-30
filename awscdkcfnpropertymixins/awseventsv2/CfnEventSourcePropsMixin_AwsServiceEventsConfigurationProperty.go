package awseventsv2


// Forward a single AWS service's events from the account's default event bus.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   awsServiceEventsConfigurationProperty := &AwsServiceEventsConfigurationProperty{
//   	AwsService: jsii.String("awsService"),
//   	OnFailureConfiguration: &OnFailureConfigurationProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	Pattern: jsii.String("pattern"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-eventsource-awsserviceeventsconfiguration.html
//
type CfnEventSourcePropsMixin_AwsServiceEventsConfigurationProperty struct {
	// A single AWS service source identifier, for example aws.s3. Wildcards and lists are not allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-eventsource-awsserviceeventsconfiguration.html#cfn-eventsv2-eventsource-awsserviceeventsconfiguration-awsservice
	//
	AwsService *string `field:"optional" json:"awsService" yaml:"awsService"`
	// The destination for events that could not be forwarded by the managed forwarding target or, for partner event sources, the managed partner event bus.
	//
	// Arn is optional. An empty object removes a configured destination on update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-eventsource-awsserviceeventsconfiguration.html#cfn-eventsv2-eventsource-awsserviceeventsconfiguration-onfailureconfiguration
	//
	OnFailureConfiguration interface{} `field:"optional" json:"onFailureConfiguration" yaml:"onFailureConfiguration"`
	// A filter pattern, as a JSON string, that defines which events from the specified AWS service are forwarded to the event bus.
	//
	// Do not include source, account, or region as top-level fields. If you do not specify a pattern, all events from the service are forwarded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-eventsource-awsserviceeventsconfiguration.html#cfn-eventsv2-eventsource-awsserviceeventsconfiguration-pattern
	//
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
}

