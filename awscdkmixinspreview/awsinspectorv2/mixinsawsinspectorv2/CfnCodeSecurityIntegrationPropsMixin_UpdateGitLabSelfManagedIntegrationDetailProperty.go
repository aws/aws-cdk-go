package mixinsawsinspectorv2


// Contains details required to update an integration with a self-managed GitLab instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   updateGitLabSelfManagedIntegrationDetailProperty := &UpdateGitLabSelfManagedIntegrationDetailProperty{
//   	AuthCode: jsii.String("authCode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityintegration-updategitlabselfmanagedintegrationdetail.html
//
type CfnCodeSecurityIntegrationPropsMixin_UpdateGitLabSelfManagedIntegrationDetailProperty struct {
	// The authorization code received from the self-managed GitLab instance to update the integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityintegration-updategitlabselfmanagedintegrationdetail.html#cfn-inspectorv2-codesecurityintegration-updategitlabselfmanagedintegrationdetail-authcode
	//
	AuthCode *string `field:"optional" json:"authCode" yaml:"authCode"`
}

