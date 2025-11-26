package previewawssagemakermixins


// Configuration for Dataset Definition inputs.
//
// The Dataset Definition input must specify exactly one of either `AthenaDatasetDefinition` or `RedshiftDatasetDefinition` types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   datasetDefinitionProperty := &DatasetDefinitionProperty{
//   	AthenaDatasetDefinition: &AthenaDatasetDefinitionProperty{
//   		Catalog: jsii.String("catalog"),
//   		Database: jsii.String("database"),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		OutputCompression: jsii.String("outputCompression"),
//   		OutputFormat: jsii.String("outputFormat"),
//   		OutputS3Uri: jsii.String("outputS3Uri"),
//   		QueryString: jsii.String("queryString"),
//   		WorkGroup: jsii.String("workGroup"),
//   	},
//   	DataDistributionType: jsii.String("dataDistributionType"),
//   	InputMode: jsii.String("inputMode"),
//   	LocalPath: jsii.String("localPath"),
//   	RedshiftDatasetDefinition: &RedshiftDatasetDefinitionProperty{
//   		ClusterId: jsii.String("clusterId"),
//   		ClusterRoleArn: jsii.String("clusterRoleArn"),
//   		Database: jsii.String("database"),
//   		DbUser: jsii.String("dbUser"),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   		OutputCompression: jsii.String("outputCompression"),
//   		OutputFormat: jsii.String("outputFormat"),
//   		OutputS3Uri: jsii.String("outputS3Uri"),
//   		QueryString: jsii.String("queryString"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-datasetdefinition.html
//
type CfnProcessingJobPropsMixin_DatasetDefinitionProperty struct {
	// Configuration for Athena Dataset Definition input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-datasetdefinition.html#cfn-sagemaker-processingjob-datasetdefinition-athenadatasetdefinition
	//
	AthenaDatasetDefinition interface{} `field:"optional" json:"athenaDatasetDefinition" yaml:"athenaDatasetDefinition"`
	// Whether the generated dataset is `FullyReplicated` or `ShardedByS3Key` (default).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-datasetdefinition.html#cfn-sagemaker-processingjob-datasetdefinition-datadistributiontype
	//
	DataDistributionType *string `field:"optional" json:"dataDistributionType" yaml:"dataDistributionType"`
	// Whether to use `File` or `Pipe` input mode.
	//
	// In `File` (default) mode, Amazon SageMaker copies the data from the input source onto the local Amazon Elastic Block Store (Amazon EBS) volumes before starting your training algorithm. This is the most commonly used input mode. In `Pipe` mode, Amazon SageMaker streams input data from the source directly to your algorithm without using the EBS volume.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-datasetdefinition.html#cfn-sagemaker-processingjob-datasetdefinition-inputmode
	//
	InputMode *string `field:"optional" json:"inputMode" yaml:"inputMode"`
	// The local path where you want Amazon SageMaker to download the Dataset Definition inputs to run a processing job.
	//
	// `LocalPath` is an absolute path to the input data. This is a required parameter when `AppManaged` is `False` (default).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-datasetdefinition.html#cfn-sagemaker-processingjob-datasetdefinition-localpath
	//
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
	// Configuration for Redshift Dataset Definition input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-datasetdefinition.html#cfn-sagemaker-processingjob-datasetdefinition-redshiftdatasetdefinition
	//
	RedshiftDatasetDefinition interface{} `field:"optional" json:"redshiftDatasetDefinition" yaml:"redshiftDatasetDefinition"`
}

