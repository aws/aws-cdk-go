package awsathena

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnWorkGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkGroupProps := &cfnWorkGroupProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	recursiveDeleteOption: jsii.Boolean(false),
//   	state: jsii.String("state"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	workGroupConfiguration: &workGroupConfigurationProperty{
//   		bytesScannedCutoffPerQuery: jsii.Number(123),
//   		enforceWorkGroupConfiguration: jsii.Boolean(false),
//   		engineVersion: &engineVersionProperty{
//   			effectiveEngineVersion: jsii.String("effectiveEngineVersion"),
//   			selectedEngineVersion: jsii.String("selectedEngineVersion"),
//   		},
//   		publishCloudWatchMetricsEnabled: jsii.Boolean(false),
//   		requesterPaysEnabled: jsii.Boolean(false),
//   		resultConfiguration: &resultConfigurationProperty{
//   			encryptionConfiguration: &encryptionConfigurationProperty{
//   				encryptionOption: jsii.String("encryptionOption"),
//
//   				// the properties below are optional
//   				kmsKey: jsii.String("kmsKey"),
//   			},
//   			outputLocation: jsii.String("outputLocation"),
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

