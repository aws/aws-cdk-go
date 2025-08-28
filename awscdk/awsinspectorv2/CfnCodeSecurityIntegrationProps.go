package awsinspectorv2


// Properties for defining a `CfnCodeSecurityIntegration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCodeSecurityIntegrationProps := &CfnCodeSecurityIntegrationProps{
//   	CreateIntegrationDetails: &CreateDetailsProperty{
//   		GitlabSelfManaged: &CreateGitLabSelfManagedIntegrationDetailProperty{
//   			AccessToken: jsii.String("accessToken"),
//   			InstanceUrl: jsii.String("instanceUrl"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Type: jsii.String("type"),
//   	UpdateIntegrationDetails: &UpdateDetailsProperty{
//   		Github: &UpdateGitHubIntegrationDetailProperty{
//   			Code: jsii.String("code"),
//   			InstallationId: jsii.String("installationId"),
//   		},
//   		GitlabSelfManaged: &UpdateGitLabSelfManagedIntegrationDetailProperty{
//   			AuthCode: jsii.String("authCode"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-codesecurityintegration.html
//
type CfnCodeSecurityIntegrationProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-codesecurityintegration.html#cfn-inspectorv2-codesecurityintegration-createintegrationdetails
	//
	CreateIntegrationDetails interface{} `field:"optional" json:"createIntegrationDetails" yaml:"createIntegrationDetails"`
	// Code Security Integration name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-codesecurityintegration.html#cfn-inspectorv2-codesecurityintegration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-codesecurityintegration.html#cfn-inspectorv2-codesecurityintegration-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-codesecurityintegration.html#cfn-inspectorv2-codesecurityintegration-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspectorv2-codesecurityintegration.html#cfn-inspectorv2-codesecurityintegration-updateintegrationdetails
	//
	UpdateIntegrationDetails interface{} `field:"optional" json:"updateIntegrationDetails" yaml:"updateIntegrationDetails"`
}

