package previewawsdlmevents


// Type definition for PolicyDetailsItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   policyDetailsItem := &PolicyDetailsItem{
//   	CopyTags: []*string{
//   		jsii.String("copyTags"),
//   	},
//   	CreateRule: &CreateRule{
//   		Interval: []*string{
//   			jsii.String("interval"),
//   		},
//   		IntervalUnit: []*string{
//   			jsii.String("intervalUnit"),
//   		},
//   		Times: []*string{
//   			jsii.String("times"),
//   		},
//   	},
//   	FastRestoreRule: &FastRestoreRule{
//   		AvailabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		Count: []*string{
//   			jsii.String("count"),
//   		},
//   	},
//   	Name: []*string{
//   		jsii.String("name"),
//   	},
//   	RetainRule: &RetainRule{
//   		Count: []*string{
//   			jsii.String("count"),
//   		},
//   	},
//   	TagsToAdd: []PolicyDetailsItemItem{
//   		&PolicyDetailsItemItem{
//   			Key: []*string{
//   				jsii.String("key"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_PolicyDetailsItem struct {
	// CopyTags property.
	//
	// Specify an array of string values to match this event if the actual value of CopyTags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CopyTags *[]*string `field:"optional" json:"copyTags" yaml:"copyTags"`
	// CreateRule property.
	//
	// Specify an array of string values to match this event if the actual value of CreateRule is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CreateRule *AWSAPICallViaCloudTrail_CreateRule `field:"optional" json:"createRule" yaml:"createRule"`
	// FastRestoreRule property.
	//
	// Specify an array of string values to match this event if the actual value of FastRestoreRule is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	FastRestoreRule *AWSAPICallViaCloudTrail_FastRestoreRule `field:"optional" json:"fastRestoreRule" yaml:"fastRestoreRule"`
	// Name property.
	//
	// Specify an array of string values to match this event if the actual value of Name is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Name *[]*string `field:"optional" json:"name" yaml:"name"`
	// RetainRule property.
	//
	// Specify an array of string values to match this event if the actual value of RetainRule is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RetainRule *AWSAPICallViaCloudTrail_RetainRule `field:"optional" json:"retainRule" yaml:"retainRule"`
	// TagsToAdd property.
	//
	// Specify an array of string values to match this event if the actual value of TagsToAdd is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TagsToAdd *[]*AWSAPICallViaCloudTrail_PolicyDetailsItemItem `field:"optional" json:"tagsToAdd" yaml:"tagsToAdd"`
}

