package awskendra


// Specifies the properties of a custom index field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentMetadataConfigurationProperty := &documentMetadataConfigurationProperty{
//   	name: jsii.String("name"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	relevance: &relevanceProperty{
//   		duration: jsii.String("duration"),
//   		freshness: jsii.Boolean(false),
//   		importance: jsii.Number(123),
//   		rankOrder: jsii.String("rankOrder"),
//   		valueImportanceItems: []interface{}{
//   			&valueImportanceItemProperty{
//   				key: jsii.String("key"),
//   				value: jsii.Number(123),
//   			},
//   		},
//   	},
//   	search: &searchProperty{
//   		displayable: jsii.Boolean(false),
//   		facetable: jsii.Boolean(false),
//   		searchable: jsii.Boolean(false),
//   		sortable: jsii.Boolean(false),
//   	},
//   }
//
type CfnIndex_DocumentMetadataConfigurationProperty struct {
	// The name of the index field.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The data type of the index field.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Provides manual tuning parameters to determine how the field affects the search results.
	Relevance interface{} `field:"optional" json:"relevance" yaml:"relevance"`
	// Provides information about how the field is used during a search.
	Search interface{} `field:"optional" json:"search" yaml:"search"`
}

