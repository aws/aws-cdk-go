package previewawscodecommitevents


// Type definition for MergeMetadata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mergeMetadata := &MergeMetadata{
//   	IsMerged: []*string{
//   		jsii.String("isMerged"),
//   	},
//   	MergeCommitId: []*string{
//   		jsii.String("mergeCommitId"),
//   	},
//   	MergedBy: []*string{
//   		jsii.String("mergedBy"),
//   	},
//   	MergeOption: []*string{
//   		jsii.String("mergeOption"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_MergeMetadata struct {
	// isMerged property.
	//
	// Specify an array of string values to match this event if the actual value of isMerged is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	IsMerged *[]*string `field:"optional" json:"isMerged" yaml:"isMerged"`
	// mergeCommitId property.
	//
	// Specify an array of string values to match this event if the actual value of mergeCommitId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MergeCommitId *[]*string `field:"optional" json:"mergeCommitId" yaml:"mergeCommitId"`
	// mergedBy property.
	//
	// Specify an array of string values to match this event if the actual value of mergedBy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MergedBy *[]*string `field:"optional" json:"mergedBy" yaml:"mergedBy"`
	// mergeOption property.
	//
	// Specify an array of string values to match this event if the actual value of mergeOption is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MergeOption *[]*string `field:"optional" json:"mergeOption" yaml:"mergeOption"`
}

