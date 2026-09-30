package awskendra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFeaturedResultsSet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFeaturedResultsSetProps := &CfnFeaturedResultsSetProps{
//   	FeaturedResultsSetName: jsii.String("featuredResultsSetName"),
//   	IndexId: jsii.String("indexId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	FeaturedDocuments: []interface{}{
//   		&FeaturedDocumentProperty{
//   			Id: jsii.String("id"),
//   		},
//   	},
//   	QueryTexts: []*string{
//   		jsii.String("queryTexts"),
//   	},
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-featuredresultsset.html
//
type CfnFeaturedResultsSetProps struct {
	// A name for the set of featured results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-featuredresultsset.html#cfn-kendra-featuredresultsset-featuredresultssetname
	//
	FeaturedResultsSetName *string `field:"required" json:"featuredResultsSetName" yaml:"featuredResultsSetName"`
	// The identifier of the index that you want to use for featuring results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-featuredresultsset.html#cfn-kendra-featuredresultsset-indexid
	//
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
	// A description for the set of featured results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-featuredresultsset.html#cfn-kendra-featuredresultsset-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of document IDs for the documents you want to feature at the top of the search results page.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-featuredresultsset.html#cfn-kendra-featuredresultsset-featureddocuments
	//
	FeaturedDocuments interface{} `field:"optional" json:"featuredDocuments" yaml:"featuredDocuments"`
	// A list of queries for featuring results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-featuredresultsset.html#cfn-kendra-featuredresultsset-querytexts
	//
	QueryTexts *[]*string `field:"optional" json:"queryTexts" yaml:"queryTexts"`
	// The current status of the set of featured results.
	//
	// When the value is ACTIVE, featured results are ready for use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-featuredresultsset.html#cfn-kendra-featuredresultsset-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// A list of key-value pairs that identify or categorize the featured results set.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kendra-featuredresultsset.html#cfn-kendra-featuredresultsset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

