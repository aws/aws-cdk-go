package previewawsmediapackagemixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAssetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAssetMixinProps := &CfnAssetMixinProps{
//   	EgressEndpoints: []interface{}{
//   		&EgressEndpointProperty{
//   			PackagingConfigurationId: jsii.String("packagingConfigurationId"),
//   			Url: jsii.String("url"),
//   		},
//   	},
//   	Id: jsii.String("id"),
//   	PackagingGroupId: jsii.String("packagingGroupId"),
//   	ResourceId: jsii.String("resourceId"),
//   	SourceArn: jsii.String("sourceArn"),
//   	SourceRoleArn: jsii.String("sourceRoleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html
//
type CfnAssetMixinProps struct {
	// List of playback endpoints that are available for this asset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-egressendpoints
	//
	EgressEndpoints interface{} `field:"optional" json:"egressEndpoints" yaml:"egressEndpoints"`
	// Unique identifier that you assign to the asset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the packaging group associated with this asset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-packaginggroupid
	//
	PackagingGroupId *string `field:"optional" json:"packagingGroupId" yaml:"packagingGroupId"`
	// Unique identifier for this asset, as it's configured in the key provider service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-resourceid
	//
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// The ARN for the source content in Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-sourcearn
	//
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
	// The ARN for the IAM role that provides AWS Elemental MediaPackage access to the Amazon S3 bucket where the source content is stored.
	//
	// Valid format: arn:aws:iam::{accountID}:role/{name}.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-sourcerolearn
	//
	SourceRoleArn *string `field:"optional" json:"sourceRoleArn" yaml:"sourceRoleArn"`
	// The tags to assign to the asset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-asset.html#cfn-mediapackage-asset-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

