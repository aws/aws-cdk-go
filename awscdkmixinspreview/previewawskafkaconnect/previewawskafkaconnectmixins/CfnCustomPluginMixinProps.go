package previewawskafkaconnectmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCustomPluginPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCustomPluginMixinProps := &CfnCustomPluginMixinProps{
//   	ContentType: jsii.String("contentType"),
//   	Description: jsii.String("description"),
//   	Location: &CustomPluginLocationProperty{
//   		S3Location: &S3LocationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			FileKey: jsii.String("fileKey"),
//   			ObjectVersion: jsii.String("objectVersion"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-customplugin.html
//
type CfnCustomPluginMixinProps struct {
	// The format of the plugin file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-customplugin.html#cfn-kafkaconnect-customplugin-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The description of the custom plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-customplugin.html#cfn-kafkaconnect-customplugin-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Information about the location of the custom plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-customplugin.html#cfn-kafkaconnect-customplugin-location
	//
	Location interface{} `field:"optional" json:"location" yaml:"location"`
	// The name of the custom plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-customplugin.html#cfn-kafkaconnect-customplugin-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kafkaconnect-customplugin.html#cfn-kafkaconnect-customplugin-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

