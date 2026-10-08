package awsnetworksecuritymanager


// Configuration settings for the deployment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deploymentConfigurationProperty := &DeploymentConfigurationProperty{
//   	EnableCrossAccountVisibility: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networksecuritymanager-deployment-deploymentconfiguration.html
//
type CfnDeployment_DeploymentConfigurationProperty struct {
	// Whether cross-account visibility is enabled for the deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networksecuritymanager-deployment-deploymentconfiguration.html#cfn-networksecuritymanager-deployment-deploymentconfiguration-enablecrossaccountvisibility
	//
	EnableCrossAccountVisibility interface{} `field:"required" json:"enableCrossAccountVisibility" yaml:"enableCrossAccountVisibility"`
}

