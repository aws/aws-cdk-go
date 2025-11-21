package awsdatabrew

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnJob`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnJobProps := &CfnJobProps{
//   	Name: jsii.String("name"),
//   	RoleArn: jsii.String("roleArn"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	DatabaseOutputs: []interface{}{
//   		&DatabaseOutputProperty{
//   			DatabaseOptions: &DatabaseTableOutputOptionsProperty{
//   				TableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				TempDirectory: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					BucketOwner: jsii.String("bucketOwner"),
//   					Key: jsii.String("key"),
//   				},
//   			},
//   			GlueConnectionName: jsii.String("glueConnectionName"),
//
//   			// the properties below are optional
//   			DatabaseOutputMode: jsii.String("databaseOutputMode"),
//   		},
//   	},
//   	DataCatalogOutputs: []interface{}{
//   		&DataCatalogOutputProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			TableName: jsii.String("tableName"),
//
//   			// the properties below are optional
//   			CatalogId: jsii.String("catalogId"),
//   			DatabaseOptions: &DatabaseTableOutputOptionsProperty{
//   				TableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				TempDirectory: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					BucketOwner: jsii.String("bucketOwner"),
//   					Key: jsii.String("key"),
//   				},
//   			},
//   			Overwrite: jsii.Boolean(false),
//   			S3Options: &S3TableOutputOptionsProperty{
//   				Location: &S3LocationProperty{
//   					Bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					BucketOwner: jsii.String("bucketOwner"),
//   					Key: jsii.String("key"),
//   				},
//   			},
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
//   	OutputLocation: &OutputLocationProperty{
//   		Bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		BucketOwner: jsii.String("bucketOwner"),
//   		Key: jsii.String("key"),
//   	},
//   	Outputs: []interface{}{
//   		&OutputProperty{
//   			Location: &S3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				BucketOwner: jsii.String("bucketOwner"),
//   				Key: jsii.String("key"),
//   			},
//
//   			// the properties below are optional
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			Format: jsii.String("format"),
//   			FormatOptions: &OutputFormatOptionsProperty{
//   				Csv: &CsvOutputOptionsProperty{
//   					Delimiter: jsii.String("delimiter"),
//   				},
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
//
//   				// the properties below are optional
//   				Selectors: []interface{}{
//   					&ColumnSelectorProperty{
//   						Name: jsii.String("name"),
//   						Regex: jsii.String("regex"),
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
//   			EntityTypes: []*string{
//   				jsii.String("entityTypes"),
//   			},
//
//   			// the properties below are optional
//   			AllowedStatistics: &AllowedStatisticsProperty{
//   				Statistics: []*string{
//   					jsii.String("statistics"),
//   				},
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
//
//   		// the properties below are optional
//   		Version: jsii.String("version"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Timeout: jsii.Number(123),
//   	ValidationConfigurations: []interface{}{
//   		&ValidationConfigurationProperty{
//   			RulesetArn: jsii.String("rulesetArn"),
//
//   			// the properties below are optional
//   			ValidationMode: jsii.String("validationMode"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html
//
type CfnJobProps struct {
	// The unique name of the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the role to be assumed for this job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The job type of the job, which must be one of the following:.
	//
	// - `PROFILE` - A job to analyze a dataset, to determine its size, data types, data distribution, and more.
	// - `RECIPE` - A job to apply one or more transformations to a dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
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
	// List of validation configurations that are applied to the profile job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-databrew-job.html#cfn-databrew-job-validationconfigurations
	//
	ValidationConfigurations interface{} `field:"optional" json:"validationConfigurations" yaml:"validationConfigurations"`
}

