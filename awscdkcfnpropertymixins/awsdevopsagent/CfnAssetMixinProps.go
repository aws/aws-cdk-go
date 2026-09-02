package awsdevopsagent


// Properties for CfnAssetPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var metadata interface{}
//
//   cfnAssetMixinProps := &CfnAssetMixinProps{
//   	AgentSpaceId: jsii.String("agentSpaceId"),
//   	AssetType: jsii.String("assetType"),
//   	Files: []interface{}{
//   		&AssetFileProperty{
//   			ContentBytes: jsii.String("contentBytes"),
//   			ContentText: jsii.String("contentText"),
//   			Metadata: metadata,
//   			Path: jsii.String("path"),
//   		},
//   	},
//   	Metadata: metadata,
//   	Zip: jsii.String("zip"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-asset.html
//
type CfnAssetMixinProps struct {
	// The unique identifier of the parent Agent Space.
	//
	// The asset is created as a child of this agent space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-asset.html#cfn-devopsagent-asset-agentspaceid
	//
	AgentSpaceId *string `field:"optional" json:"agentSpaceId" yaml:"agentSpaceId"`
	// The type of asset.
	//
	// The Asset API treats this as an open string; call ListAssetTypes for the current authoritative set of supported types. As of launch, customer-creatable types include skill, agents_md, and attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-asset.html#cfn-devopsagent-asset-assettype
	//
	AssetType *string `field:"optional" json:"assetType" yaml:"assetType"`
	// Inline file list.
	//
	// Mutually exclusive with Zip; enforced by the handler at Create/Update time. Write-only: not repopulated by Read.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-asset.html#cfn-devopsagent-asset-files
	//
	Files interface{} `field:"optional" json:"files" yaml:"files"`
	// Asset metadata document.
	//
	// Required and optional keys depend on AssetType. Values may be strings, numbers, booleans, or lists of any of those - validated server-side; see the public Asset API docs for the per-type metadata schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-asset.html#cfn-devopsagent-asset-metadata
	//
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// Base64-encoded zip bundle containing all files for the asset.
	//
	// Mutually exclusive with Files; enforced by the handler at Create/Update time. Write-only: not repopulated by Read. Server treats a zip as 'replace all files' (max 6 MiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-asset.html#cfn-devopsagent-asset-zip
	//
	Zip *string `field:"optional" json:"zip" yaml:"zip"`
}

