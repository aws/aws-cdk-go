package previewawsathenamixins


// The configuration information that will be updated for this workgroup, which includes the location in Amazon S3 where query results are stored, the encryption option, if any, used for query results, whether the Amazon CloudWatch Metrics are enabled for the workgroup, whether the workgroup settings override the client-side settings, and the data usage limit for the amount of bytes scanned per query, if it is specified.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   workGroupConfigurationUpdatesProperty := &WorkGroupConfigurationUpdatesProperty{
//   	AdditionalConfiguration: jsii.String("additionalConfiguration"),
//   	BytesScannedCutoffPerQuery: jsii.Number(123),
//   	CustomerContentEncryptionConfiguration: &CustomerContentEncryptionConfigurationProperty{
//   		KmsKey: jsii.String("kmsKey"),
//   	},
//   	EnforceWorkGroupConfiguration: jsii.Boolean(false),
//   	EngineConfiguration: &EngineConfigurationProperty{
//   		AdditionalConfigs: map[string]*string{
//   			"additionalConfigsKey": jsii.String("additionalConfigs"),
//   		},
//   		Classifications: []interface{}{
//   			&ClassificationProperty{
//   				Name: jsii.String("name"),
//   				Properties: map[string]*string{
//   					"propertiesKey": jsii.String("properties"),
//   				},
//   			},
//   		},
//   		CoordinatorDpuSize: jsii.Number(123),
//   		DefaultExecutorDpuSize: jsii.Number(123),
//   		MaxConcurrentDpus: jsii.Number(123),
//   		SparkProperties: map[string]*string{
//   			"sparkPropertiesKey": jsii.String("sparkProperties"),
//   		},
//   	},
//   	EngineVersion: &EngineVersionProperty{
//   		EffectiveEngineVersion: jsii.String("effectiveEngineVersion"),
//   		SelectedEngineVersion: jsii.String("selectedEngineVersion"),
//   	},
//   	ExecutionRole: jsii.String("executionRole"),
//   	ManagedQueryResultsConfiguration: &ManagedQueryResultsConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		EncryptionConfiguration: &ManagedStorageEncryptionConfigurationProperty{
//   			KmsKey: jsii.String("kmsKey"),
//   		},
//   	},
//   	MonitoringConfiguration: &MonitoringConfigurationProperty{
//   		CloudWatchLoggingConfiguration: &CloudWatchLoggingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroup: jsii.String("logGroup"),
//   			LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   			LogTypes: map[string][]*string{
//   				"logTypesKey": []*string{
//   					jsii.String("logTypes"),
//   				},
//   			},
//   		},
//   		ManagedLoggingConfiguration: &ManagedLoggingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			KmsKey: jsii.String("kmsKey"),
//   		},
//   		S3LoggingConfiguration: &S3LoggingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			KmsKey: jsii.String("kmsKey"),
//   			LogLocation: jsii.String("logLocation"),
//   		},
//   	},
//   	PublishCloudWatchMetricsEnabled: jsii.Boolean(false),
//   	RemoveBytesScannedCutoffPerQuery: jsii.Boolean(false),
//   	RemoveCustomerContentEncryptionConfiguration: jsii.Boolean(false),
//   	RequesterPaysEnabled: jsii.Boolean(false),
//   	ResultConfigurationUpdates: &ResultConfigurationUpdatesProperty{
//   		AclConfiguration: &AclConfigurationProperty{
//   			S3AclOption: jsii.String("s3AclOption"),
//   		},
//   		EncryptionConfiguration: &EncryptionConfigurationProperty{
//   			EncryptionOption: jsii.String("encryptionOption"),
//   			KmsKey: jsii.String("kmsKey"),
//   		},
//   		ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   		OutputLocation: jsii.String("outputLocation"),
//   		RemoveAclConfiguration: jsii.Boolean(false),
//   		RemoveEncryptionConfiguration: jsii.Boolean(false),
//   		RemoveExpectedBucketOwner: jsii.Boolean(false),
//   		RemoveOutputLocation: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html
//
type CfnWorkGroupPropsMixin_WorkGroupConfigurationUpdatesProperty struct {
	// Additional Configuration that are passed to Athena Spark Calculations running in this workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-additionalconfiguration
	//
	AdditionalConfiguration *string `field:"optional" json:"additionalConfiguration" yaml:"additionalConfiguration"`
	// The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-bytesscannedcutoffperquery
	//
	BytesScannedCutoffPerQuery *float64 `field:"optional" json:"bytesScannedCutoffPerQuery" yaml:"bytesScannedCutoffPerQuery"`
	// Indicates the KMS key for encrypting notebook content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-customercontentencryptionconfiguration
	//
	CustomerContentEncryptionConfiguration interface{} `field:"optional" json:"customerContentEncryptionConfiguration" yaml:"customerContentEncryptionConfiguration"`
	// If set to "true", the settings for the workgroup override client-side settings.
	//
	// If set to "false", client-side settings are used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-enforceworkgroupconfiguration
	//
	EnforceWorkGroupConfiguration interface{} `field:"optional" json:"enforceWorkGroupConfiguration" yaml:"enforceWorkGroupConfiguration"`
	// The engine configuration for running queries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-engineconfiguration
	//
	EngineConfiguration interface{} `field:"optional" json:"engineConfiguration" yaml:"engineConfiguration"`
	// The Athena engine version for running queries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-engineversion
	//
	EngineVersion interface{} `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The ARN of the execution role used to access user resources for Spark sessions and Identity Center enabled workgroups.
	//
	// This property applies only to Spark enabled workgroups and Identity Center enabled workgroups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-executionrole
	//
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The configuration for the managed query results and encryption option.
	//
	// ResultConfiguration and ManagedQueryResultsConfiguration cannot be set at the same time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-managedqueryresultsconfiguration
	//
	ManagedQueryResultsConfiguration interface{} `field:"optional" json:"managedQueryResultsConfiguration" yaml:"managedQueryResultsConfiguration"`
	// Contains the configuration settings for managed log persistence, delivering logs to Amazon S3 buckets, Amazon CloudWatch log groups etc.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-monitoringconfiguration
	//
	MonitoringConfiguration interface{} `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
	// Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-publishcloudwatchmetricsenabled
	//
	PublishCloudWatchMetricsEnabled interface{} `field:"optional" json:"publishCloudWatchMetricsEnabled" yaml:"publishCloudWatchMetricsEnabled"`
	// Indicates that the data usage control limit per query is removed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-removebytesscannedcutoffperquery
	//
	RemoveBytesScannedCutoffPerQuery interface{} `field:"optional" json:"removeBytesScannedCutoffPerQuery" yaml:"removeBytesScannedCutoffPerQuery"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-removecustomercontentencryptionconfiguration
	//
	RemoveCustomerContentEncryptionConfiguration interface{} `field:"optional" json:"removeCustomerContentEncryptionConfiguration" yaml:"removeCustomerContentEncryptionConfiguration"`
	// If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries.
	//
	// If set to false, workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-requesterpaysenabled
	//
	RequesterPaysEnabled interface{} `field:"optional" json:"requesterPaysEnabled" yaml:"requesterPaysEnabled"`
	// The result configuration information about the queries in this workgroup that will be updated.
	//
	// Includes the updated results location and an updated option for encrypting query results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfigurationupdates.html#cfn-athena-workgroup-workgroupconfigurationupdates-resultconfigurationupdates
	//
	ResultConfigurationUpdates interface{} `field:"optional" json:"resultConfigurationUpdates" yaml:"resultConfigurationUpdates"`
}

