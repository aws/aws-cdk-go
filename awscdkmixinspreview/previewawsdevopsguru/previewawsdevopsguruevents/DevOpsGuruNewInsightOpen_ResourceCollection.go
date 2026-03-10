package previewawsdevopsguruevents


// Type definition for ResourceCollection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceCollection := &ResourceCollection{
//   	CloudFormation: &CloudFormation{
//   		StackNames: []*string{
//   			jsii.String("stackNames"),
//   		},
//   	},
//   	Tags: []TagType{
//   		&TagType{
//   			AppBoundaryKey: []*string{
//   				jsii.String("appBoundaryKey"),
//   			},
//   			TagValues: []*string{
//   				jsii.String("tagValues"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type DevOpsGuruNewInsightOpen_ResourceCollection struct {
	// cloudFormation property.
	//
	// Specify an array of string values to match this event if the actual value of cloudFormation is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CloudFormation *DevOpsGuruNewInsightOpen_CloudFormation `field:"optional" json:"cloudFormation" yaml:"cloudFormation"`
	// tags property.
	//
	// Specify an array of string values to match this event if the actual value of tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]*DevOpsGuruNewInsightOpen_TagType `field:"optional" json:"tags" yaml:"tags"`
}

