package awsserverlessrepo


// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &CfnApplicationProps{
//   	Author: jsii.String("author"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	HomePageUrl: jsii.String("homePageUrl"),
//   	Labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	LicenseBody: jsii.String("licenseBody"),
//   	ReadmeBody: jsii.String("readmeBody"),
//   	SemanticVersion: jsii.String("semanticVersion"),
//   	SourceCodeUrl: jsii.String("sourceCodeUrl"),
//   	SpdxLicenseId: jsii.String("spdxLicenseId"),
//   	TemplateBody: jsii.String("templateBody"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverlessrepo-application.html
//
type CfnApplicationProps struct {
	// The name of the author publishing the app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverlessrepo-application.html#cfn-serverlessrepo-application-author
	//
	Author *string `field:"required" json:"author" yaml:"author"`
	// The description of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverlessrepo-application.html#cfn-serverlessrepo-application-description
	//
	Description *string `field:"required" json:"description" yaml:"description"`
	// The name of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverlessrepo-application.html#cfn-serverlessrepo-application-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A URL with more information about the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverlessrepo-application.html#cfn-serverlessrepo-application-homepageurl
	//
	HomePageUrl *string `field:"optional" json:"homePageUrl" yaml:"homePageUrl"`
	// Labels to improve discovery of apps in search results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverlessrepo-application.html#cfn-serverlessrepo-application-labels
	//
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// A local text file that contains the license of the app.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverlessrepo-application.html#cfn-serverlessrepo-application-licensebody
	//
	LicenseBody *string `field:"optional" json:"licenseBody" yaml:"licenseBody"`
	// A text readme file in Markdown language that contains a more detailed description of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverlessrepo-application.html#cfn-serverlessrepo-application-readmebody
	//
	ReadmeBody *string `field:"optional" json:"readmeBody" yaml:"readmeBody"`
	// The semantic version of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverlessrepo-application.html#cfn-serverlessrepo-application-semanticversion
	//
	SemanticVersion *string `field:"optional" json:"semanticVersion" yaml:"semanticVersion"`
	// A link to a public repository for the source code of your application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverlessrepo-application.html#cfn-serverlessrepo-application-sourcecodeurl
	//
	SourceCodeUrl *string `field:"optional" json:"sourceCodeUrl" yaml:"sourceCodeUrl"`
	// A valid identifier from https://spdx.org/licenses/.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverlessrepo-application.html#cfn-serverlessrepo-application-spdxlicenseid
	//
	SpdxLicenseId *string `field:"optional" json:"spdxLicenseId" yaml:"spdxLicenseId"`
	// The local raw packaged AWS SAM template file of your application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverlessrepo-application.html#cfn-serverlessrepo-application-templatebody
	//
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
}

