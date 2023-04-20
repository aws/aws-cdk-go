package awsappintegrations

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDataIntegration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDataIntegrationProps := &CfnDataIntegrationProps{
//   	KmsKey: jsii.String("kmsKey"),
//   	Name: jsii.String("name"),
//   	ScheduleConfig: &ScheduleConfigProperty{
//   		FirstExecutionFrom: jsii.String("firstExecutionFrom"),
//   		Object: jsii.String("object"),
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	SourceUri: jsii.String("sourceUri"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDataIntegrationProps struct {
	// The KMS key for the DataIntegration.
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
	// The name of the DataIntegration.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the data and how often it should be pulled from the source.
	ScheduleConfig interface{} `field:"required" json:"scheduleConfig" yaml:"scheduleConfig"`
	// The URI of the data source.
	SourceUri *string `field:"required" json:"sourceUri" yaml:"sourceUri"`
	// A description of the DataIntegration.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

