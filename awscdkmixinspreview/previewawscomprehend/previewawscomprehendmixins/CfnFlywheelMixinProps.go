package previewawscomprehendmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnFlywheelPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFlywheelMixinProps := &CfnFlywheelMixinProps{
//   	ActiveModelArn: jsii.String("activeModelArn"),
//   	DataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   	DataLakeS3Uri: jsii.String("dataLakeS3Uri"),
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
//   	FlywheelName: jsii.String("flywheelName"),
//   	ModelType: jsii.String("modelType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskConfig: &TaskConfigProperty{
//   		DocumentClassificationConfig: &DocumentClassificationConfigProperty{
//   			Labels: []*string{
//   				jsii.String("labels"),
//   			},
//   			Mode: jsii.String("mode"),
//   		},
//   		EntityRecognitionConfig: &EntityRecognitionConfigProperty{
//   			EntityTypes: []interface{}{
//   				&EntityTypesListItemProperty{
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		LanguageCode: jsii.String("languageCode"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-flywheel.html
//
type CfnFlywheelMixinProps struct {
	// The Amazon Resource Number (ARN) of the active model version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-flywheel.html#cfn-comprehend-flywheel-activemodelarn
	//
	ActiveModelArn *string `field:"optional" json:"activeModelArn" yaml:"activeModelArn"`
	// The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend permission to access the flywheel data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-flywheel.html#cfn-comprehend-flywheel-dataaccessrolearn
	//
	DataAccessRoleArn *string `field:"optional" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// Amazon S3 URI of the data lake location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-flywheel.html#cfn-comprehend-flywheel-datalakes3uri
	//
	DataLakeS3Uri *string `field:"optional" json:"dataLakeS3Uri" yaml:"dataLakeS3Uri"`
	// Data security configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-flywheel.html#cfn-comprehend-flywheel-datasecurityconfig
	//
	DataSecurityConfig interface{} `field:"optional" json:"dataSecurityConfig" yaml:"dataSecurityConfig"`
	// Name for the flywheel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-comprehend-flywheel.html#cfn-comprehend-flywheel-flywheelname
	//
	FlywheelName *string `field:"optional" json:"flywheelName" yaml:"flywheelName"`
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

