package awsinspectorv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   updateDetailsProperty := &UpdateDetailsProperty{
//   	Github: &UpdateGitHubIntegrationDetailProperty{
//   		Code: jsii.String("code"),
//   		InstallationId: jsii.String("installationId"),
//   	},
//   	GitlabSelfManaged: &UpdateGitLabSelfManagedIntegrationDetailProperty{
//   		AuthCode: jsii.String("authCode"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityintegration-updatedetails.html
//
type CfnCodeSecurityIntegration_UpdateDetailsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityintegration-updatedetails.html#cfn-inspectorv2-codesecurityintegration-updatedetails-github
	//
	Github interface{} `field:"optional" json:"github" yaml:"github"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityintegration-updatedetails.html#cfn-inspectorv2-codesecurityintegration-updatedetails-gitlabselfmanaged
	//
	GitlabSelfManaged interface{} `field:"optional" json:"gitlabSelfManaged" yaml:"gitlabSelfManaged"`
}

