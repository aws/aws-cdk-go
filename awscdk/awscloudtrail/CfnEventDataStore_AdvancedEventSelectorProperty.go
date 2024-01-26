package awscloudtrail


// Advanced event selectors let you create fine-grained selectors for CloudTrail management and data events.
//
// They help you control costs by logging only those events that are important to you. For more information about advanced event selectors, see [Logging management events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-management-events-with-cloudtrail.html) and [Logging data events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html) in the *AWS CloudTrail User Guide* .
//
// You cannot apply both event selectors and advanced event selectors to a trail.
//
// *Supported CloudTrail event record fields for management events*
//
// - `eventCategory` (required)
// - `eventSource`
// - `readOnly`
//
// *Supported CloudTrail event record fields for data events*
//
// - `eventCategory` (required)
// - `resources.type` (required)
// - `readOnly`
// - `eventName`
// - `resources.ARN`
//
// > For event data stores for CloudTrail Insights events, AWS Config configuration items, Audit Manager evidence, or events outside of AWS , the only supported field is `eventCategory` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   advancedEventSelectorProperty := &AdvancedEventSelectorProperty{
//   	FieldSelectors: []interface{}{
//   		&AdvancedFieldSelectorProperty{
//   			Field: jsii.String("field"),
//
//   			// the properties below are optional
//   			EndsWith: []*string{
//   				jsii.String("endsWith"),
//   			},
//   			EqualTo: []*string{
//   				jsii.String("equalTo"),
//   			},
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
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-advancedeventselector.html
//
type CfnEventDataStore_AdvancedEventSelectorProperty struct {
	// Contains all selector statements in an advanced event selector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-advancedeventselector.html#cfn-cloudtrail-eventdatastore-advancedeventselector-fieldselectors
	//
	FieldSelectors interface{} `field:"required" json:"fieldSelectors" yaml:"fieldSelectors"`
	// An optional, descriptive name for an advanced event selector, such as "Log data events for only two S3 buckets".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-eventdatastore-advancedeventselector.html#cfn-cloudtrail-eventdatastore-advancedeventselector-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

