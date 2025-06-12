package awspinpointemail


// Properties for defining a `CfnConfigurationSetEventDestination`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationSetEventDestinationProps := &CfnConfigurationSetEventDestinationProps{
//   	ConfigurationSetName: jsii.String("configurationSetName"),
//   	EventDestinationName: jsii.String("eventDestinationName"),
//
//   	// the properties below are optional
//   	EventDestination: &EventDestinationProperty{
//   		MatchingEventTypes: []*string{
//   			jsii.String("matchingEventTypes"),
//   		},
//
//   		// the properties below are optional
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
//   		KinesisFirehoseDestination: &KinesisFirehoseDestinationProperty{
//   			DeliveryStreamArn: jsii.String("deliveryStreamArn"),
//   			IamRoleArn: jsii.String("iamRoleArn"),
//   		},
//   		PinpointDestination: &PinpointDestinationProperty{
//   			ApplicationArn: jsii.String("applicationArn"),
//   		},
//   		SnsDestination: &SnsDestinationProperty{
//   			TopicArn: jsii.String("topicArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-configurationseteventdestination.html
//
type CfnConfigurationSetEventDestinationProps struct {
	// The name of the configuration set that contains the event destination that you want to modify.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-configurationseteventdestination.html#cfn-pinpointemail-configurationseteventdestination-configurationsetname
	//
	ConfigurationSetName *string `field:"required" json:"configurationSetName" yaml:"configurationSetName"`
	// The name of the event destination that you want to modify.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-configurationseteventdestination.html#cfn-pinpointemail-configurationseteventdestination-eventdestinationname
	//
	EventDestinationName *string `field:"required" json:"eventDestinationName" yaml:"eventDestinationName"`
	// An object that defines the event destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-configurationseteventdestination.html#cfn-pinpointemail-configurationseteventdestination-eventdestination
	//
	EventDestination interface{} `field:"optional" json:"eventDestination" yaml:"eventDestination"`
}

