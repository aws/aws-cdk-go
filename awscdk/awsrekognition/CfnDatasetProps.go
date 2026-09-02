package awsrekognition

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
//   	DatasetType: jsii.String("datasetType"),
//
//   	// the properties below are optional
//   	ProjectArn: jsii.String("projectArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rekognition-dataset.html
//
type CfnDatasetProps struct {
	// The type of the dataset.
	//
	// Specify TRAIN to create a training dataset. Specify TEST to create a test dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rekognition-dataset.html#cfn-rekognition-dataset-datasettype
	//
	DatasetType *string `field:"required" json:"datasetType" yaml:"datasetType"`
	// The ARN of the project to which the dataset belongs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rekognition-dataset.html#cfn-rekognition-dataset-projectarn
	//
	ProjectArn *string `field:"optional" json:"projectArn" yaml:"projectArn"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rekognition-dataset.html#cfn-rekognition-dataset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

