package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnHealthCheckSourcePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnHealthCheckSourceMixinProps := &CfnHealthCheckSourceMixinProps{
//   	HealthCheckSourceConfig: &HealthCheckSourceConfigProperty{
//   		SubnetConfigurationList: []interface{}{
//   			&SubnetConfigurationProperty{
//   				AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   				SecurityGroupIdList: []*string{
//   					jsii.String("securityGroupIdList"),
//   				},
//   				SubnetId: jsii.String("subnetId"),
//   			},
//   		},
//   		VpcId: jsii.String("vpcId"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthchecksource.html
//
type CfnHealthCheckSourceMixinProps struct {
	// <p>A complex type that contains the VPC, subnets, and security groups for a health check source.
	//
	// You can't change these values after you create the source.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthchecksource.html#cfn-route53-healthchecksource-healthchecksourceconfig
	//
	HealthCheckSourceConfig interface{} `field:"optional" json:"healthCheckSourceConfig" yaml:"healthCheckSourceConfig"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthchecksource.html#cfn-route53-healthchecksource-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

