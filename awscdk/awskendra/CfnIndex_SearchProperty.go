package awskendra


// Provides information about how a custom index field is used during a search.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   searchProperty := &SearchProperty{
//   	Displayable: jsii.Boolean(false),
//   	Facetable: jsii.Boolean(false),
//   	Searchable: jsii.Boolean(false),
//   	Sortable: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-search.html
//
type CfnIndex_SearchProperty struct {
	// Determines whether the field is returned in the query response.
	//
	// The default is `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-search.html#cfn-kendra-index-search-displayable
	//
	Displayable interface{} `field:"optional" json:"displayable" yaml:"displayable"`
	// Indicates that the field can be used to create search facets, a count of results for each value in the field.
	//
	// The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-search.html#cfn-kendra-index-search-facetable
	//
	Facetable interface{} `field:"optional" json:"facetable" yaml:"facetable"`
	// Determines whether the field is used in the search.
	//
	// If the `Searchable` field is `true` , you can use relevance tuning to manually tune how Amazon Kendra weights the field in the search. The default is `true` for string fields and `false` for number and date fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-search.html#cfn-kendra-index-search-searchable
	//
	Searchable interface{} `field:"optional" json:"searchable" yaml:"searchable"`
	// Determines whether the field can be used to sort the results of a query.
	//
	// The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-search.html#cfn-kendra-index-search-sortable
	//
	Sortable interface{} `field:"optional" json:"sortable" yaml:"sortable"`
}

