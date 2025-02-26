package awsappsync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for SourceApiAssociation which associates an AppSync Source API with an AppSync Merged API.
//
// Example:
//   sourceApi := appsync.NewGraphqlApi(this, jsii.String("FirstSourceAPI"), &GraphqlApiProps{
//   	Name: jsii.String("FirstSourceAPI"),
//   	Definition: appsync.Definition_FromFile(path.join(__dirname, jsii.String("appsync.merged-api-1.graphql"))),
//   })
//
//   importedMergedApi := appsync.GraphqlApi_FromGraphqlApiAttributes(this, jsii.String("ImportedMergedApi"), &GraphqlApiAttributes{
//   	GraphqlApiId: jsii.String("MyApiId"),
//   	GraphqlApiArn: jsii.String("MyApiArn"),
//   })
//
//   importedExecutionRole := iam.Role_FromRoleArn(this, jsii.String("ExecutionRole"), jsii.String("arn:aws:iam::ACCOUNT:role/MyExistingRole"))
//   appsync.NewSourceApiAssociation(this, jsii.String("SourceApiAssociation2"), &SourceApiAssociationProps{
//   	SourceApi: sourceApi,
//   	MergedApi: importedMergedApi,
//   	MergeType: appsync.MergeType_MANUAL_MERGE,
//   	MergedApiExecutionRole: importedExecutionRole,
//   })
//
type SourceApiAssociationProps struct {
	// The merged api to associate.
	MergedApi IGraphqlApi `field:"required" json:"mergedApi" yaml:"mergedApi"`
	// The merged api execution role for adding the access policy for the source api.
	MergedApiExecutionRole awsiam.IRole `field:"required" json:"mergedApiExecutionRole" yaml:"mergedApiExecutionRole"`
	// The source api to associate.
	SourceApi IGraphqlApi `field:"required" json:"sourceApi" yaml:"sourceApi"`
	// The description of the source api association.
	// Default: - None.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The merge type for the source.
	// Default: - AUTO_MERGE.
	//
	MergeType MergeType `field:"optional" json:"mergeType" yaml:"mergeType"`
}

