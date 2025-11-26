package previewawssagemakermixins


// Configuration for Redshift Dataset Definition input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   redshiftDatasetDefinitionProperty := &RedshiftDatasetDefinitionProperty{
//   	ClusterId: jsii.String("clusterId"),
//   	ClusterRoleArn: jsii.String("clusterRoleArn"),
//   	Database: jsii.String("database"),
//   	DbUser: jsii.String("dbUser"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	OutputCompression: jsii.String("outputCompression"),
//   	OutputFormat: jsii.String("outputFormat"),
//   	OutputS3Uri: jsii.String("outputS3Uri"),
//   	QueryString: jsii.String("queryString"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html
//
type CfnProcessingJobPropsMixin_RedshiftDatasetDefinitionProperty struct {
	// The Redshift cluster Identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html#cfn-sagemaker-processingjob-redshiftdatasetdefinition-clusterid
	//
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// The IAM role attached to your Redshift cluster that Amazon SageMaker uses to generate datasets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html#cfn-sagemaker-processingjob-redshiftdatasetdefinition-clusterrolearn
	//
	ClusterRoleArn *string `field:"optional" json:"clusterRoleArn" yaml:"clusterRoleArn"`
	// The name of the Redshift database used in Redshift query execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html#cfn-sagemaker-processingjob-redshiftdatasetdefinition-database
	//
	Database *string `field:"optional" json:"database" yaml:"database"`
	// The database user name used in Redshift query execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html#cfn-sagemaker-processingjob-redshiftdatasetdefinition-dbuser
	//
	DbUser *string `field:"optional" json:"dbUser" yaml:"dbUser"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt data from a Redshift execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html#cfn-sagemaker-processingjob-redshiftdatasetdefinition-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The compression used for Redshift query results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html#cfn-sagemaker-processingjob-redshiftdatasetdefinition-outputcompression
	//
	OutputCompression *string `field:"optional" json:"outputCompression" yaml:"outputCompression"`
	// The data storage format for Redshift query results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html#cfn-sagemaker-processingjob-redshiftdatasetdefinition-outputformat
	//
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// The location in Amazon S3 where the Redshift query results are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html#cfn-sagemaker-processingjob-redshiftdatasetdefinition-outputs3uri
	//
	OutputS3Uri *string `field:"optional" json:"outputS3Uri" yaml:"outputS3Uri"`
	// The SQL query statements to be executed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html#cfn-sagemaker-processingjob-redshiftdatasetdefinition-querystring
	//
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
}

