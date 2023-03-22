package awsivschat

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnLoggingConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLoggingConfigurationProps := &cfnLoggingConfigurationProps{
//   	destinationConfiguration: &destinationConfigurationProperty{
//   		cloudWatchLogs: &cloudWatchLogsDestinationConfigurationProperty{
//   			logGroupName: jsii.String("logGroupName"),
//   		},
//   		firehose: &firehoseDestinationConfigurationProperty{
//   			deliveryStreamName: jsii.String("deliveryStreamName"),
//   		},
//   		s3: &s3DestinationConfigurationProperty{
//   			bucketName: jsii.String("bucketName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLoggingConfigurationProps struct {
	// The DestinationConfiguration is a complex type that contains information about where chat content will be logged.
	DestinationConfiguration interface{} `field:"required" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// Logging-configuration name.
	//
	// The value does not need to be unique.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

