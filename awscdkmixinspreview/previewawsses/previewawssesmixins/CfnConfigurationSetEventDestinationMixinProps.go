package previewawssesmixins


// Properties for CfnConfigurationSetEventDestinationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfigurationSetEventDestinationMixinProps := &CfnConfigurationSetEventDestinationMixinProps{
//   	ConfigurationSetName: jsii.String("configurationSetName"),
//   	EventDestination: &EventDestinationProperty{
//   		CloudWatchDestination: &CloudWatchDestinationProperty{
//   			DimensionConfigurations: []interface{}{
//   				&DimensionConfigurationProperty{
//   					DefaultDimensionValue: jsii.String("defaultDimensionValue"),
//   					DimensionName: jsii.String("dimensionName"),
//   					DimensionValueSource: jsii.String("dimensionValueSource"),
//   				},
//   			},
//   		},
//   		Enabled: jsii.Boolean(false),
//   		EventBridgeDestination: &EventBridgeDestinationProperty{
//   			EventBusArn: jsii.String("eventBusArn"),
//   		},
//   		KinesisFirehoseDestination: &KinesisFirehoseDestinationProperty{
//   			DeliveryStreamArn: jsii.String("deliveryStreamArn"),
//   			IamRoleArn: jsii.String("iamRoleArn"),
//   		},
//   		MatchingEventTypes: []*string{
//   			jsii.String("matchingEventTypes"),
//   		},
//   		Name: jsii.String("name"),
//   		SnsDestination: &SnsDestinationProperty{
//   			TopicArn: jsii.String("topicArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationseteventdestination.html
//
type CfnConfigurationSetEventDestinationMixinProps struct {
	// The name of the configuration set that contains the event destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationseteventdestination.html#cfn-ses-configurationseteventdestination-configurationsetname
	//
	ConfigurationSetName *string `field:"optional" json:"configurationSetName" yaml:"configurationSetName"`
	// An object that defines the event destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationseteventdestination.html#cfn-ses-configurationseteventdestination-eventdestination
	//
	EventDestination interface{} `field:"optional" json:"eventDestination" yaml:"eventDestination"`
}

