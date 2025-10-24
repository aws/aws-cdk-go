package awsappsync


// Configuration of source API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var graphqlApi GraphqlApi
//
//   sourceApi := &SourceApi{
//   	SourceApi: graphqlApi,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	MergeType: awscdk.Aws_appsync.MergeType_MANUAL_MERGE,
//   }
//
type SourceApi struct {
	// Source API that is associated with the merged API.
	SourceApi IGraphqlApi `field:"required" json:"sourceApi" yaml:"sourceApi"`
	// Description of the Source API asssociation.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Merging option used to associate the source API to the Merged API.
	// Default: - Auto merge. The merge is triggered automatically when the source API has changed
	//
	MergeType MergeType `field:"optional" json:"mergeType" yaml:"mergeType"`
}

