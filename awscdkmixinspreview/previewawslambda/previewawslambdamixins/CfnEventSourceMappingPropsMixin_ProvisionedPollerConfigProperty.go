package previewawslambdamixins


// The [provisioned mode](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventsourcemapping.html#invocation-eventsourcemapping-provisioned-mode) configuration for the event source. Use Provisioned Mode to customize the minimum and maximum number of event pollers for your event source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   provisionedPollerConfigProperty := &ProvisionedPollerConfigProperty{
//   	MaximumPollers: jsii.Number(123),
//   	MinimumPollers: jsii.Number(123),
//   	PollerGroupName: jsii.String("pollerGroupName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-provisionedpollerconfig.html
//
type CfnEventSourceMappingPropsMixin_ProvisionedPollerConfigProperty struct {
	// The maximum number of event pollers this event source can scale up to.
	//
	// For Amazon SQS events source mappings, default is 200, and minimum value allowed is 2. For Amazon MSK and self-managed Apache Kafka event source mappings, default is 200, and minimum value allowed is 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-provisionedpollerconfig.html#cfn-lambda-eventsourcemapping-provisionedpollerconfig-maximumpollers
	//
	MaximumPollers *float64 `field:"optional" json:"maximumPollers" yaml:"maximumPollers"`
	// The minimum number of event pollers this event source can scale down to.
	//
	// For Amazon SQS events source mappings, default is 2, and minimum 2 required. For Amazon MSK and self-managed Apache Kafka event source mappings, default is 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-provisionedpollerconfig.html#cfn-lambda-eventsourcemapping-provisionedpollerconfig-minimumpollers
	//
	MinimumPollers *float64 `field:"optional" json:"minimumPollers" yaml:"minimumPollers"`
	// (Amazon MSK and self-managed Apache Kafka) The name of the provisioned poller group.
	//
	// Use this option to group multiple ESMs within the event source's VPC to share Event Poller Unit (EPU) capacity. You can use this option to optimize Provisioned mode costs for your ESMs. You can group up to 100 ESMs per poller group and aggregate maximum pollers across all ESMs in a group cannot exceed 2000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-provisionedpollerconfig.html#cfn-lambda-eventsourcemapping-provisionedpollerconfig-pollergroupname
	//
	PollerGroupName *string `field:"optional" json:"pollerGroupName" yaml:"pollerGroupName"`
}

