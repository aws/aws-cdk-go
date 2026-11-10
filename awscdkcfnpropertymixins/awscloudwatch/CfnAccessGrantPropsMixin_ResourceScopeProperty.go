package awscloudwatch


// A single resource scope entry within an access grant.
//
// Associates a resource type with optional ARN patterns, tag conditions, and row-level filters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   resourceScopeProperty := &ResourceScopeProperty{
//   	ResourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//   	ResourceType: jsii.String("resourceType"),
//   	RowScopeGroups: []interface{}{
//   		[]interface{}{
//   			&RowScopeProperty{
//   				Field: jsii.String("field"),
//   				Operator: jsii.String("operator"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	SignalTypes: []*string{
//   		jsii.String("signalTypes"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-resourcescope.html
//
type CfnAccessGrantPropsMixin_ResourceScopeProperty struct {
	// Specific resource ARNs or ARN patterns.
	//
	// When set, actions are limited to these resources. When absent, defaults to all resources of the type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-resourcescope.html#cfn-cloudwatch-accessgrant-resourcescope-resourcearns
	//
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// The resource type name, for example DataSet or OmniDashboard.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-resourcescope.html#cfn-cloudwatch-accessgrant-resourcescope-resourcetype
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Row-level filters for this scope, as an OR of AND-groups: a row is visible when it matches every filter in any one group.
	//
	// Requires SignalTypes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-resourcescope.html#cfn-cloudwatch-accessgrant-resourcescope-rowscopegroups
	//
	RowScopeGroups interface{} `field:"optional" json:"rowScopeGroups" yaml:"rowScopeGroups"`
	// Signal types this scope's row filtering applies to.
	//
	// Required when RowScopeGroups is set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-resourcescope.html#cfn-cloudwatch-accessgrant-resourcescope-signaltypes
	//
	SignalTypes *[]*string `field:"optional" json:"signalTypes" yaml:"signalTypes"`
	// Tag-based conditions for dynamic resource scoping.
	//
	// Access applies only to resources carrying all of the specified tag key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-accessgrant-resourcescope.html#cfn-cloudwatch-accessgrant-resourcescope-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

