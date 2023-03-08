package awsmediapackage

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnAsset`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssetProps := &cfnAssetProps{
//   	id: jsii.String("id"),
//   	packagingGroupId: jsii.String("packagingGroupId"),
//   	sourceArn: jsii.String("sourceArn"),
//   	sourceRoleArn: jsii.String("sourceRoleArn"),
//
//   	// the properties below are optional
//   	resourceId: jsii.String("resourceId"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAssetProps struct {
	// Unique identifier that you assign to the asset.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The ID of the packaging group associated with this asset.
	PackagingGroupId *string `field:"required" json:"packagingGroupId" yaml:"packagingGroupId"`
	// The ARN for the source content in Amazon S3.
	SourceArn *string `field:"required" json:"sourceArn" yaml:"sourceArn"`
	// The ARN for the IAM role that provides AWS Elemental MediaPackage access to the Amazon S3 bucket where the source content is stored.
	//
	// Valid format: arn:aws:iam::{accountID}:role/{name}.
	SourceRoleArn *string `field:"required" json:"sourceRoleArn" yaml:"sourceRoleArn"`
	// Unique identifier for this asset, as it's configured in the key provider service.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// The tags to assign to the asset.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

