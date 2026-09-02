package awsdevopsagent


// A single file inside an Asset's bundle.
//
// Path is the diff key on update; Content is write-only and not repopulated by Read.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var metadata interface{}
//
//   assetFileProperty := &AssetFileProperty{
//   	Path: jsii.String("path"),
//
//   	// the properties below are optional
//   	ContentBytes: jsii.String("contentBytes"),
//   	ContentText: jsii.String("contentText"),
//   	Metadata: metadata,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-asset-assetfile.html
//
type CfnAsset_AssetFileProperty struct {
	// Path of this file within the asset bundle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-asset-assetfile.html#cfn-devopsagent-asset-assetfile-path
	//
	Path *string `field:"required" json:"path" yaml:"path"`
	// Base64-encoded binary contents of the file.
	//
	// Mutually exclusive with ContentText (max 6 MiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-asset-assetfile.html#cfn-devopsagent-asset-assetfile-contentbytes
	//
	ContentBytes *string `field:"optional" json:"contentBytes" yaml:"contentBytes"`
	// UTF-8 text contents of the file.
	//
	// Mutually exclusive with ContentBytes (max 1.5 MiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-asset-assetfile.html#cfn-devopsagent-asset-assetfile-contenttext
	//
	ContentText *string `field:"optional" json:"contentText" yaml:"contentText"`
	// Per-file metadata document.
	//
	// Values may be strings, numbers, booleans, or lists of any of those (validated server-side).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-asset-assetfile.html#cfn-devopsagent-asset-assetfile-metadata
	//
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
}

