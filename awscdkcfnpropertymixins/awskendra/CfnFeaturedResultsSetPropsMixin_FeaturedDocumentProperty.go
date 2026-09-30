package awskendra


// A featured document.
//
// This document is displayed at the top of the search results page.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   featuredDocumentProperty := &FeaturedDocumentProperty{
//   	Id: jsii.String("id"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-featuredresultsset-featureddocument.html
//
type CfnFeaturedResultsSetPropsMixin_FeaturedDocumentProperty struct {
	// The identifier of the document to feature in the search results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-featuredresultsset-featureddocument.html#cfn-kendra-featuredresultsset-featureddocument-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
}

