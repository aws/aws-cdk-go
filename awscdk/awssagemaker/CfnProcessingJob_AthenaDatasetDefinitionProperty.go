package awssagemaker


// Configuration for Athena Dataset Definition input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   athenaDatasetDefinitionProperty := &AthenaDatasetDefinitionProperty{
//   	Catalog: jsii.String("catalog"),
//   	Database: jsii.String("database"),
//   	OutputFormat: jsii.String("outputFormat"),
//   	OutputS3Uri: jsii.String("outputS3Uri"),
//   	QueryString: jsii.String("queryString"),
//
//   	// the properties below are optional
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	OutputCompression: jsii.String("outputCompression"),
//   	WorkGroup: jsii.String("workGroup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-athenadatasetdefinition.html
//
type CfnProcessingJob_AthenaDatasetDefinitionProperty struct {
	// The name of the data catalog used in Athena query execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-athenadatasetdefinition.html#cfn-sagemaker-processingjob-athenadatasetdefinition-catalog
	//
	Catalog *string `field:"required" json:"catalog" yaml:"catalog"`
	// The name of the database used in the Athena query execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-athenadatasetdefinition.html#cfn-sagemaker-processingjob-athenadatasetdefinition-database
	//
	Database *string `field:"required" json:"database" yaml:"database"`
	// The data storage format for Athena query results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-athenadatasetdefinition.html#cfn-sagemaker-processingjob-athenadatasetdefinition-outputformat
	//
	OutputFormat *string `field:"required" json:"outputFormat" yaml:"outputFormat"`
	// The location in Amazon S3 where Athena query results are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-athenadatasetdefinition.html#cfn-sagemaker-processingjob-athenadatasetdefinition-outputs3uri
	//
	OutputS3Uri *string `field:"required" json:"outputS3Uri" yaml:"outputS3Uri"`
	// The SQL query statements, to be executed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-athenadatasetdefinition.html#cfn-sagemaker-processingjob-athenadatasetdefinition-querystring
	//
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// The AWS Key Management Service ( AWS KMS) key that Amazon SageMaker uses to encrypt data generated from an Athena query execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-athenadatasetdefinition.html#cfn-sagemaker-processingjob-athenadatasetdefinition-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The compression used for Athena query results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-athenadatasetdefinition.html#cfn-sagemaker-processingjob-athenadatasetdefinition-outputcompression
	//
	OutputCompression *string `field:"optional" json:"outputCompression" yaml:"outputCompression"`
	// The name of the workgroup in which the Athena query is being started.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-athenadatasetdefinition.html#cfn-sagemaker-processingjob-athenadatasetdefinition-workgroup
	//
	WorkGroup *string `field:"optional" json:"workGroup" yaml:"workGroup"`
}

