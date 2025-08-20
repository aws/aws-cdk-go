package awsinspectorv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   updateGitHubIntegrationDetailProperty := &UpdateGitHubIntegrationDetailProperty{
//   	Code: jsii.String("code"),
//   	InstallationId: jsii.String("installationId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityintegration-updategithubintegrationdetail.html
//
type CfnCodeSecurityIntegration_UpdateGitHubIntegrationDetailProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityintegration-updategithubintegrationdetail.html#cfn-inspectorv2-codesecurityintegration-updategithubintegrationdetail-code
	//
	Code *string `field:"required" json:"code" yaml:"code"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityintegration-updategithubintegrationdetail.html#cfn-inspectorv2-codesecurityintegration-updategithubintegrationdetail-installationid
	//
	InstallationId *string `field:"required" json:"installationId" yaml:"installationId"`
}

