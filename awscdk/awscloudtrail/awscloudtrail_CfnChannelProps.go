package awscloudtrail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnChannelProps := &cfnChannelProps{
//   	destinations: []interface{}{
//   		&destinationProperty{
//   			location: jsii.String("location"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	source: jsii.String("source"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnChannelProps struct {
	// One or more event data stores to which events arriving through a channel will be logged.
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// The name of the channel.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name of the partner or external event source.
	//
	// You cannot change this name after you create the channel. A maximum of one channel is allowed per source.
	//
	// A source can be either `Custom` for all valid non- AWS events, or the name of a partner event source. For information about the source names for available partners, see [Additional information about integration partners](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-event-data-store-integration.html#cloudtrail-lake-partner-information) in the CloudTrail User Guide.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// A list of tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

