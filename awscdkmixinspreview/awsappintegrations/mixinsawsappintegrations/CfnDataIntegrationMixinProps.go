package mixinsawsappintegrations

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDataIntegrationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var filters interface{}
//   var objectConfiguration interface{}
//
//   cfnDataIntegrationMixinProps := &CfnDataIntegrationMixinProps{
//   	Description: jsii.String("description"),
//   	FileConfiguration: &FileConfigurationProperty{
//   		Filters: filters,
//   		Folders: []*string{
//   			jsii.String("folders"),
//   		},
//   	},
//   	KmsKey: jsii.String("kmsKey"),
//   	Name: jsii.String("name"),
//   	ObjectConfiguration: objectConfiguration,
//   	ScheduleConfig: &ScheduleConfigProperty{
//   		FirstExecutionFrom: jsii.String("firstExecutionFrom"),
//   		Object: jsii.String("object"),
//   		ScheduleExpression: jsii.String("scheduleExpression"),
//   	},
//   	SourceUri: jsii.String("sourceUri"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-dataintegration.html
//
type CfnDataIntegrationMixinProps struct {
	// A description of the DataIntegration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-dataintegration.html#cfn-appintegrations-dataintegration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The configuration for what files should be pulled from the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-dataintegration.html#cfn-appintegrations-dataintegration-fileconfiguration
	//
	FileConfiguration interface{} `field:"optional" json:"fileConfiguration" yaml:"fileConfiguration"`
	// The KMS key for the DataIntegration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-dataintegration.html#cfn-appintegrations-dataintegration-kmskey
	//
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The name of the DataIntegration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-dataintegration.html#cfn-appintegrations-dataintegration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The configuration for what data should be pulled from the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-dataintegration.html#cfn-appintegrations-dataintegration-objectconfiguration
	//
	ObjectConfiguration interface{} `field:"optional" json:"objectConfiguration" yaml:"objectConfiguration"`
	// The name of the data and how often it should be pulled from the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-dataintegration.html#cfn-appintegrations-dataintegration-scheduleconfig
	//
	ScheduleConfig interface{} `field:"optional" json:"scheduleConfig" yaml:"scheduleConfig"`
	// The URI of the data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-dataintegration.html#cfn-appintegrations-dataintegration-sourceuri
	//
	SourceUri *string `field:"optional" json:"sourceUri" yaml:"sourceUri"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appintegrations-dataintegration.html#cfn-appintegrations-dataintegration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

