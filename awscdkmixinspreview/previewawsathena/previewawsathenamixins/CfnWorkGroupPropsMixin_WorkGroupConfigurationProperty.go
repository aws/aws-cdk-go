package previewawsathenamixins


// The configuration of the workgroup, which includes the location in Amazon S3 where query results are stored, the encryption option, if any, used for query results, whether Amazon CloudWatch Metrics are enabled for the workgroup, and the limit for the amount of bytes scanned (cutoff) per query, if it is specified.
//
// The `EnforceWorkGroupConfiguration` option determines whether workgroup settings override client-side query settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   workGroupConfigurationProperty := &WorkGroupConfigurationProperty{
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
//   	RequesterPaysEnabled: jsii.Boolean(false),
//   	ResultConfiguration: &ResultConfigurationProperty{
//   		AclConfiguration: &AclConfigurationProperty{
//   			S3AclOption: jsii.String("s3AclOption"),
//   		},
//   		EncryptionConfiguration: &EncryptionConfigurationProperty{
//   			EncryptionOption: jsii.String("encryptionOption"),
//   			KmsKey: jsii.String("kmsKey"),
//   		},
//   		ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   		OutputLocation: jsii.String("outputLocation"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html
//
type CfnWorkGroupPropsMixin_WorkGroupConfigurationProperty struct {
	// Specifies a user defined JSON string that is passed to the session engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-additionalconfiguration
	//
	AdditionalConfiguration *string `field:"optional" json:"additionalConfiguration" yaml:"additionalConfiguration"`
	// The upper limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan.
	//
	// No default is defined.
	//
	// > This property currently supports integer types. Support for long values is planned.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-bytesscannedcutoffperquery
	//
	BytesScannedCutoffPerQuery *float64 `field:"optional" json:"bytesScannedCutoffPerQuery" yaml:"bytesScannedCutoffPerQuery"`
	// Specifies the KMS key that is used to encrypt the user's data stores in Athena.
	//
	// This setting does not apply to Athena SQL workgroups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-customercontentencryptionconfiguration
	//
	CustomerContentEncryptionConfiguration interface{} `field:"optional" json:"customerContentEncryptionConfiguration" yaml:"customerContentEncryptionConfiguration"`
	// If set to "true", the settings for the workgroup override client-side settings.
	//
	// If set to "false", client-side settings are used. For more information, see [Override client-side settings](https://docs.aws.amazon.com/athena/latest/ug/workgroups-settings-override.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-enforceworkgroupconfiguration
	//
	EnforceWorkGroupConfiguration interface{} `field:"optional" json:"enforceWorkGroupConfiguration" yaml:"enforceWorkGroupConfiguration"`
	// The engine configuration for running queries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-engineconfiguration
	//
	EngineConfiguration interface{} `field:"optional" json:"engineConfiguration" yaml:"engineConfiguration"`
	// The engine version that all queries running on the workgroup use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-engineversion
	//
	EngineVersion interface{} `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Role used to access user resources in an Athena for Apache Spark session.
	//
	// This property applies only to Spark-enabled workgroups in Athena.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-executionrole
	//
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The configuration for storing results in Athena owned storage, which includes whether this feature is enabled;
	//
	// whether encryption configuration, if any, is used for encrypting query results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-managedqueryresultsconfiguration
	//
	ManagedQueryResultsConfiguration interface{} `field:"optional" json:"managedQueryResultsConfiguration" yaml:"managedQueryResultsConfiguration"`
	// Contains the configuration settings for managed log persistence, delivering logs to Amazon S3 buckets, Amazon CloudWatch log groups etc.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-monitoringconfiguration
	//
	MonitoringConfiguration interface{} `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
	// Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-publishcloudwatchmetricsenabled
	//
	PublishCloudWatchMetricsEnabled interface{} `field:"optional" json:"publishCloudWatchMetricsEnabled" yaml:"publishCloudWatchMetricsEnabled"`
	// If set to `true` , allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries.
	//
	// If set to `false` , workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error. The default is `false` . For more information about Requester Pays buckets, see [Requester Pays Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html) in the *Amazon Simple Storage Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-requesterpaysenabled
	//
	RequesterPaysEnabled interface{} `field:"optional" json:"requesterPaysEnabled" yaml:"requesterPaysEnabled"`
	// Specifies the location in Amazon S3 where query results are stored and the encryption option, if any, used for query results.
	//
	// For more information, see [Work with query results and recent queries](https://docs.aws.amazon.com/athena/latest/ug/querying.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-workgroupconfiguration.html#cfn-athena-workgroup-workgroupconfiguration-resultconfiguration
	//
	ResultConfiguration interface{} `field:"optional" json:"resultConfiguration" yaml:"resultConfiguration"`
}

