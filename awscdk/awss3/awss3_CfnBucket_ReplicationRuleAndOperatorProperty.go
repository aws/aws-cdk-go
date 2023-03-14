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
//   replicationRuleAndOperatorProperty := &replicationRuleAndOperatorProperty{
//   	prefix: jsii.String("prefix"),
//   	tagFilters: []interface{}{
//   		&tagFilterProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBucket_ReplicationRuleAndOperatorProperty struct {
	// An object key name prefix that identifies the subset of objects to which the rule applies.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// An array of tags containing key and value pairs.
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

