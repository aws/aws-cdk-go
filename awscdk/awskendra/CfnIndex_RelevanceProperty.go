package awskendra


// Provides information for tuning the relevance of a field in a search.
//
// When a query includes terms that match the field, the results are given a boost in the response based on these tuning parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   relevanceProperty := &RelevanceProperty{
//   	Duration: jsii.String("duration"),
//   	Freshness: jsii.Boolean(false),
//   	Importance: jsii.Number(123),
//   	RankOrder: jsii.String("rankOrder"),
//   	ValueImportanceItems: []interface{}{
//   		&ValueImportanceItemProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-relevance.html
//
type CfnIndex_RelevanceProperty struct {
	// Specifies the time period that the boost applies to.
	//
	// For example, to make the boost apply to documents with the field value within the last month, you would use "2628000s". Once the field value is beyond the specified range, the effect of the boost drops off. The higher the importance, the faster the effect drops off. If you don't specify a value, the default is 3 months. The value of the field is a numeric string followed by the character "s", for example "86400s" for one day, or "604800s" for one week.
	//
	// Only applies to `DATE` fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-relevance.html#cfn-kendra-index-relevance-duration
	//
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Indicates that this field determines how "fresh" a document is.
	//
	// For example, if document 1 was created on November 5, and document 2 was created on October 31, document 1 is "fresher" than document 2. You can only set the `Freshness` field on one `DATE` type field. Only applies to `DATE` fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-relevance.html#cfn-kendra-index-relevance-freshness
	//
	Freshness interface{} `field:"optional" json:"freshness" yaml:"freshness"`
	// The relative importance of the field in the search.
	//
	// Larger numbers provide more of a boost than smaller numbers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-relevance.html#cfn-kendra-index-relevance-importance
	//
	Importance *float64 `field:"optional" json:"importance" yaml:"importance"`
	// Determines how values should be interpreted.
	//
	// When the `RankOrder` field is `ASCENDING` , higher numbers are better. For example, a document with a rating score of 10 is higher ranking than a document with a rating score of 1.
	//
	// When the `RankOrder` field is `DESCENDING` , lower numbers are better. For example, in a task tracking application, a priority 1 task is more important than a priority 5 task.
	//
	// Only applies to `LONG` and `DOUBLE` fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-relevance.html#cfn-kendra-index-relevance-rankorder
	//
	RankOrder *string `field:"optional" json:"rankOrder" yaml:"rankOrder"`
	// An array of key-value pairs for different boosts when they appear in the search result list.
	//
	// For example, if you want to boost query terms that match the "department" field in the result, query terms that match this field are boosted in the result. You can add entries from the department field to boost documents with those values higher.
	//
	// For example, you can add entries to the map with names of departments. If you add "HR", 5 and "Legal",3 those departments are given special attention when they appear in the metadata of a document.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-index-relevance.html#cfn-kendra-index-relevance-valueimportanceitems
	//
	ValueImportanceItems interface{} `field:"optional" json:"valueImportanceItems" yaml:"valueImportanceItems"`
}

