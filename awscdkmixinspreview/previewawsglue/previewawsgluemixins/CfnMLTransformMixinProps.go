package previewawsgluemixins


// Properties for CfnMLTransformPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var tags interface{}
//
//   cfnMLTransformMixinProps := &CfnMLTransformMixinProps{
//   	Description: jsii.String("description"),
//   	GlueVersion: jsii.String("glueVersion"),
//   	InputRecordTables: &InputRecordTablesProperty{
//   		GlueTables: []interface{}{
//   			&GlueTablesProperty{
//   				CatalogId: jsii.String("catalogId"),
//   				ConnectionName: jsii.String("connectionName"),
//   				DatabaseName: jsii.String("databaseName"),
//   				TableName: jsii.String("tableName"),
//   			},
//   		},
//   	},
//   	MaxCapacity: jsii.Number(123),
//   	MaxRetries: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	NumberOfWorkers: jsii.Number(123),
//   	Role: jsii.String("role"),
//   	Tags: tags,
//   	Timeout: jsii.Number(123),
//   	TransformEncryption: &TransformEncryptionProperty{
//   		MlUserDataEncryption: &MLUserDataEncryptionProperty{
//   			KmsKeyId: jsii.String("kmsKeyId"),
//   			MlUserDataEncryptionMode: jsii.String("mlUserDataEncryptionMode"),
//   		},
//   		TaskRunSecurityConfigurationName: jsii.String("taskRunSecurityConfigurationName"),
//   	},
//   	TransformParameters: &TransformParametersProperty{
//   		FindMatchesParameters: &FindMatchesParametersProperty{
//   			AccuracyCostTradeoff: jsii.Number(123),
//   			EnforceProvidedLabels: jsii.Boolean(false),
//   			PrecisionRecallTradeoff: jsii.Number(123),
//   			PrimaryKeyColumnName: jsii.String("primaryKeyColumnName"),
//   		},
//   		TransformType: jsii.String("transformType"),
//   	},
//   	WorkerType: jsii.String("workerType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html
//
type CfnMLTransformMixinProps struct {
	// A user-defined, long-form description text for the machine learning transform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// This value determines which version of AWS Glue this machine learning transform is compatible with.
	//
	// Glue 1.0 is recommended for most customers. If the value is not set, the Glue compatibility defaults to Glue 0.9. For more information, see [AWS Glue Versions](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html#release-notes-versions) in the developer guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-glueversion
	//
	GlueVersion *string `field:"optional" json:"glueVersion" yaml:"glueVersion"`
	// A list of AWS Glue table definitions used by the transform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-inputrecordtables
	//
	InputRecordTables interface{} `field:"optional" json:"inputRecordTables" yaml:"inputRecordTables"`
	// The number of AWS Glue data processing units (DPUs) that are allocated to task runs for this transform.
	//
	// You can allocate from 2 to 100 DPUs; the default is 10. A DPU is a relative measure of processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory. For more information, see the [AWS Glue pricing page](https://docs.aws.amazon.com/glue/pricing/) .
	//
	// `MaxCapacity` is a mutually exclusive option with `NumberOfWorkers` and `WorkerType` .
	//
	// - If either `NumberOfWorkers` or `WorkerType` is set, then `MaxCapacity` cannot be set.
	// - If `MaxCapacity` is set then neither `NumberOfWorkers` or `WorkerType` can be set.
	// - If `WorkerType` is set, then `NumberOfWorkers` is required (and vice versa).
	// - `MaxCapacity` and `NumberOfWorkers` must both be at least 1.
	//
	// When the `WorkerType` field is set to a value other than `Standard` , the `MaxCapacity` field is set automatically and becomes read-only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-maxcapacity
	//
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The maximum number of times to retry after an `MLTaskRun` of the machine learning transform fails.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-maxretries
	//
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// A user-defined name for the machine learning transform. Names are required to be unique. `Name` is optional:.
	//
	// - If you supply `Name` , the stack cannot be repeatedly created.
	// - If `Name` is not provided, a randomly generated name will be used instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The number of workers of a defined `workerType` that are allocated when a task of the transform runs.
	//
	// If `WorkerType` is set, then `NumberOfWorkers` is required (and vice versa).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-numberofworkers
	//
	NumberOfWorkers *float64 `field:"optional" json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// The name or Amazon Resource Name (ARN) of the IAM role with the required permissions.
	//
	// The required permissions include both AWS Glue service role permissions to AWS Glue resources, and Amazon S3 permissions required by the transform.
	//
	// - This role needs AWS Glue service role permissions to allow access to resources in AWS Glue . See [Attach a Policy to IAM Users That Access AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/attach-policy-iam-user.html) .
	// - This role needs permission to your Amazon Simple Storage Service (Amazon S3) sources, targets, temporary directory, scripts, and any libraries used by the task run for this transform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-role
	//
	Role *string `field:"optional" json:"role" yaml:"role"`
	// The tags to use with this machine learning transform.
	//
	// You may use tags to limit access to the machine learning transform. For more information about tags in AWS Glue , see [AWS Tags in AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/monitor-tags.html) in the developer guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The timeout in minutes of the machine learning transform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-timeout
	//
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// The encryption-at-rest settings of the transform that apply to accessing user data.
	//
	// Machine learning
	// transforms can access user data encrypted in Amazon S3 using KMS.
	//
	// Additionally, imported labels and trained transforms can now be encrypted using a customer provided
	// KMS key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-transformencryption
	//
	TransformEncryption interface{} `field:"optional" json:"transformEncryption" yaml:"transformEncryption"`
	// The algorithm-specific parameters that are associated with the machine learning transform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-transformparameters
	//
	TransformParameters interface{} `field:"optional" json:"transformParameters" yaml:"transformParameters"`
	// The type of predefined worker that is allocated when a task of this transform runs.
	//
	// Accepts a value of Standard, G.1X, or G.2X.
	//
	// - For the `Standard` worker type, each worker provides 4 vCPU, 16 GB of memory and a 50GB disk, and 2 executors per worker.
	// - For the `G.1X` worker type, each worker provides 4 vCPU, 16 GB of memory and a 64GB disk, and 1 executor per worker.
	// - For the `G.2X` worker type, each worker provides 8 vCPU, 32 GB of memory and a 128GB disk, and 1 executor per worker.
	//
	// `MaxCapacity` is a mutually exclusive option with `NumberOfWorkers` and `WorkerType` .
	//
	// - If either `NumberOfWorkers` or `WorkerType` is set, then `MaxCapacity` cannot be set.
	// - If `MaxCapacity` is set then neither `NumberOfWorkers` or `WorkerType` can be set.
	// - If `WorkerType` is set, then `NumberOfWorkers` is required (and vice versa).
	// - `MaxCapacity` and `NumberOfWorkers` must both be at least 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-mltransform.html#cfn-glue-mltransform-workertype
	//
	WorkerType *string `field:"optional" json:"workerType" yaml:"workerType"`
}

