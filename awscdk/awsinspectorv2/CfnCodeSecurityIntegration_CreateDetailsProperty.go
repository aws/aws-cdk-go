package awsinspectorv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   createDetailsProperty := &CreateDetailsProperty{
//   	GitlabSelfManaged: &CreateGitLabSelfManagedIntegrationDetailProperty{
//   		AccessToken: jsii.String("accessToken"),
//   		InstanceUrl: jsii.String("instanceUrl"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityintegration-createdetails.html
//
type CfnCodeSecurityIntegration_CreateDetailsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-inspectorv2-codesecurityintegration-createdetails.html#cfn-inspectorv2-codesecurityintegration-createdetails-gitlabselfmanaged
	//
	GitlabSelfManaged interface{} `field:"required" json:"gitlabSelfManaged" yaml:"gitlabSelfManaged"`
}

