package previewawsm2mixins


// Properties for CfnDeploymentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDeploymentMixinProps := &CfnDeploymentMixinProps{
//   	ApplicationId: jsii.String("applicationId"),
//   	ApplicationVersion: jsii.Number(123),
//   	EnvironmentId: jsii.String("environmentId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-deployment.html
//
type CfnDeploymentMixinProps struct {
	// The unique identifier of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-deployment.html#cfn-m2-deployment-applicationid
	//
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// The version of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-deployment.html#cfn-m2-deployment-applicationversion
	//
	ApplicationVersion *float64 `field:"optional" json:"applicationVersion" yaml:"applicationVersion"`
	// The unique identifier of the runtime environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-m2-deployment.html#cfn-m2-deployment-environmentid
	//
	EnvironmentId *string `field:"optional" json:"environmentId" yaml:"environmentId"`
}

