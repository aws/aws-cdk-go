package awsopensearchserverless


// Properties for defining a `CfnCollectionIndex`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCollectionIndexProps := &CfnCollectionIndexProps{
//   	Id: jsii.String("id"),
//   	IndexName: jsii.String("indexName"),
//
//   	// the properties below are optional
//   	IndexSchema: jsii.String("indexSchema"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collectionindex.html
//
type CfnCollectionIndexProps struct {
	// The identifier of the collection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collectionindex.html#cfn-opensearchserverless-collectionindex-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The name of the collection index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collectionindex.html#cfn-opensearchserverless-collectionindex-indexname
	//
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// The Mappings for the collection index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchserverless-collectionindex.html#cfn-opensearchserverless-collectionindex-indexschema
	//
	IndexSchema *string `field:"optional" json:"indexSchema" yaml:"indexSchema"`
}

