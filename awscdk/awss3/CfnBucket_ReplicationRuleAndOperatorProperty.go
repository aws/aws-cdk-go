package awss3


// A container for specifying rule filters.
//
// The filters determine the subset of objects to which the rule applies. This element is required only if you specify more than one filter.
//
// For example:
//
// - If you specify both a `Prefix` and a `TagFilter` , wrap these filters in an `And` tag.
// - If you specify a filter based on multiple tags, wrap the `TagFilter` elements in an `And` tag.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationRuleAndOperatorProperty := &ReplicationRuleAndOperatorProperty{
//   	Prefix: jsii.String("prefix"),
//   	TagFilters: []interface{}{
//   		&TagFilterProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationruleandoperator.html
//
type CfnBucket_ReplicationRuleAndOperatorProperty struct {
	// An object key name prefix that identifies the subset of objects to which the rule applies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationruleandoperator.html#cfn-s3-bucket-replicationruleandoperator-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// An array of tags containing key and value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-replicationruleandoperator.html#cfn-s3-bucket-replicationruleandoperator-tagfilters
	//
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

