package awscloudtrail


// Advanced event selectors let you create fine-grained selectors for the following AWS CloudTrail event record ﬁelds.
//
// They help you control costs by logging only those events that are important to you. For more information about advanced event selectors, see [Logging data events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html) in the *AWS CloudTrail User Guide* .
//
// - `readOnly`
// - `eventSource`
// - `eventName`
// - `eventCategory`
// - `resources.type`
// - `resources.ARN`
//
// You cannot apply both event selectors and advanced event selectors to a trail.
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
type CfnEventDataStore_AdvancedEventSelectorProperty struct {
	// Contains all selector statements in an advanced event selector.
	FieldSelectors interface{} `field:"required" json:"fieldSelectors" yaml:"fieldSelectors"`
	// An optional, descriptive name for an advanced event selector, such as "Log data events for only two S3 buckets".
	Name *string `field:"optional" json:"name" yaml:"name"`
}

