package awsconnect


// Properties for defining a `CfnViewVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnViewVersionProps := &CfnViewVersionProps{
//   	ViewArn: jsii.String("viewArn"),
//
//   	// the properties below are optional
//   	VersionDescription: jsii.String("versionDescription"),
//   	ViewContentSha256: jsii.String("viewContentSha256"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-viewversion.html
//
type CfnViewVersionProps struct {
	// The unqualified Amazon Resource Name (ARN) of the view.
	//
	// For example:
	//
	// `arn:<partition>:connect:<region>:<accountId>:instance/00000000-0000-0000-0000-000000000000/view/00000000-0000-0000-0000-000000000000`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-viewversion.html#cfn-connect-viewversion-viewarn
	//
	ViewArn *string `field:"required" json:"viewArn" yaml:"viewArn"`
	// The description of the view version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-viewversion.html#cfn-connect-viewversion-versiondescription
	//
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
	// Indicates the checksum value of the latest published view content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-viewversion.html#cfn-connect-viewversion-viewcontentsha256
	//
	ViewContentSha256 *string `field:"optional" json:"viewContentSha256" yaml:"viewContentSha256"`
}

