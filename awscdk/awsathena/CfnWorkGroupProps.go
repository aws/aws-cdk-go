package awsathena

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnWorkGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkGroupProps := &CfnWorkGroupProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	RecursiveDeleteOption: jsii.Boolean(false),
//   	State: jsii.String("state"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkGroupConfiguration: &WorkGroupConfigurationProperty{
//   		BytesScannedCutoffPerQuery: jsii.Number(123),
//   		EnforceWorkGroupConfiguration: jsii.Boolean(false),
//   		EngineVersion: &EngineVersionProperty{
//   			EffectiveEngineVersion: jsii.String("effectiveEngineVersion"),
//   			SelectedEngineVersion: jsii.String("selectedEngineVersion"),
//   		},
//   		PublishCloudWatchMetricsEnabled: jsii.Boolean(false),
//   		RequesterPaysEnabled: jsii.Boolean(false),
//   		ResultConfiguration: &ResultConfigurationProperty{
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				EncryptionOption: jsii.String("encryptionOption"),
//
//   				// the properties below are optional
//   				KmsKey: jsii.String("kmsKey"),
//   			},
//   			OutputLocation: jsii.String("outputLocation"),
//   		},
//   	},
//   }
//
type CfnWorkGroupProps struct {
	// The workgroup name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The workgroup description.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The option to delete a workgroup and its contents even if the workgroup contains any named queries.
	//
	// The default is false.
	RecursiveDeleteOption interface{} `field:"optional" json:"recursiveDeleteOption" yaml:"recursiveDeleteOption"`
	// The state of the workgroup: ENABLED or DISABLED.
	State *string `field:"optional" json:"state" yaml:"state"`
	// The tags (key-value pairs) to associate with this resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The configuration of the workgroup, which includes the location in Amazon S3 where query results are stored, the encryption option, if any, used for query results, whether Amazon CloudWatch Metrics are enabled for the workgroup, and the limit for the amount of bytes scanned (cutoff) per query, if it is specified.
	//
	// The `EnforceWorkGroupConfiguration` option determines whether workgroup settings override client-side query settings.
	WorkGroupConfiguration interface{} `field:"optional" json:"workGroupConfiguration" yaml:"workGroupConfiguration"`
}

