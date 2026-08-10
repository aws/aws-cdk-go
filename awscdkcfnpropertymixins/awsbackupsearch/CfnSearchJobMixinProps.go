package awsbackupsearch


// Properties for CfnSearchJobPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnSearchJobMixinProps := &CfnSearchJobMixinProps{
//   	Name: jsii.String("name"),
//   	SearchScope: &SearchScopeProperty{
//   		BackupResourceTypes: []*string{
//   			jsii.String("backupResourceTypes"),
//   		},
//   	},
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
type CfnSearchJobMixinProps struct {
	// The name of the search job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupsearch-searchjob.html#cfn-backupsearch-searchjob-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The search scope for the search job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupsearch-searchjob.html#cfn-backupsearch-searchjob-searchscope
	//
	SearchScope interface{} `field:"optional" json:"searchScope" yaml:"searchScope"`
	// Tags associated with the search job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backupsearch-searchjob.html#cfn-backupsearch-searchjob-tags
	//
	Tags *[]*CfnSearchJobPropsMixin_TagsItemsProperty `field:"optional" json:"tags" yaml:"tags"`
}

