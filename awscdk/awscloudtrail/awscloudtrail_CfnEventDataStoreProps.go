package awscloudtrail

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnEventDataStore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventDataStoreProps := &cfnEventDataStoreProps{
//   	advancedEventSelectors: []interface{}{
//   		&advancedEventSelectorProperty{
//   			fieldSelectors: []interface{}{
//   				&advancedFieldSelectorProperty{
//   					field: jsii.String("field"),
//
//   					// the properties below are optional
//   					endsWith: []*string{
//   						jsii.String("endsWith"),
//   					},
//   					equalTo: []*string{
//   						jsii.String("equalTo"),
//   					},
//   					notEndsWith: []*string{
//   						jsii.String("notEndsWith"),
//   					},
//   					notEquals: []*string{
//   						jsii.String("notEquals"),
//   					},
//   					notStartsWith: []*string{
//   						jsii.String("notStartsWith"),
//   					},
//   					startsWith: []*string{
//   						jsii.String("startsWith"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			name: jsii.String("name"),
//   		},
//   	},
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	multiRegionEnabled: jsii.Boolean(false),
//   	name: jsii.String("name"),
//   	organizationEnabled: jsii.Boolean(false),
//   	retentionPeriod: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	terminationProtectionEnabled: jsii.Boolean(false),
//   }
//
type CfnEventDataStoreProps struct {
	// The advanced event selectors to use to select the events for the data store.
	//
	// You can configure up to five advanced event selectors for each event data store.
	//
	// For more information about how to use advanced event selectors to log CloudTrail events, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.
	//
	// For more information about how to use advanced event selectors to include AWS Config configuration items in your event data store, see [Create an event data store for AWS Config configuration items](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/lake-cli-create-eds-config.html) in the CloudTrail User Guide.
	AdvancedEventSelectors interface{} `field:"optional" json:"advancedEventSelectors" yaml:"advancedEventSelectors"`
	// Specifies the AWS KMS key ID to use to encrypt the events delivered by CloudTrail.
	//
	// The value can be an alias name prefixed by `alias/` , a fully specified ARN to an alias, a fully specified ARN to a key, or a globally unique identifier.
	//
	// > Disabling or deleting the KMS key, or removing CloudTrail permissions on the key, prevents CloudTrail from logging events to the event data store, and prevents users from querying the data in the event data store that was encrypted with the key. After you associate an event data store with a KMS key, the KMS key cannot be removed or changed. Before you disable or delete a KMS key that you are using with an event data store, delete or back up your event data store.
	//
	// CloudTrail also supports AWS KMS multi-Region keys. For more information about multi-Region keys, see [Using multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the *AWS Key Management Service Developer Guide* .
	//
	// Examples:
	//
	// - `alias/MyAliasName`
	// - `arn:aws:kms:us-east-2:123456789012:alias/MyAliasName`
	// - `arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012`
	// - `12345678-1234-1234-1234-123456789012`.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies whether the event data store includes events from all regions, or only from the region in which the event data store is created.
	MultiRegionEnabled interface{} `field:"optional" json:"multiRegionEnabled" yaml:"multiRegionEnabled"`
	// The name of the event data store.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies whether an event data store collects events logged for an organization in AWS Organizations .
	OrganizationEnabled interface{} `field:"optional" json:"organizationEnabled" yaml:"organizationEnabled"`
	// The retention period of the event data store, in days.
	//
	// You can set a retention period of up to 2557 days, the equivalent of seven years.
	RetentionPeriod *float64 `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// A list of tags.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies whether termination protection is enabled for the event data store.
	//
	// If termination protection is enabled, you cannot delete the event data store until termination protection is disabled.
	TerminationProtectionEnabled interface{} `field:"optional" json:"terminationProtectionEnabled" yaml:"terminationProtectionEnabled"`
}

