package awscomprehend

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnFlywheel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFlywheelProps := &cfnFlywheelProps{
//   	dataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   	dataLakeS3Uri: jsii.String("dataLakeS3Uri"),
//   	flywheelName: jsii.String("flywheelName"),
//
//   	// the properties below are optional
//   	activeModelArn: jsii.String("activeModelArn"),
//   	dataSecurityConfig: &dataSecurityConfigProperty{
//   		dataLakeKmsKeyId: jsii.String("dataLakeKmsKeyId"),
//   		modelKmsKeyId: jsii.String("modelKmsKeyId"),
//   		volumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   		vpcConfig: &vpcConfigProperty{
//   			securityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	modelType: jsii.String("modelType"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	taskConfig: &taskConfigProperty{
//   		languageCode: jsii.String("languageCode"),
//
//   		// the properties below are optional
//   		documentClassificationConfig: &documentClassificationConfigProperty{
//   			mode: jsii.String("mode"),
//
//   			// the properties below are optional
//   			labels: []*string{
//   				jsii.String("labels"),
//   			},
//   		},
//   		entityRecognitionConfig: &entityRecognitionConfigProperty{
//   			entityTypes: []interface{}{
//   				&entityTypesListItemProperty{
//   					type: jsii.String("type"),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnFlywheelProps struct {
	// The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend permission to access the flywheel data.
	DataAccessRoleArn *string `field:"required" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// Amazon S3 URI of the data lake location.
	DataLakeS3Uri *string `field:"required" json:"dataLakeS3Uri" yaml:"dataLakeS3Uri"`
	// Name for the flywheel.
	FlywheelName *string `field:"required" json:"flywheelName" yaml:"flywheelName"`
	// The Amazon Resource Number (ARN) of the active model version.
	ActiveModelArn *string `field:"optional" json:"activeModelArn" yaml:"activeModelArn"`
	// Data security configuration.
	DataSecurityConfig interface{} `field:"optional" json:"dataSecurityConfig" yaml:"dataSecurityConfig"`
	// Model type of the flywheel's model.
	ModelType *string `field:"optional" json:"modelType" yaml:"modelType"`
	// Tags associated with the endpoint being created.
	//
	// A tag is a key-value pair that adds metadata to the endpoint. For example, a tag with "Sales" as the key might be added to an endpoint to indicate its use by the sales department.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Configuration about the custom classifier associated with the flywheel.
	TaskConfig interface{} `field:"optional" json:"taskConfig" yaml:"taskConfig"`
}

