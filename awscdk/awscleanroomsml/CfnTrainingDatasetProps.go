package awscleanroomsml

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTrainingDataset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrainingDatasetProps := &CfnTrainingDatasetProps{
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	TrainingData: []interface{}{
//   		&DatasetProperty{
//   			InputConfig: &DatasetInputConfigProperty{
//   				DataSource: &DataSourceProperty{
//   					GlueDataSource: &GlueDataSourceProperty{
//   						DatabaseName: jsii.String("databaseName"),
//   						TableName: jsii.String("tableName"),
//
//   						// the properties below are optional
//   						CatalogId: jsii.String("catalogId"),
//   					},
//   				},
//   				Schema: []interface{}{
//   					&ColumnSchemaProperty{
//   						ColumnName: jsii.String("columnName"),
//   						ColumnTypes: []*string{
//   							jsii.String("columnTypes"),
//   						},
//   					},
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-trainingdataset.html
//
type CfnTrainingDatasetProps struct {
	// The name of the training dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-trainingdataset.html#cfn-cleanroomsml-trainingdataset-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the IAM role that Clean Rooms ML can assume to read the data referred to in the `dataSource` field of each dataset.
	//
	// Passing a role across accounts is not allowed. If you pass a role that isn't in your account, you get an `AccessDeniedException` error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-trainingdataset.html#cfn-cleanroomsml-trainingdataset-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// An array of information that lists the Dataset objects, which specifies the dataset type and details on its location and schema.
	//
	// You must provide a role that has read access to these tables.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-trainingdataset.html#cfn-cleanroomsml-trainingdataset-trainingdata
	//
	TrainingData interface{} `field:"required" json:"trainingData" yaml:"trainingData"`
	// The description of the training dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-trainingdataset.html#cfn-cleanroomsml-trainingdataset-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The optional metadata that you apply to the resource to help you categorize and organize them.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	//
	// The following basic restrictions apply to tags:
	//
	// - Maximum number of tags per resource - 50.
	// - For each resource, each tag key must be unique, and each tag key can have only one value.
	// - Maximum key length - 128 Unicode characters in UTF-8.
	// - Maximum value length - 256 Unicode characters in UTF-8.
	// - If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : /
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cleanroomsml-trainingdataset.html#cfn-cleanroomsml-trainingdataset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

