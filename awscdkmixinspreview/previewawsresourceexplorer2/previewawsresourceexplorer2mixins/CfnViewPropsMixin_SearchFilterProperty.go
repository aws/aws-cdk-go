package previewawsresourceexplorer2mixins


// A search filter defines which resources can be part of a search query result set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   searchFilterProperty := &SearchFilterProperty{
//   	FilterString: jsii.String("filterString"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourceexplorer2-view-searchfilter.html
//
type CfnViewPropsMixin_SearchFilterProperty struct {
	// The string that contains the search keywords, prefixes, and operators to control the results that can be returned by a Search operation.
	//
	// For information about the supported syntax, see [Search query reference](https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html) in the *AWS Resource Explorer User Guide* .
	//
	// > This query string in the context of this operation supports only [filter prefixes](https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html#query-syntax-filters) with optional [operators](https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html#query-syntax-operators) . It doesn't support free-form text. For example, the string `region:us* service:ec2 -tag:stage=prod` includes all Amazon EC2 resources in any AWS Region that begin with the letters `us` and are *not* tagged with a key `Stage` that has the value `prod` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resourceexplorer2-view-searchfilter.html#cfn-resourceexplorer2-view-searchfilter-filterstring
	//
	FilterString *string `field:"optional" json:"filterString" yaml:"filterString"`
}

