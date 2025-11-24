package mixinsawsarcregionswitch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   healthCheckStateProperty := &HealthCheckStateProperty{
//   	HealthCheckId: jsii.String("healthCheckId"),
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-healthcheckstate.html
//
type CfnPlanPropsMixin_HealthCheckStateProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-healthcheckstate.html#cfn-arcregionswitch-plan-healthcheckstate-healthcheckid
	//
	HealthCheckId *string `field:"optional" json:"healthCheckId" yaml:"healthCheckId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-healthcheckstate.html#cfn-arcregionswitch-plan-healthcheckstate-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}

