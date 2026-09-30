package awsscn


// The DataIntegrationFlow source parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dataIntegrationFlowSourceProperty := &DataIntegrationFlowSourceProperty{
//   	DatasetSource: &DatasetSourceConfigurationProperty{
//   		DatasetIdentifier: jsii.String("datasetIdentifier"),
//   		Options: &DatasetOptionsProperty{
//   			DedupeRecords: jsii.Boolean(false),
//   			DedupeStrategy: &DedupeStrategyProperty{
//   				FieldPriority: &FieldPriorityDedupeStrategyConfigurationProperty{
//   					Fields: []interface{}{
//   						&FieldPriorityDedupeFieldProperty{
//   							Name: jsii.String("name"),
//   							SortOrder: jsii.String("sortOrder"),
//   						},
//   					},
//   				},
//   				Type: jsii.String("type"),
//   			},
//   			LoadType: jsii.String("loadType"),
//   		},
//   	},
//   	S3Source: &S3SourceConfigurationProperty{
//   		BucketName: jsii.String("bucketName"),
//   		Options: &S3OptionsProperty{
//   			FileType: jsii.String("fileType"),
//   		},
//   		Prefix: jsii.String("prefix"),
//   	},
//   	SourceName: jsii.String("sourceName"),
//   	SourceType: jsii.String("sourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowsource.html
//
type CfnDataIntegrationFlowPropsMixin_DataIntegrationFlowSourceProperty struct {
	// The dataset source configuration parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowsource.html#cfn-scn-dataintegrationflow-dataintegrationflowsource-datasetsource
	//
	DatasetSource interface{} `field:"optional" json:"datasetSource" yaml:"datasetSource"`
	// The S3 source configuration parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowsource.html#cfn-scn-dataintegrationflow-dataintegrationflowsource-s3source
	//
	S3Source interface{} `field:"optional" json:"s3Source" yaml:"s3Source"`
	// The source name that can be used as table alias in SQL transformation query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowsource.html#cfn-scn-dataintegrationflow-dataintegrationflowsource-sourcename
	//
	SourceName *string `field:"optional" json:"sourceName" yaml:"sourceName"`
	// The source type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scn-dataintegrationflow-dataintegrationflowsource.html#cfn-scn-dataintegrationflow-dataintegrationflowsource-sourcetype
	//
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

