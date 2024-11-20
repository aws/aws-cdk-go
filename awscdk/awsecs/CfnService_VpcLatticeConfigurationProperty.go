package awsecs


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-vpclatticeconfiguration.html#cfn-ecs-service-vpclatticeconfiguration-portname
	//
	PortName *string `field:"required" json:"portName" yaml:"portName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-vpclatticeconfiguration.html#cfn-ecs-service-vpclatticeconfiguration-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-vpclatticeconfiguration.html#cfn-ecs-service-vpclatticeconfiguration-targetgrouparn
	//
	TargetGroupArn *string `field:"required" json:"targetGroupArn" yaml:"targetGroupArn"`
}

