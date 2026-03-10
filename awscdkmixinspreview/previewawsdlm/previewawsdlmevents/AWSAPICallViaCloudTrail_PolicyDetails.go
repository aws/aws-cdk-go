package previewawsdlmevents


// Type definition for PolicyDetails.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   policyDetails := &PolicyDetails{
//   	PolicyType: []*string{
//   		jsii.String("policyType"),
//   	},
//   	ResourceTypes: []*string{
//   		jsii.String("resourceTypes"),
//   	},
//   	Schedules: []PolicyDetailsItem{
//   		&PolicyDetailsItem{
//   			CopyTags: []*string{
//   				jsii.String("copyTags"),
//   			},
//   			CreateRule: &CreateRule{
//   				Interval: []*string{
//   					jsii.String("interval"),
//   				},
//   				IntervalUnit: []*string{
//   					jsii.String("intervalUnit"),
//   				},
//   				Times: []*string{
//   					jsii.String("times"),
//   				},
//   			},
//   			FastRestoreRule: &FastRestoreRule{
//   				AvailabilityZones: []*string{
//   					jsii.String("availabilityZones"),
//   				},
//   				Count: []*string{
//   					jsii.String("count"),
//   				},
//   			},
//   			Name: []*string{
//   				jsii.String("name"),
//   			},
//   			RetainRule: &RetainRule{
//   				Count: []*string{
//   					jsii.String("count"),
//   				},
//   			},
//   			TagsToAdd: []PolicyDetailsItemItem{
//   				&PolicyDetailsItemItem{
//   					Key: []*string{
//   						jsii.String("key"),
//   					},
//   					Value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	TargetTags: []PolicyDetailsItem1{
//   		&PolicyDetailsItem1{
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
type AWSAPICallViaCloudTrail_PolicyDetails struct {
	// PolicyType property.
	//
	// Specify an array of string values to match this event if the actual value of PolicyType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PolicyType *[]*string `field:"optional" json:"policyType" yaml:"policyType"`
	// ResourceTypes property.
	//
	// Specify an array of string values to match this event if the actual value of ResourceTypes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
	// Schedules property.
	//
	// Specify an array of string values to match this event if the actual value of Schedules is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Schedules *[]*AWSAPICallViaCloudTrail_PolicyDetailsItem `field:"optional" json:"schedules" yaml:"schedules"`
	// TargetTags property.
	//
	// Specify an array of string values to match this event if the actual value of TargetTags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TargetTags *[]*AWSAPICallViaCloudTrail_PolicyDetailsItem1 `field:"optional" json:"targetTags" yaml:"targetTags"`
}

