package awsappintegrations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataIntegration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var filters interface{}
//   var objectConfiguration interface{}
//
//   cfnDataIntegrationProps := &CfnDataIntegrationProps{
//   	KmsKey: jsii.String("kmsKey"),
//   	Name: jsii.String("name"),
//   	ScheduleConfig: &ScheduleConfigProperty{
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//
//   		// the properties below are optional
//   		FirstExecutionFrom: jsii.String("firstExecutionFrom"),
//   		Object: jsii.String("object"),
//   	},
//   	SourceUri: jsii.String("sourceUri"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	FileConfiguration: &FileConfigurationProperty{
//   		Folders: []*string{
//   			jsii.String("folders"),
//   		},
//
//   		// the properties below are optional
//   		Filters: filters,
//   	},
//   	ObjectConfiguration: objectConfiguration,
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-dataintegration.html
//
type CfnDataIntegrationProps struct {
	// The KMS key for the DataIntegration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-dataintegration.html#cfn-appintegrations-dataintegration-kmskey
	//
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
	// The name of the DataIntegration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-dataintegration.html#cfn-appintegrations-dataintegration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the data and how often it should be pulled from the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-dataintegration.html#cfn-appintegrations-dataintegration-scheduleconfig
	//
	ScheduleConfig interface{} `field:"required" json:"scheduleConfig" yaml:"scheduleConfig"`
	// The URI of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-dataintegration.html#cfn-appintegrations-dataintegration-sourceuri
	//
	SourceUri *string `field:"required" json:"sourceUri" yaml:"sourceUri"`
	// A description of the DataIntegration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-dataintegration.html#cfn-appintegrations-dataintegration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The configuration for what files should be pulled from the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-dataintegration.html#cfn-appintegrations-dataintegration-fileconfiguration
	//
	FileConfiguration interface{} `field:"optional" json:"fileConfiguration" yaml:"fileConfiguration"`
	// The configuration for what data should be pulled from the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-dataintegration.html#cfn-appintegrations-dataintegration-objectconfiguration
	//
	ObjectConfiguration interface{} `field:"optional" json:"objectConfiguration" yaml:"objectConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-dataintegration.html#cfn-appintegrations-dataintegration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

