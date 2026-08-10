package awsbackupsearch


// Properties for defining a `CfnSearchJob`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSearchJobProps := &CfnSearchJobProps{
//   	SearchScope: &SearchScopeProperty{
//   		BackupResourceTypes: []*string{
//   			jsii.String("backupResourceTypes"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupsearch-searchjob.html
//
type CfnSearchJobProps struct {
	// The search scope for the search job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupsearch-searchjob.html#cfn-backupsearch-searchjob-searchscope
	//
	SearchScope interface{} `field:"required" json:"searchScope" yaml:"searchScope"`
	// The name of the search job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupsearch-searchjob.html#cfn-backupsearch-searchjob-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags associated with the search job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupsearch-searchjob.html#cfn-backupsearch-searchjob-tags
	//
	Tags *[]*CfnSearchJob_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

