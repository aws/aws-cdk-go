package previewawsathenamixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnWorkGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWorkGroupMixinProps := &CfnWorkGroupMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	RecursiveDeleteOption: jsii.Boolean(false),
//   	State: jsii.String("state"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkGroupConfiguration: &WorkGroupConfigurationProperty{
//   		AdditionalConfiguration: jsii.String("additionalConfiguration"),
//   		BytesScannedCutoffPerQuery: jsii.Number(123),
//   		CustomerContentEncryptionConfiguration: &CustomerContentEncryptionConfigurationProperty{
//   			KmsKey: jsii.String("kmsKey"),
//   		},
//   		EnforceWorkGroupConfiguration: jsii.Boolean(false),
//   		EngineConfiguration: &EngineConfigurationProperty{
//   			AdditionalConfigs: map[string]*string{
//   				"additionalConfigsKey": jsii.String("additionalConfigs"),
//   			},
//   			Classifications: []interface{}{
//   				&ClassificationProperty{
//   					Name: jsii.String("name"),
//   					Properties: map[string]*string{
//   						"propertiesKey": jsii.String("properties"),
//   					},
//   				},
//   			},
//   			CoordinatorDpuSize: jsii.Number(123),
//   			DefaultExecutorDpuSize: jsii.Number(123),
//   			MaxConcurrentDpus: jsii.Number(123),
//   			SparkProperties: map[string]*string{
//   				"sparkPropertiesKey": jsii.String("sparkProperties"),
//   			},
//   		},
//   		EngineVersion: &EngineVersionProperty{
//   			EffectiveEngineVersion: jsii.String("effectiveEngineVersion"),
//   			SelectedEngineVersion: jsii.String("selectedEngineVersion"),
//   		},
//   		ExecutionRole: jsii.String("executionRole"),
//   		ManagedQueryResultsConfiguration: &ManagedQueryResultsConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			EncryptionConfiguration: &ManagedStorageEncryptionConfigurationProperty{
//   				KmsKey: jsii.String("kmsKey"),
//   			},
//   		},
//   		MonitoringConfiguration: &MonitoringConfigurationProperty{
//   			CloudWatchLoggingConfiguration: &CloudWatchLoggingConfigurationProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroup: jsii.String("logGroup"),
//   				LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   				LogTypes: map[string][]*string{
//   					"logTypesKey": []*string{
//   						jsii.String("logTypes"),
//   					},
//   				},
//   			},
//   			ManagedLoggingConfiguration: &ManagedLoggingConfigurationProperty{
//   				Enabled: jsii.Boolean(false),
//   				KmsKey: jsii.String("kmsKey"),
//   			},
//   			S3LoggingConfiguration: &S3LoggingConfigurationProperty{
//   				Enabled: jsii.Boolean(false),
//   				KmsKey: jsii.String("kmsKey"),
//   				LogLocation: jsii.String("logLocation"),
//   			},
//   		},
//   		PublishCloudWatchMetricsEnabled: jsii.Boolean(false),
//   		RequesterPaysEnabled: jsii.Boolean(false),
//   		ResultConfiguration: &ResultConfigurationProperty{
//   			AclConfiguration: &AclConfigurationProperty{
//   				S3AclOption: jsii.String("s3AclOption"),
//   			},
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				EncryptionOption: jsii.String("encryptionOption"),
//   				KmsKey: jsii.String("kmsKey"),
//   			},
//   			ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   			OutputLocation: jsii.String("outputLocation"),
//   		},
//   	},
//   	WorkGroupConfigurationUpdates: &WorkGroupConfigurationUpdatesProperty{
//   		AdditionalConfiguration: jsii.String("additionalConfiguration"),
//   		BytesScannedCutoffPerQuery: jsii.Number(123),
//   		CustomerContentEncryptionConfiguration: &CustomerContentEncryptionConfigurationProperty{
//   			KmsKey: jsii.String("kmsKey"),
//   		},
//   		EnforceWorkGroupConfiguration: jsii.Boolean(false),
//   		EngineConfiguration: &EngineConfigurationProperty{
//   			AdditionalConfigs: map[string]*string{
//   				"additionalConfigsKey": jsii.String("additionalConfigs"),
//   			},
//   			Classifications: []interface{}{
//   				&ClassificationProperty{
//   					Name: jsii.String("name"),
//   					Properties: map[string]*string{
//   						"propertiesKey": jsii.String("properties"),
//   					},
//   				},
//   			},
//   			CoordinatorDpuSize: jsii.Number(123),
//   			DefaultExecutorDpuSize: jsii.Number(123),
//   			MaxConcurrentDpus: jsii.Number(123),
//   			SparkProperties: map[string]*string{
//   				"sparkPropertiesKey": jsii.String("sparkProperties"),
//   			},
//   		},
//   		EngineVersion: &EngineVersionProperty{
//   			EffectiveEngineVersion: jsii.String("effectiveEngineVersion"),
//   			SelectedEngineVersion: jsii.String("selectedEngineVersion"),
//   		},
//   		ExecutionRole: jsii.String("executionRole"),
//   		ManagedQueryResultsConfiguration: &ManagedQueryResultsConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			EncryptionConfiguration: &ManagedStorageEncryptionConfigurationProperty{
//   				KmsKey: jsii.String("kmsKey"),
//   			},
//   		},
//   		MonitoringConfiguration: &MonitoringConfigurationProperty{
//   			CloudWatchLoggingConfiguration: &CloudWatchLoggingConfigurationProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroup: jsii.String("logGroup"),
//   				LogStreamNamePrefix: jsii.String("logStreamNamePrefix"),
//   				LogTypes: map[string][]*string{
//   					"logTypesKey": []*string{
//   						jsii.String("logTypes"),
//   					},
//   				},
//   			},
//   			ManagedLoggingConfiguration: &ManagedLoggingConfigurationProperty{
//   				Enabled: jsii.Boolean(false),
//   				KmsKey: jsii.String("kmsKey"),
//   			},
//   			S3LoggingConfiguration: &S3LoggingConfigurationProperty{
//   				Enabled: jsii.Boolean(false),
//   				KmsKey: jsii.String("kmsKey"),
//   				LogLocation: jsii.String("logLocation"),
//   			},
//   		},
//   		PublishCloudWatchMetricsEnabled: jsii.Boolean(false),
//   		RemoveBytesScannedCutoffPerQuery: jsii.Boolean(false),
//   		RemoveCustomerContentEncryptionConfiguration: jsii.Boolean(false),
//   		RequesterPaysEnabled: jsii.Boolean(false),
//   		ResultConfigurationUpdates: &ResultConfigurationUpdatesProperty{
//   			AclConfiguration: &AclConfigurationProperty{
//   				S3AclOption: jsii.String("s3AclOption"),
//   			},
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				EncryptionOption: jsii.String("encryptionOption"),
//   				KmsKey: jsii.String("kmsKey"),
//   			},
//   			ExpectedBucketOwner: jsii.String("expectedBucketOwner"),
//   			OutputLocation: jsii.String("outputLocation"),
//   			RemoveAclConfiguration: jsii.Boolean(false),
//   			RemoveEncryptionConfiguration: jsii.Boolean(false),
//   			RemoveExpectedBucketOwner: jsii.Boolean(false),
//   			RemoveOutputLocation: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html
//
type CfnWorkGroupMixinProps struct {
	// The workgroup description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The workgroup name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The option to delete a workgroup and its contents even if the workgroup contains any named queries.
	//
	// The default is false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-recursivedeleteoption
	//
	RecursiveDeleteOption interface{} `field:"optional" json:"recursiveDeleteOption" yaml:"recursiveDeleteOption"`
	// The state of the workgroup: ENABLED or DISABLED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// The tags (key-value pairs) to associate with this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The configuration of the workgroup, which includes the location in Amazon S3 where query results are stored, the encryption option, if any, used for query results, whether Amazon CloudWatch Metrics are enabled for the workgroup, and the limit for the amount of bytes scanned (cutoff) per query, if it is specified.
	//
	// The `EnforceWorkGroupConfiguration` option determines whether workgroup settings override client-side query settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-workgroupconfiguration
	//
	WorkGroupConfiguration interface{} `field:"optional" json:"workGroupConfiguration" yaml:"workGroupConfiguration"`
	// The configuration information that will be updated for this workgroup, which includes the location in Amazon S3 where query results are stored, the encryption option, if any, used for query results, whether the Amazon CloudWatch Metrics are enabled for the workgroup, whether the workgroup settings override the client-side settings, and the data usage limit for the amount of bytes scanned per query, if it is specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html#cfn-athena-workgroup-workgroupconfigurationupdates
	//
	// Deprecated: this property has been deprecated.
	WorkGroupConfigurationUpdates interface{} `field:"optional" json:"workGroupConfigurationUpdates" yaml:"workGroupConfigurationUpdates"`
}

