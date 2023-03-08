package awsrekognition

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCollection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCollectionProps := &CfnCollectionProps{
//   	CollectionId: jsii.String("collectionId"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCollectionProps struct {
	// ID for the collection that you are creating.
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// A set of tags (key-value pairs) that you want to attach to the collection.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

