package previewawsdatabrewmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnJobPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnJobMixinProps := &CfnJobMixinProps{
//   	DatabaseOutputs: []interface{}{
//   		&DatabaseOutputProperty{
//   			DatabaseOptions: &DatabaseTableOutputOptionsProperty{
//   				TableName: jsii.String("tableName"),
//   				TempDirectory: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//   					BucketOwner: jsii.String("bucketOwner"),
//   					Key: jsii.String("key"),
//   				},
//   			},
//   			DatabaseOutputMode: jsii.String("databaseOutputMode"),
//   			GlueConnectionName: jsii.String("glueConnectionName"),
//   		},
//   	},
//   	DataCatalogOutputs: []interface{}{
//   		&DataCatalogOutputProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			DatabaseName: jsii.String("databaseName"),
//   			DatabaseOptions: &DatabaseTableOutputOptionsProperty{
//   				TableName: jsii.String("tableName"),
//   				TempDirectory: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//   					BucketOwner: jsii.String("bucketOwner"),
//   					Key: jsii.String("key"),
//   				},
//   			},
//   			Overwrite: jsii.Boolean(false),
//   			S3Options: &S3TableOutputOptionsProperty{
//   				Location: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//   					BucketOwner: jsii.String("bucketOwner"),
//   					Key: jsii.String("key"),
//   				},
//   			},
//   			TableName: jsii.String("tableName"),
//   		},
//   	},
//   	DatasetName: jsii.String("datasetName"),
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	EncryptionMode: jsii.String("encryptionMode"),
//   	JobSample: &JobSampleProperty{
//   		Mode: jsii.String("mode"),
//   		Size: jsii.Number(123),
//   	},
//   	LogSubscription: jsii.String("logSubscription"),
//   	MaxCapacity: jsii.Number(123),
//   	MaxRetries: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	OutputLocation: &OutputLocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		BucketOwner: jsii.String("bucketOwner"),
//   		Key: jsii.String("key"),
//   	},
//   	Outputs: []interface{}{
//   		&OutputProperty{
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			Format: jsii.String("format"),
//   			FormatOptions: &OutputFormatOptionsProperty{
//   				Csv: &CsvOutputOptionsProperty{
//   					Delimiter: jsii.String("delimiter"),
//   				},
//   			},
//   			Location: &S3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//   				BucketOwner: jsii.String("bucketOwner"),
//   				Key: jsii.String("key"),
//   			},
//   			MaxOutputFiles: jsii.Number(123),
//   			Overwrite: jsii.Boolean(false),
//   			PartitionColumns: []*string{
//   				jsii.String("partitionColumns"),
//   			},
//   		},
//   	},
//   	ProfileConfiguration: &ProfileConfigurationProperty{
//   		ColumnStatisticsConfigurations: []interface{}{
//   			&ColumnStatisticsConfigurationProperty{
//   				Selectors: []interface{}{
//   					&ColumnSelectorProperty{
//   						Name: jsii.String("name"),
//   						Regex: jsii.String("regex"),
//   					},
//   				},
//   				Statistics: &StatisticsConfigurationProperty{
//   					IncludedStatistics: []*string{
//   						jsii.String("includedStatistics"),
//   					},
//   					Overrides: []interface{}{
//   						&StatisticOverrideProperty{
//   							Parameters: map[string]*string{
//   								"parametersKey": jsii.String("parameters"),
//   							},
//   							Statistic: jsii.String("statistic"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		DatasetStatisticsConfiguration: &StatisticsConfigurationProperty{
//   			IncludedStatistics: []*string{
//   				jsii.String("includedStatistics"),
//   			},
//   			Overrides: []interface{}{
//   				&StatisticOverrideProperty{
//   					Parameters: map[string]*string{
//   						"parametersKey": jsii.String("parameters"),
//   					},
//   					Statistic: jsii.String("statistic"),
//   				},
//   			},
//   		},
//   		EntityDetectorConfiguration: &EntityDetectorConfigurationProperty{
//   			AllowedStatistics: &AllowedStatisticsProperty{
//   				Statistics: []*string{
//   					jsii.String("statistics"),
//   				},
//   			},
//   			EntityTypes: []*string{
//   				jsii.String("entityTypes"),
//   			},
//   		},
//   		ProfileColumns: []interface{}{
//   			&ColumnSelectorProperty{
//   				Name: jsii.String("name"),
//   				Regex: jsii.String("regex"),
//   			},
//   		},
//   	},
//   	ProjectName: jsii.String("projectName"),
//   	Recipe: &RecipeProperty{
//   		Name: jsii.String("name"),
//   		Version: jsii.String("version"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Timeout: jsii.Number(123),
//   	Type: jsii.String("type"),
//   	ValidationConfigurations: []interface{}{
//   		&ValidationConfigurationProperty{
//   			RulesetArn: jsii.String("rulesetArn"),
//   			ValidationMode: jsii.String("validationMode"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html
//
type CfnJobMixinProps struct {
	// Represents a list of JDBC database output objects which defines the output destination for a DataBrew recipe job to write into.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-databaseoutputs
	//
	DatabaseOutputs interface{} `field:"optional" json:"databaseOutputs" yaml:"databaseOutputs"`
	// One or more artifacts that represent the AWS Glue Data Catalog output from running the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-datacatalogoutputs
	//
	DataCatalogOutputs interface{} `field:"optional" json:"dataCatalogOutputs" yaml:"dataCatalogOutputs"`
	// A dataset that the job is to process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-datasetname
	//
	DatasetName *string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// The Amazon Resource Name (ARN) of an encryption key that is used to protect the job output.
	//
	// For more information, see [Encrypting data written by DataBrew jobs](https://docs.aws.amazon.com/databrew/latest/dg/encryption-security-configuration.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-encryptionkeyarn
	//
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// The encryption mode for the job, which can be one of the following:.
	//
	// - `SSE-KMS` - Server-side encryption with keys managed by AWS  .
	// - `SSE-S3` - Server-side encryption with keys managed by Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-encryptionmode
	//
	EncryptionMode *string `field:"optional" json:"encryptionMode" yaml:"encryptionMode"`
	// A sample configuration for profile jobs only, which determines the number of rows on which the profile job is run.
	//
	// If a `JobSample` value isn't provided, the default value is used. The default value is CUSTOM_ROWS for the mode parameter and 20,000 for the size parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-jobsample
	//
	JobSample interface{} `field:"optional" json:"jobSample" yaml:"jobSample"`
	// The current status of Amazon CloudWatch logging for the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-logsubscription
	//
	LogSubscription *string `field:"optional" json:"logSubscription" yaml:"logSubscription"`
	// The maximum number of nodes that can be consumed when the job processes data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-maxcapacity
	//
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The maximum number of times to retry the job after a job run fails.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-maxretries
	//
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// The unique name of the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The location in Amazon S3 where the job writes its output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-outputlocation
	//
	OutputLocation interface{} `field:"optional" json:"outputLocation" yaml:"outputLocation"`
	// One or more artifacts that represent output from running the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-outputs
	//
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
	// Configuration for profile jobs.
	//
	// Configuration can be used to select columns, do evaluations, and override default parameters of evaluations. When configuration is undefined, the profile job will apply default settings to all supported columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-profileconfiguration
	//
	ProfileConfiguration interface{} `field:"optional" json:"profileConfiguration" yaml:"profileConfiguration"`
	// The name of the project that the job is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-projectname
	//
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
	// A series of data transformation steps that the job runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-recipe
	//
	Recipe interface{} `field:"optional" json:"recipe" yaml:"recipe"`
	// The Amazon Resource Name (ARN) of the role to be assumed for this job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Metadata tags that have been applied to the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The job's timeout in minutes.
	//
	// A job that attempts to run longer than this timeout period ends with a status of `TIMEOUT` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-timeout
	//
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// The job type of the job, which must be one of the following:.
	//
	// - `PROFILE` - A job to analyze a dataset, to determine its size, data types, data distribution, and more.
	// - `RECIPE` - A job to apply one or more transformations to a dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// List of validation configurations that are applied to the profile job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-validationconfigurations
	//
	ValidationConfigurations interface{} `field:"optional" json:"validationConfigurations" yaml:"validationConfigurations"`
}

