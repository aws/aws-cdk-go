package awsresourceexplorer2


// Properties for defining a `CfnView`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnViewProps := &CfnViewProps{
//   	ViewName: jsii.String("viewName"),
//
//   	// the properties below are optional
//   	Filters: &FiltersProperty{
//   		FilterString: jsii.String("filterString"),
//   	},
//   	IncludedProperties: []interface{}{
//   		&IncludedPropertyProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Scope: jsii.String("scope"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourceexplorer2-view.html
//
type CfnViewProps struct {
	// The name of the new view.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourceexplorer2-view.html#cfn-resourceexplorer2-view-viewname
	//
	ViewName *string `field:"required" json:"viewName" yaml:"viewName"`
	// An array of strings that include search keywords, prefixes, and operators that filter the results that are returned for queries made using this view.
	//
	// When you use this view in a [Search](https://docs.aws.amazon.com/resource-explorer/latest/apireference/API_Search.html) operation, the filter string is combined with the search's `QueryString` parameter using a logical `AND` operator.
	//
	// For information about the supported syntax, see [Search query reference for Resource Explorer](https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html) in the *AWS Resource Explorer User Guide* .
	//
	// > This query string in the context of this operation supports only [filter prefixes](https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html#query-syntax-filters) with optional [operators](https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html#query-syntax-operators) . It doesn't support free-form text. For example, the string `region:us* service:ec2 -tag:stage=prod` includes all Amazon EC2 resources in any AWS Region that begin with the letters `us` and are *not* tagged with a key `Stage` that has the value `prod` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourceexplorer2-view.html#cfn-resourceexplorer2-view-filters
	//
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// A list of fields that provide additional information about the view.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourceexplorer2-view.html#cfn-resourceexplorer2-view-includedproperties
	//
	IncludedProperties interface{} `field:"optional" json:"includedProperties" yaml:"includedProperties"`
	// The root ARN of the account, an organizational unit (OU), or an organization ARN.
	//
	// If left empty, the default is account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourceexplorer2-view.html#cfn-resourceexplorer2-view-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Tag key and value pairs that are attached to the view.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-resourceexplorer2-view.html#cfn-resourceexplorer2-view-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

