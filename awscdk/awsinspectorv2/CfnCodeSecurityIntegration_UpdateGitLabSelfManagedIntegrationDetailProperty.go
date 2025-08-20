package awsinspectorv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   updateGitLabSelfManagedIntegrationDetailProperty := &UpdateGitLabSelfManagedIntegrationDetailProperty{
//   	AuthCode: jsii.String("authCode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityintegration-updategitlabselfmanagedintegrationdetail.html
//
type CfnCodeSecurityIntegration_UpdateGitLabSelfManagedIntegrationDetailProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityintegration-updategitlabselfmanagedintegrationdetail.html#cfn-inspectorv2-codesecurityintegration-updategitlabselfmanagedintegrationdetail-authcode
	//
	AuthCode *string `field:"required" json:"authCode" yaml:"authCode"`
}

