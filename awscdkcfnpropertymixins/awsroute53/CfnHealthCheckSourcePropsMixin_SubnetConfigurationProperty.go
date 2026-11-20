package awsroute53


// <p>A complex type that identifies one subnet for a health check source and the security groups that Route 53 associates with the network interfaces that it creates in that subnet.</p>.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   subnetConfigurationProperty := &SubnetConfigurationProperty{
//   	AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   	SecurityGroupIdList: []*string{
//   		jsii.String("securityGroupIdList"),
//   	},
//   	SubnetId: jsii.String("subnetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthchecksource-subnetconfiguration.html
//
type CfnHealthCheckSourcePropsMixin_SubnetConfigurationProperty struct {
	// <p>The ID of the Availability Zone that contains the subnet, for example <code>use1-az1</code>.
	//
	// Specify the Availability Zone ID, not the Availability Zone name. Each subnet in a health check source must be in a different Availability Zone.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthchecksource-subnetconfiguration.html#cfn-route53-healthchecksource-subnetconfiguration-availabilityzoneid
	//
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// <p>The IDs of one to ten security groups in the same VPC to associate with the network interfaces that Route 53 creates in this subnet.
	//
	// The security groups must allow outbound traffic to the endpoints that the health checks monitor.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthchecksource-subnetconfiguration.html#cfn-route53-healthchecksource-subnetconfiguration-securitygroupidlist
	//
	SecurityGroupIdList *[]*string `field:"optional" json:"securityGroupIdList" yaml:"securityGroupIdList"`
	// <p>The ID of the subnet.
	//
	// The subnet must have free IP addresses for the network interfaces that Route 53 creates.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthchecksource-subnetconfiguration.html#cfn-route53-healthchecksource-subnetconfiguration-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

