package awsappsync


// The attributes for imported AppSync Source Api Association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var graphqlApi graphqlApi
//
//   sourceApiAssociationAttributes := &SourceApiAssociationAttributes{
//   	AssociationArn: jsii.String("associationArn"),
//   	MergedApi: graphqlApi,
//   	SourceApi: graphqlApi,
//   }
//
type SourceApiAssociationAttributes struct {
	// The association arn.
	AssociationArn *string `field:"required" json:"associationArn" yaml:"associationArn"`
	// The merged api in the association.
	MergedApi IGraphqlApi `field:"required" json:"mergedApi" yaml:"mergedApi"`
	// The source api in the association.
	SourceApi IGraphqlApi `field:"required" json:"sourceApi" yaml:"sourceApi"`
}

