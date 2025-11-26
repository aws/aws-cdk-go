package previewawscodedeploymixins


// Information about the type of deployment, either in-place or blue/green, you want to run and whether to route deployment traffic behind a load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   deploymentStyleProperty := &DeploymentStyleProperty{
//   	DeploymentOption: jsii.String("deploymentOption"),
//   	DeploymentType: jsii.String("deploymentType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deploymentstyle.html
//
type CfnDeploymentGroupPropsMixin_DeploymentStyleProperty struct {
	// Indicates whether to route deployment traffic behind a load balancer.
	//
	// > An Amazon EC2 Application Load Balancer or Network Load Balancer is required for an Amazon ECS deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deploymentstyle.html#cfn-codedeploy-deploymentgroup-deploymentstyle-deploymentoption
	//
	DeploymentOption *string `field:"optional" json:"deploymentOption" yaml:"deploymentOption"`
	// Indicates whether to run an in-place or blue/green deployment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deploymentstyle.html#cfn-codedeploy-deploymentgroup-deploymentstyle-deploymenttype
	//
	DeploymentType *string `field:"optional" json:"deploymentType" yaml:"deploymentType"`
}

