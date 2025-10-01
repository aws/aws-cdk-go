package awssagemaker


// Configuration for Redshift Dataset Definition input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftDatasetDefinitionProperty := &RedshiftDatasetDefinitionProperty{
//   	ClusterId: jsii.String("clusterId"),
//   	ClusterRoleArn: jsii.String("clusterRoleArn"),
//   	Database: jsii.String("database"),
//   	DbUser: jsii.String("dbUser"),
//   	OutputFormat: jsii.String("outputFormat"),
//   	OutputS3Uri: jsii.String("outputS3Uri"),
//   	QueryString: jsii.String("queryString"),
//
//   	// the properties below are optional
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	OutputCompression: jsii.String("outputCompression"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html
//
type CfnProcessingJob_RedshiftDatasetDefinitionProperty struct {
	// The Redshift cluster Identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html#cfn-sagemaker-processingjob-redshiftdatasetdefinition-clusterid
	//
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// The IAM role attached to your Redshift cluster that Amazon SageMaker uses to generate datasets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html#cfn-sagemaker-processingjob-redshiftdatasetdefinition-clusterrolearn
	//
	ClusterRoleArn *string `field:"required" json:"clusterRoleArn" yaml:"clusterRoleArn"`
	// The name of the Redshift database used in Redshift query execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html#cfn-sagemaker-processingjob-redshiftdatasetdefinition-database
	//
	Database *string `field:"required" json:"database" yaml:"database"`
	// The database user name used in Redshift query execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html#cfn-sagemaker-processingjob-redshiftdatasetdefinition-dbuser
	//
	DbUser *string `field:"required" json:"dbUser" yaml:"dbUser"`
	// The data storage format for Redshift query results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html#cfn-sagemaker-processingjob-redshiftdatasetdefinition-outputformat
	//
	OutputFormat *string `field:"required" json:"outputFormat" yaml:"outputFormat"`
	// The location in Amazon S3 where the Redshift query results are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html#cfn-sagemaker-processingjob-redshiftdatasetdefinition-outputs3uri
	//
	OutputS3Uri *string `field:"required" json:"outputS3Uri" yaml:"outputS3Uri"`
	// The SQL query statements to be executed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html#cfn-sagemaker-processingjob-redshiftdatasetdefinition-querystring
	//
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt data from a Redshift execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html#cfn-sagemaker-processingjob-redshiftdatasetdefinition-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The compression used for Redshift query results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-redshiftdatasetdefinition.html#cfn-sagemaker-processingjob-redshiftdatasetdefinition-outputcompression
	//
	OutputCompression *string `field:"optional" json:"outputCompression" yaml:"outputCompression"`
}

