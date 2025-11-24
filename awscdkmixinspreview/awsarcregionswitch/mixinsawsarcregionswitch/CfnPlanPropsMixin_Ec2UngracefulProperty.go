package mixinsawsarcregionswitch


// Configuration for handling failures when performing operations on EC2 resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ec2UngracefulProperty := &Ec2UngracefulProperty{
//   	MinimumSuccessPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ec2ungraceful.html
//
type CfnPlanPropsMixin_Ec2UngracefulProperty struct {
	// The minimum success percentage that you specify for EC2 Auto Scaling groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-ec2ungraceful.html#cfn-arcregionswitch-plan-ec2ungraceful-minimumsuccesspercentage
	//
	MinimumSuccessPercentage *float64 `field:"optional" json:"minimumSuccessPercentage" yaml:"minimumSuccessPercentage"`
}

