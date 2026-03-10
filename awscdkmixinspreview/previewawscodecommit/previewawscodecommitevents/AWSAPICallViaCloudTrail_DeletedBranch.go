package previewawscodecommitevents


// Type definition for DeletedBranch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   deletedBranch := &DeletedBranch{
//   	BranchName: []*string{
//   		jsii.String("branchName"),
//   	},
//   	CommitId: []*string{
//   		jsii.String("commitId"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_DeletedBranch struct {
	// branchName property.
	//
	// Specify an array of string values to match this event if the actual value of branchName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BranchName *[]*string `field:"optional" json:"branchName" yaml:"branchName"`
	// commitId property.
	//
	// Specify an array of string values to match this event if the actual value of commitId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	CommitId *[]*string `field:"optional" json:"commitId" yaml:"commitId"`
}

