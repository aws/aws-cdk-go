package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnIntegrationResourceProperty`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIntegrationResourcePropertyProps := &CfnIntegrationResourcePropertyProps{
//   	ResourceArn: jsii.String("resourceArn"),
//
//   	// the properties below are optional
//   	SourceProcessingProperties: &SourceProcessingPropertiesProperty{
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetProcessingProperties: &TargetProcessingPropertiesProperty{
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		ConnectionName: jsii.String("connectionName"),
//   		EventBusArn: jsii.String("eventBusArn"),
//   		KmsArn: jsii.String("kmsArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integrationresourceproperty.html
//
type CfnIntegrationResourcePropertyProps struct {
	// The connection ARN of the source, or the database ARN of the target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integrationresourceproperty.html#cfn-glue-integrationresourceproperty-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The resource properties associated with the integration source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integrationresourceproperty.html#cfn-glue-integrationresourceproperty-sourceprocessingproperties
	//
	SourceProcessingProperties interface{} `field:"optional" json:"sourceProcessingProperties" yaml:"sourceProcessingProperties"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integrationresourceproperty.html#cfn-glue-integrationresourceproperty-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The resource properties associated with the integration target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-integrationresourceproperty.html#cfn-glue-integrationresourceproperty-targetprocessingproperties
	//
	TargetProcessingProperties interface{} `field:"optional" json:"targetProcessingProperties" yaml:"targetProcessingProperties"`
}

