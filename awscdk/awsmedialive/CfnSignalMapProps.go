package awsmedialive


// Properties for defining a `CfnSignalMap`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSignalMapProps := &CfnSignalMapProps{
//   	DiscoveryEntryPointArn: jsii.String("discoveryEntryPointArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	CloudWatchAlarmTemplateGroupIdentifiers: []*string{
//   		jsii.String("cloudWatchAlarmTemplateGroupIdentifiers"),
//   	},
//   	Description: jsii.String("description"),
//   	EventBridgeRuleTemplateGroupIdentifiers: []*string{
//   		jsii.String("eventBridgeRuleTemplateGroupIdentifiers"),
//   	},
//   	ForceRediscovery: jsii.Boolean(false),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-signalmap.html
//
type CfnSignalMapProps struct {
	// A top-level supported Amazon Web Services resource ARN to discover a signal map from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-signalmap.html#cfn-medialive-signalmap-discoveryentrypointarn
	//
	DiscoveryEntryPointArn *string `field:"required" json:"discoveryEntryPointArn" yaml:"discoveryEntryPointArn"`
	// A resource's name.
	//
	// Names must be unique within the scope of a resource type in a specific region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-signalmap.html#cfn-medialive-signalmap-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A cloudwatch alarm template group's identifier.
	//
	// Can be either be its id or current name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-signalmap.html#cfn-medialive-signalmap-cloudwatchalarmtemplategroupidentifiers
	//
	CloudWatchAlarmTemplateGroupIdentifiers *[]*string `field:"optional" json:"cloudWatchAlarmTemplateGroupIdentifiers" yaml:"cloudWatchAlarmTemplateGroupIdentifiers"`
	// A resource's optional description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-signalmap.html#cfn-medialive-signalmap-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An eventbridge rule template group's identifier.
	//
	// Can be either be its id or current name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-signalmap.html#cfn-medialive-signalmap-eventbridgeruletemplategroupidentifiers
	//
	EventBridgeRuleTemplateGroupIdentifiers *[]*string `field:"optional" json:"eventBridgeRuleTemplateGroupIdentifiers" yaml:"eventBridgeRuleTemplateGroupIdentifiers"`
	// If true, will force a rediscovery of a signal map if an unchanged discoveryEntryPointArn is provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-signalmap.html#cfn-medialive-signalmap-forcerediscovery
	//
	// Default: - false.
	//
	ForceRediscovery interface{} `field:"optional" json:"forceRediscovery" yaml:"forceRediscovery"`
	// Represents the tags associated with a resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-signalmap.html#cfn-medialive-signalmap-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

