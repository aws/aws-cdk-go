package awsses


// Properties for defining a `CfnConfigurationSetEventDestination`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationSetEventDestinationProps := &cfnConfigurationSetEventDestinationProps{
//   	configurationSetName: jsii.String("configurationSetName"),
//   	eventDestination: &eventDestinationProperty{
//   		matchingEventTypes: []*string{
//   			jsii.String("matchingEventTypes"),
//   		},
//
//   		// the properties below are optional
//   		cloudWatchDestination: &cloudWatchDestinationProperty{
//   			dimensionConfigurations: []interface{}{
//   				&dimensionConfigurationProperty{
//   					defaultDimensionValue: jsii.String("defaultDimensionValue"),
//   					dimensionName: jsii.String("dimensionName"),
//   					dimensionValueSource: jsii.String("dimensionValueSource"),
//   				},
//   			},
//   		},
//   		enabled: jsii.Boolean(false),
//   		kinesisFirehoseDestination: &kinesisFirehoseDestinationProperty{
//   			deliveryStreamArn: jsii.String("deliveryStreamArn"),
//   			iamRoleArn: jsii.String("iamRoleArn"),
//   		},
//   		name: jsii.String("name"),
//   		snsDestination: &snsDestinationProperty{
//   			topicArn: jsii.String("topicArn"),
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

