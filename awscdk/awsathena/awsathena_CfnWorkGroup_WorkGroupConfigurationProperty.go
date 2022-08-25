package awsathena


// The configuration of the workgroup, which includes the location in Amazon S3 where query results are stored, the encryption option, if any, used for query results, whether Amazon CloudWatch Metrics are enabled for the workgroup, and the limit for the amount of bytes scanned (cutoff) per query, if it is specified.
//
// The `EnforceWorkGroupConfiguration` option determines whether workgroup settings override client-side query settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workGroupConfigurationProperty := &workGroupConfigurationProperty{
//   	bytesScannedCutoffPerQuery: jsii.Number(123),
//   	enforceWorkGroupConfiguration: jsii.Boolean(false),
//   	engineVersion: &engineVersionProperty{
//   		effectiveEngineVersion: jsii.String("effectiveEngineVersion"),
//   		selectedEngineVersion: jsii.String("selectedEngineVersion"),
//   	},
//   	publishCloudWatchMetricsEnabled: jsii.Boolean(false),
//   	requesterPaysEnabled: jsii.Boolean(false),
//   	resultConfiguration: &resultConfigurationProperty{
//   		encryptionConfiguration: &encryptionConfigurationProperty{
//   			encryptionOption: jsii.String("encryptionOption"),
//
//   			// the properties below are optional
//   			kmsKey: jsii.String("kmsKey"),
//   		},
//   		outputLocation: jsii.String("outputLocation"),
//   	},
//   }
//
type CfnWorkGroup_WorkGroupConfigurationProperty struct {
	// The upper limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan.
	//
	// No default is defined.
	//
	// > This property currently supports integer types. Support for long values is planned.
	BytesScannedCutoffPerQuery *float64 `field:"optional" json:"bytesScannedCutoffPerQuery" yaml:"bytesScannedCutoffPerQuery"`
	// If set to "true", the settings for the workgroup override client-side settings.
	//
	// If set to "false", client-side settings are used. For more information, see [Workgroup Settings Override Client-Side Settings](https://docs.aws.amazon.com/athena/latest/ug/workgroups-settings-override.html) .
	EnforceWorkGroupConfiguration interface{} `field:"optional" json:"enforceWorkGroupConfiguration" yaml:"enforceWorkGroupConfiguration"`
	// The engine version that all queries running on the workgroup use.
	//
	// Queries on the `AmazonAthenaPreviewFunctionality` workgroup run on the preview engine regardless of this setting.
	EngineVersion interface{} `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.
	PublishCloudWatchMetricsEnabled interface{} `field:"optional" json:"publishCloudWatchMetricsEnabled" yaml:"publishCloudWatchMetricsEnabled"`
	// If set to `true` , allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries.
	//
	// If set to `false` , workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error. The default is `false` . For more information about Requester Pays buckets, see [Requester Pays Buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html) in the *Amazon Simple Storage Service Developer Guide* .
	RequesterPaysEnabled interface{} `field:"optional" json:"requesterPaysEnabled" yaml:"requesterPaysEnabled"`
	// Specifies the location in Amazon S3 where query results are stored and the encryption option, if any, used for query results.
	//
	// For more information, see [Working with Query Results, Output Files, and Query History](https://docs.aws.amazon.com/athena/latest/ug/querying.html) .
	ResultConfiguration interface{} `field:"optional" json:"resultConfiguration" yaml:"resultConfiguration"`
}

