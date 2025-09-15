package awsarcregionswitch


// Configuration for an Amazon EC2 Auto Scaling group used in a Region switch plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   asgProperty := &AsgProperty{
//   	Arn: jsii.String("arn"),
//   	CrossAccountRole: jsii.String("crossAccountRole"),
//   	ExternalId: jsii.String("externalId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-asg.html
//
type CfnPlan_AsgProperty struct {
	// The Amazon Resource Name (ARN) of the EC2 Auto Scaling group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-asg.html#cfn-arcregionswitch-plan-asg-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The cross account role for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-asg.html#cfn-arcregionswitch-plan-asg-crossaccountrole
	//
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// The external ID (secret key) for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-asg.html#cfn-arcregionswitch-plan-asg-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
}

