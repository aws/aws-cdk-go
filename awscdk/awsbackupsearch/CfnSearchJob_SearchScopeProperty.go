package awsbackupsearch


// The search scope for the search job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   searchScopeProperty := &SearchScopeProperty{
//   	BackupResourceTypes: []*string{
//   		jsii.String("backupResourceTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backupsearch-searchjob-searchscope.html
//
type CfnSearchJob_SearchScopeProperty struct {
	// The resource types included in a search.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backupsearch-searchjob-searchscope.html#cfn-backupsearch-searchjob-searchscope-backupresourcetypes
	//
	BackupResourceTypes *[]*string `field:"required" json:"backupResourceTypes" yaml:"backupResourceTypes"`
}

