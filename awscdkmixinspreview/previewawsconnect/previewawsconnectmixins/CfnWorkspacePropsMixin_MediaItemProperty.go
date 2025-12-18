package previewawsconnectmixins


// Contains information about a media asset used in a workspace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mediaItemProperty := &MediaItemProperty{
//   	Source: jsii.String("source"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-mediaitem.html
//
type CfnWorkspacePropsMixin_MediaItemProperty struct {
	// The source URL or data for the media asset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-mediaitem.html#cfn-connect-workspace-mediaitem-source
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
	// The type of media.
	//
	// Valid values are: `IMAGE_LOGO_FAVICON` and `IMAGE_LOGO_HORIZONTAL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-mediaitem.html#cfn-connect-workspace-mediaitem-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

