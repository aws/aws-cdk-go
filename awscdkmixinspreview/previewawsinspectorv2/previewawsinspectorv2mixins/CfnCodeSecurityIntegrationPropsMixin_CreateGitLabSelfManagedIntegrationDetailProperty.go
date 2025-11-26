package previewawsinspectorv2mixins


// Contains details required to create an integration with a self-managed GitLab instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   createGitLabSelfManagedIntegrationDetailProperty := &CreateGitLabSelfManagedIntegrationDetailProperty{
//   	AccessToken: jsii.String("accessToken"),
//   	InstanceUrl: jsii.String("instanceUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityintegration-creategitlabselfmanagedintegrationdetail.html
//
type CfnCodeSecurityIntegrationPropsMixin_CreateGitLabSelfManagedIntegrationDetailProperty struct {
	// The personal access token used to authenticate with the self-managed GitLab instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityintegration-creategitlabselfmanagedintegrationdetail.html#cfn-inspectorv2-codesecurityintegration-creategitlabselfmanagedintegrationdetail-accesstoken
	//
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// The URL of the self-managed GitLab instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityintegration-creategitlabselfmanagedintegrationdetail.html#cfn-inspectorv2-codesecurityintegration-creategitlabselfmanagedintegrationdetail-instanceurl
	//
	InstanceUrl *string `field:"optional" json:"instanceUrl" yaml:"instanceUrl"`
}

