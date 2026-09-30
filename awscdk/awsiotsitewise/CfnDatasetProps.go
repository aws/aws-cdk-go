package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDataset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDatasetProps := &CfnDatasetProps{
//   	DatasetName: jsii.String("datasetName"),
//
//   	// the properties below are optional
//   	DatasetConfig: &DatasetConfigProperty{
//   		Session: &SessionConfigProperty{
//   			SessionEndTime: jsii.String("sessionEndTime"),
//   			SessionStartTime: jsii.String("sessionStartTime"),
//   		},
//   	},
//   	DatasetDescription: jsii.String("datasetDescription"),
//   	DatasetSource: &DatasetSourceProperty{
//   		SourceFormat: jsii.String("sourceFormat"),
//   		SourceType: jsii.String("sourceType"),
//
//   		// the properties below are optional
//   		SourceDetail: &SourceDetailProperty{
//   			Kendra: &KendraSourceDetailProperty{
//   				KnowledgeBaseArn: jsii.String("knowledgeBaseArn"),
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   		},
//   	},
//   	DatasetType: jsii.String("datasetType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkspaceName: jsii.String("workspaceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dataset.html
//
type CfnDatasetProps struct {
	// The name of the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dataset.html#cfn-iotsitewise-dataset-datasetname
	//
	DatasetName *string `field:"required" json:"datasetName" yaml:"datasetName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dataset.html#cfn-iotsitewise-dataset-datasetconfig
	//
	DatasetConfig interface{} `field:"optional" json:"datasetConfig" yaml:"datasetConfig"`
	// A description about the dataset, and its functionality.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dataset.html#cfn-iotsitewise-dataset-datasetdescription
	//
	DatasetDescription *string `field:"optional" json:"datasetDescription" yaml:"datasetDescription"`
	// The data source for the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dataset.html#cfn-iotsitewise-dataset-datasetsource
	//
	DatasetSource interface{} `field:"optional" json:"datasetSource" yaml:"datasetSource"`
	// The type of the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dataset.html#cfn-iotsitewise-dataset-datasettype
	//
	DatasetType *string `field:"optional" json:"datasetType" yaml:"datasetType"`
	// A list of key-value pairs that contain metadata for the access policy.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dataset.html#cfn-iotsitewise-dataset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the workspace associated with the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-dataset.html#cfn-iotsitewise-dataset-workspacename
	//
	WorkspaceName *string `field:"optional" json:"workspaceName" yaml:"workspaceName"`
}

