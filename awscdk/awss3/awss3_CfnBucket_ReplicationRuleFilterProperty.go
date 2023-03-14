package awss3


// A filter that identifies the subset of objects to which the replication rule applies.
//
// A `Filter` must specify exactly one `Prefix` , `TagFilter` , or an `And` child element.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   replicationRuleFilterProperty := &replicationRuleFilterProperty{
//   	and: &replicationRuleAndOperatorProperty{
//   		prefix: jsii.String("prefix"),
//   		tagFilters: []interface{}{
//   			&tagFilterProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	prefix: jsii.String("prefix"),
//   	tagFilter: &tagFilterProperty{
//   		key: jsii.String("key"),
//   		value: jsii.String("value"),
//   	},
//   }
//
type CfnBucket_ReplicationRuleFilterProperty struct {
	// A container for specifying rule filters.
	//
	// The filters determine the subset of objects to which the rule applies. This element is required only if you specify more than one filter. For example:
	//
	// - If you specify both a `Prefix` and a `TagFilter` , wrap these filters in an `And` tag.
	// - If you specify a filter based on multiple tags, wrap the `TagFilter` elements in an `And` tag.
	And interface{} `field:"optional" json:"and" yaml:"and"`
	// An object key name prefix that identifies the subset of objects to which the rule applies.
	//
	// > Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints) .
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// A container for specifying a tag key and value.
	//
	// The rule applies only to objects that have the tag in their tag set.
	TagFilter interface{} `field:"optional" json:"tagFilter" yaml:"tagFilter"`
}

