package previewawsdlmevents


// Type definition for RequestParameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   requestParameters := &RequestParameters{
//   	Description: []*string{
//   		jsii.String("description"),
//   	},
//   	ExecutionRoleArn: []*string{
//   		jsii.String("executionRoleArn"),
//   	},
//   	PolicyDetails: &PolicyDetails{
//   		PolicyType: []*string{
//   			jsii.String("policyType"),
//   		},
//   		ResourceTypes: []*string{
//   			jsii.String("resourceTypes"),
//   		},
//   		Schedules: []PolicyDetailsItem{
//   			&PolicyDetailsItem{
//   				CopyTags: []*string{
//   					jsii.String("copyTags"),
//   				},
//   				CreateRule: &CreateRule{
//   					Interval: []*string{
//   						jsii.String("interval"),
//   					},
//   					IntervalUnit: []*string{
//   						jsii.String("intervalUnit"),
//   					},
//   					Times: []*string{
//   						jsii.String("times"),
//   					},
//   				},
//   				FastRestoreRule: &FastRestoreRule{
//   					AvailabilityZones: []*string{
//   						jsii.String("availabilityZones"),
//   					},
//   					Count: []*string{
//   						jsii.String("count"),
//   					},
//   				},
//   				Name: []*string{
//   					jsii.String("name"),
//   				},
//   				RetainRule: &RetainRule{
//   					Count: []*string{
//   						jsii.String("count"),
//   					},
//   				},
//   				TagsToAdd: []PolicyDetailsItemItem{
//   					&PolicyDetailsItemItem{
//   						Key: []*string{
//   							jsii.String("key"),
//   						},
//   						Value: []*string{
//   							jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		TargetTags: []PolicyDetailsItem1{
//   			&PolicyDetailsItem1{
//   				Key: []*string{
//   					jsii.String("key"),
//   				},
//   				Value: []*string{
//   					jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	PolicyId: []*string{
//   		jsii.String("policyId"),
//   	},
//   	State: []*string{
//   		jsii.String("state"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_RequestParameters struct {
	// Description property.
	//
	// Specify an array of string values to match this event if the actual value of Description is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Description *[]*string `field:"optional" json:"description" yaml:"description"`
	// ExecutionRoleArn property.
	//
	// Specify an array of string values to match this event if the actual value of ExecutionRoleArn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ExecutionRoleArn *[]*string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// PolicyDetails property.
	//
	// Specify an array of string values to match this event if the actual value of PolicyDetails is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PolicyDetails *AWSAPICallViaCloudTrail_PolicyDetails `field:"optional" json:"policyDetails" yaml:"policyDetails"`
	// policyId property.
	//
	// Specify an array of string values to match this event if the actual value of policyId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PolicyId *[]*string `field:"optional" json:"policyId" yaml:"policyId"`
	// State property.
	//
	// Specify an array of string values to match this event if the actual value of State is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	State *[]*string `field:"optional" json:"state" yaml:"state"`
}

