package awsses


// Properties for defining a `CfnConfigurationSetEventDestination`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationSetEventDestinationProps := &CfnConfigurationSetEventDestinationProps{
//   	ConfigurationSetName: jsii.String("configurationSetName"),
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
//   		Name: jsii.String("name"),
//   		SnsDestination: &SnsDestinationProperty{
//   			TopicArn: jsii.String("topicArn"),
//   		},
//   	},
//   }
//
type CfnConfigurationSetEventDestinationProps struct {
	// The name of the configuration set that contains the event destination.
	ConfigurationSetName *string `field:"required" json:"configurationSetName" yaml:"configurationSetName"`
	// The event destination object.
	EventDestination interface{} `field:"required" json:"eventDestination" yaml:"eventDestination"`
}

