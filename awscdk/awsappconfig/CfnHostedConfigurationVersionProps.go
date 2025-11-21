package awsappconfig


// Properties for defining a `CfnHostedConfigurationVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnHostedConfigurationVersionProps := &CfnHostedConfigurationVersionProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	ConfigurationProfileId: jsii.String("configurationProfileId"),
//   	Content: jsii.String("content"),
//   	ContentType: jsii.String("contentType"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	LatestVersionNumber: jsii.Number(123),
//   	VersionLabel: jsii.String("versionLabel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-hostedconfigurationversion.html
//
type CfnHostedConfigurationVersionProps struct {
	// The application ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-hostedconfigurationversion.html#cfn-appconfig-hostedconfigurationversion-applicationid
	//
	ApplicationId interface{} `field:"required" json:"applicationId" yaml:"applicationId"`
	// The configuration profile ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-hostedconfigurationversion.html#cfn-appconfig-hostedconfigurationversion-configurationprofileid
	//
	ConfigurationProfileId interface{} `field:"required" json:"configurationProfileId" yaml:"configurationProfileId"`
	// The configuration data, as bytes.
	//
	// > AWS AppConfig accepts any type of data, including text formats like JSON or TOML, or binary formats like protocol buffers or compressed data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-hostedconfigurationversion.html#cfn-appconfig-hostedconfigurationversion-content
	//
	Content *string `field:"required" json:"content" yaml:"content"`
	// A standard MIME type describing the format of the configuration content.
	//
	// For more information, see [Content-Type](https://docs.aws.amazon.com/https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-hostedconfigurationversion.html#cfn-appconfig-hostedconfigurationversion-contenttype
	//
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// A description of the configuration.
	//
	// > Due to HTTP limitations, this field only supports ASCII characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-hostedconfigurationversion.html#cfn-appconfig-hostedconfigurationversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An optional locking token used to prevent race conditions from overwriting configuration updates when creating a new version.
	//
	// To ensure your data is not overwritten when creating multiple hosted configuration versions in rapid succession, specify the version number of the latest hosted configuration version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-hostedconfigurationversion.html#cfn-appconfig-hostedconfigurationversion-latestversionnumber
	//
	LatestVersionNumber *float64 `field:"optional" json:"latestVersionNumber" yaml:"latestVersionNumber"`
	// A user-defined label for an AWS AppConfig hosted configuration version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-hostedconfigurationversion.html#cfn-appconfig-hostedconfigurationversion-versionlabel
	//
	VersionLabel *string `field:"optional" json:"versionLabel" yaml:"versionLabel"`
}

