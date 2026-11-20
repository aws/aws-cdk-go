package awsroute53


// <p>A complex type that contains the VPC, subnets, and security groups for a health check source.
//
// You can't change these values after you create the source.</p>
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   healthCheckSourceConfigProperty := &HealthCheckSourceConfigProperty{
//   	SubnetConfigurationList: []interface{}{
//   		&SubnetConfigurationProperty{
//   			AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   			SecurityGroupIdList: []*string{
//   				jsii.String("securityGroupIdList"),
//   			},
//   			SubnetId: jsii.String("subnetId"),
//   		},
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthchecksource-healthchecksourceconfig.html
//
type CfnHealthCheckSource_HealthCheckSourceConfigProperty struct {
	// <p>A complex type that contains one <code>SubnetConfiguration</code> element for each Availability Zone that Route 53 uses for private endpoint health checks in the Amazon Web Services Region.
	//
	// Each element must be in a different Availability Zone. Most Regions require three elements. Regions that have only two such Availability Zones, such as US West (N. California), require two elements. Route 53 creates network interfaces in each subnet and sends health check requests from all of them.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthchecksource-healthchecksourceconfig.html#cfn-route53-healthchecksource-healthchecksourceconfig-subnetconfigurationlist
	//
	SubnetConfigurationList interface{} `field:"required" json:"subnetConfigurationList" yaml:"subnetConfigurationList"`
	// <p>The ID of the VPC that contains the subnets for the health check source.
	//
	// All subnets in <code>SubnetConfigurationList</code> must belong to this VPC.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthchecksource-healthchecksourceconfig.html#cfn-route53-healthchecksource-healthchecksourceconfig-vpcid
	//
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

