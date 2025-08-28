package awscodedeploy


// Information about the minimum number of healthy instances per Availability Zone.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   minimumHealthyHostsPerZoneProperty := &MinimumHealthyHostsPerZoneProperty{
//   	Type: jsii.String("type"),
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-minimumhealthyhostsperzone.html
//
type CfnDeploymentConfig_MinimumHealthyHostsPerZoneProperty struct {
	// The `type` associated with the `MinimumHealthyHostsPerZone` option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-minimumhealthyhostsperzone.html#cfn-codedeploy-deploymentconfig-minimumhealthyhostsperzone-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The `value` associated with the `MinimumHealthyHostsPerZone` option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-minimumhealthyhostsperzone.html#cfn-codedeploy-deploymentconfig-minimumhealthyhostsperzone-value
	//
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

