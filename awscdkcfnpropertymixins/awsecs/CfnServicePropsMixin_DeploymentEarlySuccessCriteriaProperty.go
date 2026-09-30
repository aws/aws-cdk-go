package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   deploymentEarlySuccessCriteriaProperty := &DeploymentEarlySuccessCriteriaProperty{
//   	Enable: jsii.Boolean(false),
//   	HealthyPercent: jsii.Number(123),
//   	SourceServiceRevisionCleanup: jsii.String("sourceServiceRevisionCleanup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentearlysuccesscriteria.html
//
type CfnServicePropsMixin_DeploymentEarlySuccessCriteriaProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentearlysuccesscriteria.html#cfn-ecs-service-deploymentearlysuccesscriteria-enable
	//
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentearlysuccesscriteria.html#cfn-ecs-service-deploymentearlysuccesscriteria-healthypercent
	//
	HealthyPercent *float64 `field:"optional" json:"healthyPercent" yaml:"healthyPercent"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-deploymentearlysuccesscriteria.html#cfn-ecs-service-deploymentearlysuccesscriteria-sourceservicerevisioncleanup
	//
	SourceServiceRevisionCleanup *string `field:"optional" json:"sourceServiceRevisionCleanup" yaml:"sourceServiceRevisionCleanup"`
}

