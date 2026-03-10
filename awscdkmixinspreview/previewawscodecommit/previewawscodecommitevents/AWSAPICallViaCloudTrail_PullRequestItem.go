package previewawscodecommitevents


// Type definition for PullRequestItem.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pullRequestItem := &PullRequestItem{
//   	DestinationCommit: []*string{
//   		jsii.String("destinationCommit"),
//   	},
//   	DestinationReference: []*string{
//   		jsii.String("destinationReference"),
//   	},
//   	MergeBase: []*string{
//   		jsii.String("mergeBase"),
//   	},
//   	MergeMetadata: &MergeMetadata{
//   		IsMerged: []*string{
//   			jsii.String("isMerged"),
//   		},
//   		MergeCommitId: []*string{
//   			jsii.String("mergeCommitId"),
//   		},
//   		MergedBy: []*string{
//   			jsii.String("mergedBy"),
//   		},
//   		MergeOption: []*string{
//   			jsii.String("mergeOption"),
//   		},
//   	},
//   	RepositoryName: []*string{
//   		jsii.String("repositoryName"),
//   	},
//   	SourceCommit: []*string{
//   		jsii.String("sourceCommit"),
//   	},
//   	SourceReference: []*string{
//   		jsii.String("sourceReference"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_PullRequestItem struct {
	// destinationCommit property.
	//
	// Specify an array of string values to match this event if the actual value of destinationCommit is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DestinationCommit *[]*string `field:"optional" json:"destinationCommit" yaml:"destinationCommit"`
	// destinationReference property.
	//
	// Specify an array of string values to match this event if the actual value of destinationReference is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DestinationReference *[]*string `field:"optional" json:"destinationReference" yaml:"destinationReference"`
	// mergeBase property.
	//
	// Specify an array of string values to match this event if the actual value of mergeBase is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MergeBase *[]*string `field:"optional" json:"mergeBase" yaml:"mergeBase"`
	// mergeMetadata property.
	//
	// Specify an array of string values to match this event if the actual value of mergeMetadata is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MergeMetadata *AWSAPICallViaCloudTrail_MergeMetadata `field:"optional" json:"mergeMetadata" yaml:"mergeMetadata"`
	// repositoryName property.
	//
	// Specify an array of string values to match this event if the actual value of repositoryName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RepositoryName *[]*string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// sourceCommit property.
	//
	// Specify an array of string values to match this event if the actual value of sourceCommit is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceCommit *[]*string `field:"optional" json:"sourceCommit" yaml:"sourceCommit"`
	// sourceReference property.
	//
	// Specify an array of string values to match this event if the actual value of sourceReference is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceReference *[]*string `field:"optional" json:"sourceReference" yaml:"sourceReference"`
}

