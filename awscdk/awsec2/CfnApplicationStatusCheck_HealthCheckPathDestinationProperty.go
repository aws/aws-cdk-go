package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckPathDestinationProperty := &HealthCheckPathDestinationProperty{
//   	SecurityGroupId: jsii.String("securityGroupId"),
//   	SubnetId: jsii.String("subnetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-applicationstatuscheck-healthcheckpathdestination.html
//
type CfnApplicationStatusCheck_HealthCheckPathDestinationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-applicationstatuscheck-healthcheckpathdestination.html#cfn-ec2-applicationstatuscheck-healthcheckpathdestination-securitygroupid
	//
	SecurityGroupId *string `field:"optional" json:"securityGroupId" yaml:"securityGroupId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-applicationstatuscheck-healthcheckpathdestination.html#cfn-ec2-applicationstatuscheck-healthcheckpathdestination-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

