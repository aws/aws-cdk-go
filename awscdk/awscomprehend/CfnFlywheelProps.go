package awscomprehend

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFlywheel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlywheelProps := &CfnFlywheelProps{
//   	DataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   	DataLakeS3Uri: jsii.String("dataLakeS3Uri"),
//   	FlywheelName: jsii.String("flywheelName"),
//
//   	// the properties below are optional
//   	ActiveModelArn: jsii.String("activeModelArn"),
//   	DataSecurityConfig: &DataSecurityConfigProperty{
//   		DataLakeKmsKeyId: jsii.String("dataLakeKmsKeyId"),
//   		ModelKmsKeyId: jsii.String("modelKmsKeyId"),
//   		VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		VpcConfig: &VpcConfigProperty{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	ModelType: jsii.String("modelType"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskConfig: &TaskConfigProperty{
//   		LanguageCode: jsii.String("languageCode"),
//
//   		// the properties below are optional
//   		DocumentClassificationConfig: &DocumentClassificationConfigProperty{
//   			Mode: jsii.String("mode"),
//
//   			// the properties below are optional
//   			Labels: []*string{
//   				jsii.String("labels"),
//   			},
//   		},
//   		EntityRecognitionConfig: &EntityRecognitionConfigProperty{
//   			EntityTypes: []interface{}{
//   				&EntityTypesListItemProperty{
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-flywheel.html
//
type CfnFlywheelProps struct {
	// The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend permission to access the flywheel data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-flywheel.html#cfn-comprehend-flywheel-dataaccessrolearn
	//
	DataAccessRoleArn *string `field:"required" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// Amazon S3 URI of the data lake location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-flywheel.html#cfn-comprehend-flywheel-datalakes3uri
	//
	DataLakeS3Uri *string `field:"required" json:"dataLakeS3Uri" yaml:"dataLakeS3Uri"`
	// Name for the flywheel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-flywheel.html#cfn-comprehend-flywheel-flywheelname
	//
	FlywheelName *string `field:"required" json:"flywheelName" yaml:"flywheelName"`
	// The Amazon Resource Number (ARN) of the active model version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-flywheel.html#cfn-comprehend-flywheel-activemodelarn
	//
	ActiveModelArn *string `field:"optional" json:"activeModelArn" yaml:"activeModelArn"`
	// Data security configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-flywheel.html#cfn-comprehend-flywheel-datasecurityconfig
	//
	DataSecurityConfig interface{} `field:"optional" json:"dataSecurityConfig" yaml:"dataSecurityConfig"`
	// Model type of the flywheel's model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-flywheel.html#cfn-comprehend-flywheel-modeltype
	//
	ModelType *string `field:"optional" json:"modelType" yaml:"modelType"`
	// Tags associated with the endpoint being created.
	//
	// A tag is a key-value pair that adds metadata to the endpoint. For example, a tag with "Sales" as the key might be added to an endpoint to indicate its use by the sales department.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-flywheel.html#cfn-comprehend-flywheel-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Configuration about the model associated with a flywheel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-flywheel.html#cfn-comprehend-flywheel-taskconfig
	//
	TaskConfig interface{} `field:"optional" json:"taskConfig" yaml:"taskConfig"`
}

