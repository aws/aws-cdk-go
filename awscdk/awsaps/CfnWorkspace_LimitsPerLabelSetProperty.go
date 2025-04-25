package awsaps


// Label set and its associated limits.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   limitsPerLabelSetProperty := &LimitsPerLabelSetProperty{
//   	LabelSet: []interface{}{
//   		&LabelProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Limits: &LimitsPerLabelSetEntryProperty{
//   		MaxSeries: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-limitsperlabelset.html
//
type CfnWorkspace_LimitsPerLabelSetProperty struct {
	// An array of series labels.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-limitsperlabelset.html#cfn-aps-workspace-limitsperlabelset-labelset
	//
	LabelSet interface{} `field:"required" json:"labelSet" yaml:"labelSet"`
	// Limits that can be applied to a label set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aps-workspace-limitsperlabelset.html#cfn-aps-workspace-limitsperlabelset-limits
	//
	Limits interface{} `field:"required" json:"limits" yaml:"limits"`
}

