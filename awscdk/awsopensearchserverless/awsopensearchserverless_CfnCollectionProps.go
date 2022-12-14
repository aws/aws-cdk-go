package awsopensearchserverless

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
	// `AWS::OpenSearchServerless::Collection.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::OpenSearchServerless::Collection.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::OpenSearchServerless::Collection.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::OpenSearchServerless::Collection.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

