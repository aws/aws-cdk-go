package awscloudtrail

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEventDataStore`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventDataStoreProps := &CfnEventDataStoreProps{
//   	AdvancedEventSelectors: []interface{}{
//   		&AdvancedEventSelectorProperty{
//   			FieldSelectors: []interface{}{
//   				&AdvancedFieldSelectorProperty{
//   					Field: jsii.String("field"),
//
//   					// the properties below are optional
//   					EndsWith: []*string{
//   						jsii.String("endsWith"),
//   					},
//   					EqualTo: []*string{
//   						jsii.String("equalTo"),
//   					},
//   					NotEndsWith: []*string{
//   						jsii.String("notEndsWith"),
//   					},
//   					NotEquals: []*string{
//   						jsii.String("notEquals"),
//   					},
//   					NotStartsWith: []*string{
//   						jsii.String("notStartsWith"),
//   					},
//   					StartsWith: []*string{
//   						jsii.String("startsWith"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	BillingMode: jsii.String("billingMode"),
//   	ContextKeySelectors: []interface{}{
//   		&ContextKeySelectorProperty{
//   			EqualTo: []*string{
//   				jsii.String("equalTo"),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	FederationEnabled: jsii.Boolean(false),
//   	FederationRoleArn: jsii.String("federationRoleArn"),
//   	IngestionEnabled: jsii.Boolean(false),
//   	InsightsDestination: jsii.String("insightsDestination"),
//   	InsightSelectors: []interface{}{
//   		&InsightSelectorProperty{
//   			InsightType: jsii.String("insightType"),
//   		},
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	MaxEventSize: jsii.String("maxEventSize"),
//   	MultiRegionEnabled: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	OrganizationEnabled: jsii.Boolean(false),
//   	RetentionPeriod: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TerminationProtectionEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-eventdatastore.html
//
type CfnEventDataStoreProps struct {
	// The advanced event selectors to use to select the events for the data store.
	//
	// You can configure up to five advanced event selectors for each event data store.
	//
	// For more information about how to use advanced event selectors to log CloudTrail events, see [Log events by using advanced event selectors](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#creating-data-event-selectors-advanced) in the CloudTrail User Guide.
	//
	// For more information about how to use advanced event selectors to include AWS Config configuration items in your event data store, see [Create an event data store for AWS Config configuration items](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/lake-eds-cli.html#lake-cli-create-eds-config) in the CloudTrail User Guide.
	//
	// For more information about how to use advanced event selectors to include events outside of AWS events in your event data store, see [Create an integration to log events from outside AWS](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/lake-integrations-cli.html#lake-cli-create-integration) in the CloudTrail User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-eventdatastore.html#cfn-cloudtrail-eventdatastore-advancedeventselectors
	//
	AdvancedEventSelectors interface{} `field:"optional" json:"advancedEventSelectors" yaml:"advancedEventSelectors"`
	// The billing mode for the event data store determines the cost for ingesting events and the default and maximum retention period for the event data store.
	//
	// The following are the possible values:
	//
	// - `EXTENDABLE_RETENTION_PRICING` - This billing mode is generally recommended if you want a flexible retention period of up to 3653 days (about 10 years). The default retention period for this billing mode is 366 days.
	// - `FIXED_RETENTION_PRICING` - This billing mode is recommended if you expect to ingest more than 25 TB of event data per month and need a retention period of up to 2557 days (about 7 years). The default retention period for this billing mode is 2557 days.
	//
	// The default value is `EXTENDABLE_RETENTION_PRICING` .
	//
	// For more information about CloudTrail pricing, see [AWS CloudTrail Pricing](https://docs.aws.amazon.com/cloudtrail/pricing/) and [Managing CloudTrail Lake costs](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-lake-manage-costs.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-eventdatastore.html#cfn-cloudtrail-eventdatastore-billingmode
	//
	BillingMode *string `field:"optional" json:"billingMode" yaml:"billingMode"`
	// The list of context key selectors that are configured for the event data store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-eventdatastore.html#cfn-cloudtrail-eventdatastore-contextkeyselectors
	//
	ContextKeySelectors interface{} `field:"optional" json:"contextKeySelectors" yaml:"contextKeySelectors"`
	// Indicates if [Lake query federation](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-federation.html) is enabled. By default, Lake query federation is disabled. You cannot delete an event data store if Lake query federation is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-eventdatastore.html#cfn-cloudtrail-eventdatastore-federationenabled
	//
	FederationEnabled interface{} `field:"optional" json:"federationEnabled" yaml:"federationEnabled"`
	// If Lake query federation is enabled, provides the ARN of the federation role used to access the resources for the federated event data store.
	//
	// The federation role must exist in your account and provide the [required minimum permissions](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-federation.html#query-federation-permissions-role) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-eventdatastore.html#cfn-cloudtrail-eventdatastore-federationrolearn
	//
	FederationRoleArn *string `field:"optional" json:"federationRoleArn" yaml:"federationRoleArn"`
	// Specifies whether the event data store should start ingesting live events.
	//
	// The default is true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-eventdatastore.html#cfn-cloudtrail-eventdatastore-ingestionenabled
	//
	IngestionEnabled interface{} `field:"optional" json:"ingestionEnabled" yaml:"ingestionEnabled"`
	// The ARN (or ID suffix of the ARN) of the destination event data store that logs Insights events.
	//
	// For more information, see [Create an event data store for CloudTrail Insights events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-event-data-store-insights.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-eventdatastore.html#cfn-cloudtrail-eventdatastore-insightsdestination
	//
	InsightsDestination *string `field:"optional" json:"insightsDestination" yaml:"insightsDestination"`
	// A JSON string that contains the Insights types you want to log on an event data store.
	//
	// `ApiCallRateInsight` and `ApiErrorRateInsight` are valid Insight types.
	//
	// The `ApiCallRateInsight` Insights type analyzes write-only management API calls that are aggregated per minute against a baseline API call volume.
	//
	// The `ApiErrorRateInsight` Insights type analyzes management API calls that result in error codes. The error is shown if the API call is unsuccessful.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-eventdatastore.html#cfn-cloudtrail-eventdatastore-insightselectors
	//
	InsightSelectors interface{} `field:"optional" json:"insightSelectors" yaml:"insightSelectors"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-eventdatastore.html#cfn-cloudtrail-eventdatastore-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The maximum allowed size for events to be stored in the specified event data store.
	//
	// If you are using context key selectors, MaxEventSize must be set to Large.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-eventdatastore.html#cfn-cloudtrail-eventdatastore-maxeventsize
	//
	MaxEventSize *string `field:"optional" json:"maxEventSize" yaml:"maxEventSize"`
	// Specifies whether the event data store includes events from all Regions, or only from the Region in which the event data store is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-eventdatastore.html#cfn-cloudtrail-eventdatastore-multiregionenabled
	//
	MultiRegionEnabled interface{} `field:"optional" json:"multiRegionEnabled" yaml:"multiRegionEnabled"`
	// The name of the event data store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-eventdatastore.html#cfn-cloudtrail-eventdatastore-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies whether an event data store collects events logged for an organization in AWS Organizations .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-eventdatastore.html#cfn-cloudtrail-eventdatastore-organizationenabled
	//
	OrganizationEnabled interface{} `field:"optional" json:"organizationEnabled" yaml:"organizationEnabled"`
	// The retention period of the event data store, in days.
	//
	// If `BillingMode` is set to `EXTENDABLE_RETENTION_PRICING` , you can set a retention period of up to 3653 days, the equivalent of 10 years. If `BillingMode` is set to `FIXED_RETENTION_PRICING` , you can set a retention period of up to 2557 days, the equivalent of seven years.
	//
	// CloudTrail Lake determines whether to retain an event by checking if the `eventTime` of the event is within the specified retention period. For example, if you set a retention period of 90 days, CloudTrail will remove events when the `eventTime` is older than 90 days.
	//
	// > If you plan to copy trail events to this event data store, we recommend that you consider both the age of the events that you want to copy as well as how long you want to keep the copied events in your event data store. For example, if you copy trail events that are 5 years old and specify a retention period of 7 years, the event data store will retain those events for two years.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-eventdatastore.html#cfn-cloudtrail-eventdatastore-retentionperiod
	//
	RetentionPeriod *float64 `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// A list of tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-eventdatastore.html#cfn-cloudtrail-eventdatastore-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Specifies whether termination protection is enabled for the event data store.
	//
	// If termination protection is enabled, you cannot delete the event data store until termination protection is disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudtrail-eventdatastore.html#cfn-cloudtrail-eventdatastore-terminationprotectionenabled
	//
	TerminationProtectionEnabled interface{} `field:"optional" json:"terminationProtectionEnabled" yaml:"terminationProtectionEnabled"`
}

