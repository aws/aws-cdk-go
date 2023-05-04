package awskendra


// Specifies the properties, such as relevance tuning and searchability, of an index field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentMetadataConfigurationProperty := &DocumentMetadataConfigurationProperty{
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Relevance: &RelevanceProperty{
//   		Duration: jsii.String("duration"),
//   		Freshness: jsii.Boolean(false),
//   		Importance: jsii.Number(123),
//   		RankOrder: jsii.String("rankOrder"),
//   		ValueImportanceItems: []interface{}{
//   			&ValueImportanceItemProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Search: &SearchProperty{
//   		Displayable: jsii.Boolean(false),
//   		Facetable: jsii.Boolean(false),
//   		Searchable: jsii.Boolean(false),
//   		Sortable: jsii.Boolean(false),
//   	},
//   }
//
type CfnIndex_DocumentMetadataConfigurationProperty struct {
	// The name of the index field.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The data type of the index field.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Provides tuning parameters to determine how the field affects the search results.
	Relevance interface{} `field:"optional" json:"relevance" yaml:"relevance"`
	// Provides information about how the field is used during a search.
	Search interface{} `field:"optional" json:"search" yaml:"search"`
}

