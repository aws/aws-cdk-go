package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDatasetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var examples interface{}
//
//   cfnDatasetMixinProps := &CfnDatasetMixinProps{
//   	DatasetName: jsii.String("datasetName"),
//   	Description: jsii.String("description"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	SchemaType: jsii.String("schemaType"),
//   	Source: &DataSourceTypeProperty{
//   		InlineExamples: &InlineExamplesSourceProperty{
//   			Examples: []interface{}{
//   				examples,
//   			},
//   		},
//   		S3Source: &S3SourceProperty{
//   			S3Uri: jsii.String("s3Uri"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-dataset.html
//
type CfnDatasetMixinProps struct {
	// Human-readable name for the dataset.
	//
	// Unique within the account (case-insensitive). Immutable after creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-dataset.html#cfn-bedrockagentcore-dataset-datasetname
	//
	DatasetName *string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// A description of the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-dataset.html#cfn-bedrockagentcore-dataset-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Optional AWS KMS key ARN for SSE-KMS on service S3 writes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-dataset.html#cfn-bedrockagentcore-dataset-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Versioned schema type governing the structure of examples.
	//
	// Immutable after creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-dataset.html#cfn-bedrockagentcore-dataset-schematype
	//
	SchemaType *string `field:"optional" json:"schemaType" yaml:"schemaType"`
	// Source of initial examples.
	//
	// Provide either inline examples or an S3 URI pointing to a JSONL file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-dataset.html#cfn-bedrockagentcore-dataset-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// A list of tags to assign to the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-dataset.html#cfn-bedrockagentcore-dataset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

