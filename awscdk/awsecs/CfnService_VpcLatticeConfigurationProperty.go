package awsecs


// The VPC Lattice configuration for your service that holds the information for the target group(s) Amazon ECS tasks will be registered to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcLatticeConfigurationProperty := &VpcLatticeConfigurationProperty{
//   	PortName: jsii.String("portName"),
//   	RoleArn: jsii.String("roleArn"),
//   	TargetGroupArn: jsii.String("targetGroupArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-vpclatticeconfiguration.html
//
type CfnService_VpcLatticeConfigurationProperty struct {
	// The name of the port mapping to register in the VPC Lattice target group.
	//
	// This is the name of the `portMapping` you defined in your task definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-vpclatticeconfiguration.html#cfn-ecs-service-vpclatticeconfiguration-portname
	//
	PortName *string `field:"required" json:"portName" yaml:"portName"`
	// The ARN of the IAM role to associate with this VPC Lattice configuration.
	//
	// This is the Amazon ECS infrastructure IAM role that is used to manage your VPC Lattice infrastructure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-vpclatticeconfiguration.html#cfn-ecs-service-vpclatticeconfiguration-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The full Amazon Resource Name (ARN) of the target group or groups associated with the VPC Lattice configuration that the Amazon ECS tasks will be registered to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-vpclatticeconfiguration.html#cfn-ecs-service-vpclatticeconfiguration-targetgrouparn
	//
	TargetGroupArn *string `field:"required" json:"targetGroupArn" yaml:"targetGroupArn"`
}

