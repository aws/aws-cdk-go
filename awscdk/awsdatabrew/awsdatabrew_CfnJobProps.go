package awsdatabrew

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnJob`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnJobProps := &cfnJobProps{
//   	name: jsii.String("name"),
//   	roleArn: jsii.String("roleArn"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	databaseOutputs: []interface{}{
//   		&databaseOutputProperty{
//   			databaseOptions: &databaseTableOutputOptionsProperty{
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				tempDirectory: &s3LocationProperty{
//   					bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					bucketOwner: jsii.String("bucketOwner"),
//   					key: jsii.String("key"),
//   				},
//   			},
//   			glueConnectionName: jsii.String("glueConnectionName"),
//
//   			// the properties below are optional
//   			databaseOutputMode: jsii.String("databaseOutputMode"),
//   		},
//   	},
//   	dataCatalogOutputs: []interface{}{
//   		&dataCatalogOutputProperty{
//   			databaseName: jsii.String("databaseName"),
//   			tableName: jsii.String("tableName"),
//
//   			// the properties below are optional
//   			catalogId: jsii.String("catalogId"),
//   			databaseOptions: &databaseTableOutputOptionsProperty{
//   				tableName: jsii.String("tableName"),
//
//   				// the properties below are optional
//   				tempDirectory: &s3LocationProperty{
//   					bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					bucketOwner: jsii.String("bucketOwner"),
//   					key: jsii.String("key"),
//   				},
//   			},
//   			overwrite: jsii.Boolean(false),
//   			s3Options: &s3TableOutputOptionsProperty{
//   				location: &s3LocationProperty{
//   					bucket: jsii.String("bucket"),
//
//   					// the properties below are optional
//   					bucketOwner: jsii.String("bucketOwner"),
//   					key: jsii.String("key"),
//   				},
//   			},
//   		},
//   	},
//   	datasetName: jsii.String("datasetName"),
//   	encryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	encryptionMode: jsii.String("encryptionMode"),
//   	jobSample: &jobSampleProperty{
//   		mode: jsii.String("mode"),
//   		size: jsii.Number(123),
//   	},
//   	logSubscription: jsii.String("logSubscription"),
//   	maxCapacity: jsii.Number(123),
//   	maxRetries: jsii.Number(123),
//   	outputLocation: &outputLocationProperty{
//   		bucket: jsii.String("bucket"),
//
//   		// the properties below are optional
//   		bucketOwner: jsii.String("bucketOwner"),
//   		key: jsii.String("key"),
//   	},
//   	outputs: []interface{}{
//   		&outputProperty{
//   			location: &s3LocationProperty{
//   				bucket: jsii.String("bucket"),
//
//   				// the properties below are optional
//   				bucketOwner: jsii.String("bucketOwner"),
//   				key: jsii.String("key"),
//   			},
//
//   			// the properties below are optional
//   			compressionFormat: jsii.String("compressionFormat"),
//   			format: jsii.String("format"),
//   			formatOptions: &outputFormatOptionsProperty{
//   				csv: &csvOutputOptionsProperty{
//   					delimiter: jsii.String("delimiter"),
//   				},
//   			},
//   			maxOutputFiles: jsii.Number(123),
//   			overwrite: jsii.Boolean(false),
//   			partitionColumns: []*string{
//   				jsii.String("partitionColumns"),
//   			},
//   		},
//   	},
//   	profileConfiguration: &profileConfigurationProperty{
//   		columnStatisticsConfigurations: []interface{}{
//   			&columnStatisticsConfigurationProperty{
//   				statistics: &statisticsConfigurationProperty{
//   					includedStatistics: []*string{
//   						jsii.String("includedStatistics"),
//   					},
//   					overrides: []interface{}{
//   						&statisticOverrideProperty{
//   							parameters: map[string]*string{
//   								"parametersKey": jsii.String("parameters"),
//   							},
//   							statistic: jsii.String("statistic"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				selectors: []interface{}{
//   					&columnSelectorProperty{
//   						name: jsii.String("name"),
//   						regex: jsii.String("regex"),
//   					},
//   				},
//   			},
//   		},
//   		datasetStatisticsConfiguration: &statisticsConfigurationProperty{
//   			includedStatistics: []*string{
//   				jsii.String("includedStatistics"),
//   			},
//   			overrides: []interface{}{
//   				&statisticOverrideProperty{
//   					parameters: map[string]*string{
//   						"parametersKey": jsii.String("parameters"),
//   					},
//   					statistic: jsii.String("statistic"),
//   				},
//   			},
//   		},
//   		entityDetectorConfiguration: &entityDetectorConfigurationProperty{
//   			entityTypes: []*string{
//   				jsii.String("entityTypes"),
//   			},
//
//   			// the properties below are optional
//   			allowedStatistics: &allowedStatisticsProperty{
//   				statistics: []*string{
//   					jsii.String("statistics"),
//   				},
//   			},
//   		},
//   		profileColumns: []interface{}{
//   			&columnSelectorProperty{
//   				name: jsii.String("name"),
//   				regex: jsii.String("regex"),
//   			},
//   		},
//   	},
//   	projectName: jsii.String("projectName"),
//   	recipe: &recipeProperty{
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		version: jsii.String("version"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	timeout: jsii.Number(123),
//   	validationConfigurations: []interface{}{
//   		&validationConfigurationProperty{
//   			rulesetArn: jsii.String("rulesetArn"),
//
//   			// the properties below are optional
//   			validationMode: jsii.String("validationMode"),
//   		},
//   	},
//   }
//
type CfnJobProps struct {
	// The unique name of the job.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Resource Name (ARN) of the role to be assumed for this job.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The job type of the job, which must be one of the following:.
	//
	// - `PROFILE` - A job to analyze a dataset, to determine its size, data types, data distribution, and more.
	// - `RECIPE` - A job to apply one or more transformations to a dataset.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Represents a list of JDBC database output objects which defines the output destination for a DataBrew recipe job to write into.
	DatabaseOutputs interface{} `field:"optional" json:"databaseOutputs" yaml:"databaseOutputs"`
	// One or more artifacts that represent the AWS Glue Data Catalog output from running the job.
	DataCatalogOutputs interface{} `field:"optional" json:"dataCatalogOutputs" yaml:"dataCatalogOutputs"`
	// A dataset that the job is to process.
	DatasetName *string `field:"optional" json:"datasetName" yaml:"datasetName"`
	// The Amazon Resource Name (ARN) of an encryption key that is used to protect the job output.
	//
	// For more information, see [Encrypting data written by DataBrew jobs](https://docs.aws.amazon.com/databrew/latest/dg/encryption-security-configuration.html)
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// The encryption mode for the job, which can be one of the following:.
	//
	// - `SSE-KMS` - Server-side encryption with keys managed by AWS KMS .
	// - `SSE-S3` - Server-side encryption with keys managed by Amazon S3.
	EncryptionMode *string `field:"optional" json:"encryptionMode" yaml:"encryptionMode"`
	// A sample configuration for profile jobs only, which determines the number of rows on which the profile job is run.
	//
	// If a `JobSample` value isn't provided, the default value is used. The default value is CUSTOM_ROWS for the mode parameter and 20,000 for the size parameter.
	JobSample interface{} `field:"optional" json:"jobSample" yaml:"jobSample"`
	// The current status of Amazon CloudWatch logging for the job.
	LogSubscription *string `field:"optional" json:"logSubscription" yaml:"logSubscription"`
	// The maximum number of nodes that can be consumed when the job processes data.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The maximum number of times to retry the job after a job run fails.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// `AWS::DataBrew::Job.OutputLocation`.
	OutputLocation interface{} `field:"optional" json:"outputLocation" yaml:"outputLocation"`
	// One or more artifacts that represent output from running the job.
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
	// Configuration for profile jobs.
	//
	// Configuration can be used to select columns, do evaluations, and override default parameters of evaluations. When configuration is undefined, the profile job will apply default settings to all supported columns.
	ProfileConfiguration interface{} `field:"optional" json:"profileConfiguration" yaml:"profileConfiguration"`
	// The name of the project that the job is associated with.
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
	// A series of data transformation steps that the job runs.
	Recipe interface{} `field:"optional" json:"recipe" yaml:"recipe"`
	// Metadata tags that have been applied to the job.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The job's timeout in minutes.
	//
	// A job that attempts to run longer than this timeout period ends with a status of `TIMEOUT` .
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// List of validation configurations that are applied to the profile job.
	ValidationConfigurations interface{} `field:"optional" json:"validationConfigurations" yaml:"validationConfigurations"`
}

