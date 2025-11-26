package previewawscloudtrailmixins


// Advanced event selectors let you create fine-grained selectors for AWS CloudTrail management, data, and network activity events.
//
// They help you control costs by logging only those events that are important to you. For more information about configuring advanced event selectors, see the [Logging data events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html) , [Logging network activity events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-network-events-with-cloudtrail.html) , and [Logging management events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-management-events-with-cloudtrail.html) topics in the *AWS CloudTrail User Guide* .
//
// You cannot apply both event selectors and advanced event selectors to a trail.
//
// *Supported CloudTrail event record fields for management events*
//
// - `eventCategory` (required)
// - `eventSource`
// - `readOnly`
//
// The following additional fields are available for event data stores:
//
// - `eventName`
// - `eventType`
// - `sessionCredentialFromConsole`
// - `userIdentity.arn`
//
// *Supported CloudTrail event record fields for data events*
//
// - `eventCategory` (required)
// - `eventName`
// - `eventSource`
// - `eventType`
// - `resources.ARN`
// - `resources.type` (required)
// - `readOnly`
// - `sessionCredentialFromConsole`
// - `userIdentity.arn`
//
// *Supported CloudTrail event record fields for network activity events*
//
// - `eventCategory` (required)
// - `eventSource` (required)
// - `eventName`
// - `errorCode` - The only valid value for `errorCode` is `VpceAccessDenied` .
// - `vpcEndpointId`
//
// > For event data stores for CloudTrail Insights events, AWS Config configuration items, Audit Manager evidence, or events outside of AWS , the only supported field is `eventCategory` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   advancedEventSelectorProperty := &AdvancedEventSelectorProperty{
//   	FieldSelectors: []interface{}{
//   		&AdvancedFieldSelectorProperty{
//   			EndsWith: []*string{
//   				jsii.String("endsWith"),
//   			},
//   			EqualTo: []*string{
//   				jsii.String("equalTo"),
//   			},
//   			Field: jsii.String("field"),
//   			NotEndsWith: []*string{
//   				jsii.String("notEndsWith"),
//   			},
//   			NotEquals: []*string{
//   				jsii.String("notEquals"),
//   			},
//   			NotStartsWith: []*string{
//   				jsii.String("notStartsWith"),
//   			},
//   			StartsWith: []*string{
//   				jsii.String("startsWith"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedeventselector.html
//
type CfnTrailPropsMixin_AdvancedEventSelectorProperty struct {
	// Contains all selector statements in an advanced event selector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedeventselector.html#cfn-cloudtrail-trail-advancedeventselector-fieldselectors
	//
	FieldSelectors interface{} `field:"optional" json:"fieldSelectors" yaml:"fieldSelectors"`
	// An optional, descriptive name for an advanced event selector, such as "Log data events for only two S3 buckets".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-trail-advancedeventselector.html#cfn-cloudtrail-trail-advancedeventselector-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

