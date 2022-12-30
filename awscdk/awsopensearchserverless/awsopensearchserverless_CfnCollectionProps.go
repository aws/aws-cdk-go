package awsopensearchserverless

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCollection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCollectionProps := &cfnCollectionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnCollectionProps struct {
	// The name of the collection.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the collection.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An arbitrary set of tags (key–value pairs) to associate with the collection.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of collection.
	//
	// Possible values are `SEARCH` and `TIMESERIES` . For more information, see [Choosing a collection type](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-overview.html#serverless-usecase) .
	Type *string `field:"optional" json:"type" yaml:"type"`
}

