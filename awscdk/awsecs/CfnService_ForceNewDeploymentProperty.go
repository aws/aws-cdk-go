package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   forceNewDeploymentProperty := &ForceNewDeploymentProperty{
//   	EnableForceNewDeployment: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	ForceNewDeploymentNonce: jsii.String("forceNewDeploymentNonce"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-forcenewdeployment.html
//
type CfnService_ForceNewDeploymentProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-forcenewdeployment.html#cfn-ecs-service-forcenewdeployment-enableforcenewdeployment
	//
	EnableForceNewDeployment interface{} `field:"required" json:"enableForceNewDeployment" yaml:"enableForceNewDeployment"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-forcenewdeployment.html#cfn-ecs-service-forcenewdeployment-forcenewdeploymentnonce
	//
	ForceNewDeploymentNonce *string `field:"optional" json:"forceNewDeploymentNonce" yaml:"forceNewDeploymentNonce"`
}

