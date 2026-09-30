package awsscn


// The DataIntegrationFlow source parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataIntegrationFlowSourceProperty := &DataIntegrationFlowSourceProperty{
//   	SourceName: jsii.String("sourceName"),
//   	SourceType: jsii.String("sourceType"),
//
//   	// the properties below are optional
//   	DatasetSource: &DatasetSourceConfigurationProperty{
//   		DatasetIdentifier: jsii.String("datasetIdentifier"),
//
//   		// the properties below are optional
//   		Options: &DatasetOptionsProperty{
//   			DedupeRecords: jsii.Boolean(false),
//   			DedupeStrategy: &DedupeStrategyProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				FieldPriority: &FieldPriorityDedupeStrategyConfigurationProperty{
//   					Fields: []interface{}{
//   						&FieldPriorityDedupeFieldProperty{
//   							Name: jsii.String("name"),
//   							SortOrder: jsii.String("sortOrder"),
//   						},
//   					},
//   				},
//   			},
//   			LoadType: jsii.String("loadType"),
//   		},
//   	},
//   	S3Source: &S3SourceConfigurationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		Prefix: jsii.String("prefix"),
//
//   		// the properties below are optional
//   		Options: &S3OptionsProperty{
//   			FileType: jsii.String("fileType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowsource.html
//
type CfnDataIntegrationFlow_DataIntegrationFlowSourceProperty struct {
	// The source name that can be used as table alias in SQL transformation query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowsource.html#cfn-scn-dataintegrationflow-dataintegrationflowsource-sourcename
	//
	SourceName *string `field:"required" json:"sourceName" yaml:"sourceName"`
	// The source type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowsource.html#cfn-scn-dataintegrationflow-dataintegrationflowsource-sourcetype
	//
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
	// The dataset source configuration parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowsource.html#cfn-scn-dataintegrationflow-dataintegrationflowsource-datasetsource
	//
	DatasetSource interface{} `field:"optional" json:"datasetSource" yaml:"datasetSource"`
	// The S3 source configuration parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowsource.html#cfn-scn-dataintegrationflow-dataintegrationflowsource-s3source
	//
	S3Source interface{} `field:"optional" json:"s3Source" yaml:"s3Source"`
}

