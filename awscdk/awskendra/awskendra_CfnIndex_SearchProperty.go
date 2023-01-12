package awskendra


// Provides information about how a custom index field is used during a search.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   searchProperty := &searchProperty{
//   	displayable: jsii.Boolean(false),
//   	facetable: jsii.Boolean(false),
//   	searchable: jsii.Boolean(false),
//   	sortable: jsii.Boolean(false),
//   }
//
type CfnIndex_SearchProperty struct {
	// Determines whether the field is returned in the query response.
	//
	// The default is `true` .
	Displayable interface{} `field:"optional" json:"displayable" yaml:"displayable"`
	// Indicates that the field can be used to create search facets, a count of results for each value in the field.
	//
	// The default is `false` .
	Facetable interface{} `field:"optional" json:"facetable" yaml:"facetable"`
	// Determines whether the field is used in the search.
	//
	// If the `Searchable` field is `true` , you can use relevance tuning to manually tune how Amazon Kendra weights the field in the search. The default is `true` for string fields and `false` for number and date fields.
	Searchable interface{} `field:"optional" json:"searchable" yaml:"searchable"`
	// Indicates that the field can be used to sort the search results.
	//
	// The default is `false` .
	Sortable interface{} `field:"optional" json:"sortable" yaml:"sortable"`
}

