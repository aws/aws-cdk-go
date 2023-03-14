package awsresourcegroups


// Specifies a single tag key and optional values that you can use to specify membership in a tag-based group.
//
// An AWS resource that doesn't have a matching tag key and value is rejected as a member of the group.
//
// A `TagFilter` object includes two properties: `Key` (a string) and `Values` (a list of strings). Only resources in the account that are tagged with a matching key-value pair are members of the group. The `Values` property of `TagFilter` is optional, but specifying it narrows the query results.
//
// As an example, suppose the `TagFilters` string is `[{"Key": "Stage", "Values": ["Test", "Beta"]}, {"Key": "Storage"}]` . In this case, only resources with all of the following tags are members of the group:
//
// - `Stage` tag key with a value of either `Test` or `Beta`
// - `Storage` tag key with any value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagFilterProperty := &tagFilterProperty{
//   	key: jsii.String("key"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnGroup_TagFilterProperty struct {
	// A string that defines a tag key.
	//
	// Only resources in the account that are tagged with a specified tag key are members of the tag-based resource group.
	//
	// This field is required when the `ResourceQuery` structure's `Type` property is `TAG_FILTERS_1_0` . You must specify at least one tag key.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A list of tag values that can be included in the tag-based resource group.
	//
	// This is optional. If you don't specify a value or values for a key, then an AWS resource with any value for that key is a member.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

