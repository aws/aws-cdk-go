package awsappconfig


// Properties for defining a `CfnHostedConfigurationVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnHostedConfigurationVersionProps := &cfnHostedConfigurationVersionProps{
//   	applicationId: jsii.String("applicationId"),
//   	configurationProfileId: jsii.String("configurationProfileId"),
//   	content: jsii.String("content"),
//   	contentType: jsii.String("contentType"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	latestVersionNumber: jsii.Number(123),
//   }
//
type CfnHostedConfigurationVersionProps struct {
	// The application ID.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The configuration profile ID.
	ConfigurationProfileId *string `field:"required" json:"configurationProfileId" yaml:"configurationProfileId"`
	// The content of the configuration or the configuration data.
	Content *string `field:"required" json:"content" yaml:"content"`
	// A standard MIME type describing the format of the configuration content.
	//
	// For more information, see [Content-Type](https://docs.aws.amazon.com/https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17) .
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// A description of the configuration.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An optional locking token used to prevent race conditions from overwriting configuration updates when creating a new version.
	//
	// To ensure your data is not overwritten when creating multiple hosted configuration versions in rapid succession, specify the version number of the latest hosted configuration version.
	LatestVersionNumber *float64 `field:"optional" json:"latestVersionNumber" yaml:"latestVersionNumber"`
}

