package awscloudtrail


// A single selector statement in an advanced event selector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   advancedFieldSelectorProperty := &AdvancedFieldSelectorProperty{
//   	Field: jsii.String("field"),
//
//   	// the properties below are optional
//   	EndsWith: []*string{
//   		jsii.String("endsWith"),
//   	},
//   	EqualTo: []*string{
//   		jsii.String("equalTo"),
//   	},
//   	NotEndsWith: []*string{
//   		jsii.String("notEndsWith"),
//   	},
//   	NotEquals: []*string{
//   		jsii.String("notEquals"),
//   	},
//   	NotStartsWith: []*string{
//   		jsii.String("notStartsWith"),
//   	},
//   	StartsWith: []*string{
//   		jsii.String("startsWith"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-advancedfieldselector.html
//
type CfnEventDataStore_AdvancedFieldSelectorProperty struct {
	// A field in a CloudTrail event record on which to filter events to be logged.
	//
	// For event data stores for CloudTrail Insights events, AWS Config configuration items, Audit Manager evidence, or events outside of AWS , the field is used only for selecting events as filtering is not supported.
	//
	// For CloudTrail management events, supported fields include `eventCategory` (required), `eventSource` , and `readOnly` . The following additional fields are available for event data stores: `eventName` , `eventType` , `sessionCredentialFromConsole` , and `userIdentity.arn` .
	//
	// For CloudTrail data events, supported fields include `eventCategory` (required), `resources.type` (required), `eventName` , `readOnly` , and `resources.ARN` . The following additional fields are available for event data stores: `eventSource` , `eventType` , `sessionCredentialFromConsole` , and `userIdentity.arn` .
	//
	// For CloudTrail network activity events, supported fields include `eventCategory` (required), `eventSource` (required), `eventName` , `errorCode` , and `vpcEndpointId` .
	//
	// For event data stores for CloudTrail Insights events, AWS Config configuration items, Audit Manager evidence, or events outside of AWS , the only supported field is `eventCategory` .
	//
	// > Selectors don't support the use of wildcards like `*` . To match multiple values with a single condition, you may use `StartsWith` , `EndsWith` , `NotStartsWith` , or `NotEndsWith` to explicitly match the beginning or end of the event field.
	//
	// - *`readOnly`* - This is an optional field that is only used for management events and data events. This field can be set to `Equals` with a value of `true` or `false` . If you do not add this field, CloudTrail logs both `read` and `write` events. A value of `true` logs only `read` events. A value of `false` logs only `write` events.
	// - *`eventSource`* - This field is only used for management events, data events (for event data stores only), and network activity events.
	//
	// For management events for trails, this is an optional field that can be set to `NotEquals` `kms.amazonaws.com` to exclude KMS management events, or `NotEquals` `rdsdata.amazonaws.com` to exclude RDS management events.
	//
	// For management and data events for event data stores, you can use it to include or exclude any event source and can use any operator.
	//
	// For network activity events, this is a required field that only uses the `Equals` operator. Set this field to the event source for which you want to log network activity events. If you want to log network activity events for multiple event sources, you must create a separate field selector for each event source.
	//
	// The following are valid values for network activity events:
	//
	// - `cloudtrail.amazonaws.com`
	// - `ec2.amazonaws.com`
	// - `kms.amazonaws.com`
	// - `secretsmanager.amazonaws.com`
	// - *`eventName`* - This is an optional field that is only used for data events, management events (for event data stores only), and network activity events. You can use any operator with `eventName` . You can use it to ﬁlter in or ﬁlter out specific events. You can have multiple values for this ﬁeld, separated by commas.
	// - *`eventCategory`* - This field is required and must be set to `Equals` .
	//
	// - For CloudTrail management events, the value must be `Management` .
	// - For CloudTrail data events, the value must be `Data` .
	// - For CloudTrail network activity events, the value must be `NetworkActivity` .
	//
	// The following are used only for event data stores:
	//
	// - For CloudTrail Insights events, the value must be `Insight` .
	// - For AWS Config configuration items, the value must be `ConfigurationItem` .
	// - For Audit Manager evidence, the value must be `Evidence` .
	// - For events outside of AWS , the value must be `ActivityAuditLog` .
	// - *`eventType`* - This is an optional field available only for event data stores, which is used to filter management and data events on the event type. For information about available event types, see [CloudTrail record contents](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-event-reference-record-contents.html#ct-event-type) in the *AWS CloudTrail user guide* .
	// - *`errorCode`* - This ﬁeld is only used to filter CloudTrail network activity events and is optional. This is the error code to filter on. Currently, the only valid `errorCode` is `VpceAccessDenied` . `errorCode` can only use the `Equals` operator.
	// - *`sessionCredentialFromConsole`* - This is an optional field available only for event data stores, which is used to filter management and data events based on whether the events originated from an AWS Management Console session. `sessionCredentialFromConsole` can only use the `Equals` and `NotEquals` operators.
	// - *`resources.type`* - This ﬁeld is required for CloudTrail data events. `resources.type` can only use the `Equals` operator.
	//
	// For a list of available resource types for data events, see [Data events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html#logging-data-events) in the *AWS CloudTrail User Guide* .
	//
	// You can have only one `resources.type` ﬁeld per selector. To log events on more than one resource type, add another selector.
	// - *`resources.ARN`* - The `resources.ARN` is an optional field for data events. You can use any operator with `resources.ARN` , but if you use `Equals` or `NotEquals` , the value must exactly match the ARN of a valid resource of the type you've speciﬁed in the template as the value of resources.type. To log all data events for all objects in a specific S3 bucket, use the `StartsWith` operator, and include only the bucket ARN as the matching value.
	//
	// For information about filtering data events on the `resources.ARN` field, see [Filtering data events by resources.ARN](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/filtering-data-events.html#filtering-data-events-resourcearn) in the *AWS CloudTrail User Guide* .
	//
	// > You can't use the `resources.ARN` field to filter resource types that do not have ARNs.
	// - *`userIdentity.arn`* - This is an optional field available only for event data stores, which is used to filter management and data events on the userIdentity ARN. You can use any operator with `userIdentity.arn` . For more information on the userIdentity element, see [CloudTrail userIdentity element](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-event-reference-user-identity.html) in the *AWS CloudTrail User Guide* .
	// - *`vpcEndpointId`* - This ﬁeld is only used to filter CloudTrail network activity events and is optional. This field identifies the VPC endpoint that the request passed through. You can use any operator with `vpcEndpointId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-advancedfieldselector.html#cfn-cloudtrail-eventdatastore-advancedfieldselector-field
	//
	Field *string `field:"required" json:"field" yaml:"field"`
	// An operator that includes events that match the last few characters of the event record field specified as the value of `Field` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-advancedfieldselector.html#cfn-cloudtrail-eventdatastore-advancedfieldselector-endswith
	//
	EndsWith *[]*string `field:"optional" json:"endsWith" yaml:"endsWith"`
	// An operator that includes events that match the exact value of the event record field specified as the value of `Field` .
	//
	// This is the only valid operator that you can use with the `readOnly` , `eventCategory` , and `resources.type` fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-advancedfieldselector.html#cfn-cloudtrail-eventdatastore-advancedfieldselector-equals
	//
	EqualTo *[]*string `field:"optional" json:"equalTo" yaml:"equalTo"`
	// An operator that excludes events that match the last few characters of the event record field specified as the value of `Field` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-advancedfieldselector.html#cfn-cloudtrail-eventdatastore-advancedfieldselector-notendswith
	//
	NotEndsWith *[]*string `field:"optional" json:"notEndsWith" yaml:"notEndsWith"`
	// An operator that excludes events that match the exact value of the event record field specified as the value of `Field` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-advancedfieldselector.html#cfn-cloudtrail-eventdatastore-advancedfieldselector-notequals
	//
	NotEquals *[]*string `field:"optional" json:"notEquals" yaml:"notEquals"`
	// An operator that excludes events that match the first few characters of the event record field specified as the value of `Field` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-advancedfieldselector.html#cfn-cloudtrail-eventdatastore-advancedfieldselector-notstartswith
	//
	NotStartsWith *[]*string `field:"optional" json:"notStartsWith" yaml:"notStartsWith"`
	// An operator that includes events that match the first few characters of the event record field specified as the value of `Field` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-advancedfieldselector.html#cfn-cloudtrail-eventdatastore-advancedfieldselector-startswith
	//
	StartsWith *[]*string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

