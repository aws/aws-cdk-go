package mixinsawsrtbfabric

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnResponderGatewayPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResponderGatewayMixinProps := &CfnResponderGatewayMixinProps{
//   	Description: jsii.String("description"),
//   	DomainName: jsii.String("domainName"),
//   	ManagedEndpointConfiguration: &ManagedEndpointConfigurationProperty{
//   		AutoScalingGroupsConfiguration: &AutoScalingGroupsConfigurationProperty{
//   			AutoScalingGroupNameList: []*string{
//   				jsii.String("autoScalingGroupNameList"),
//   			},
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   		EksEndpointsConfiguration: &EksEndpointsConfigurationProperty{
//   			ClusterApiServerCaCertificateChain: jsii.String("clusterApiServerCaCertificateChain"),
//   			ClusterApiServerEndpointUri: jsii.String("clusterApiServerEndpointUri"),
//   			ClusterName: jsii.String("clusterName"),
//   			EndpointsResourceName: jsii.String("endpointsResourceName"),
//   			EndpointsResourceNamespace: jsii.String("endpointsResourceNamespace"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   	},
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrustStoreConfiguration: &TrustStoreConfigurationProperty{
//   		CertificateAuthorityCertificates: []*string{
//   			jsii.String("certificateAuthorityCertificates"),
//   		},
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-respondergateway.html
//
type CfnResponderGatewayMixinProps struct {
	// An optional description for the responder gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-respondergateway.html#cfn-rtbfabric-respondergateway-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The domain name for the responder gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-respondergateway.html#cfn-rtbfabric-respondergateway-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The configuration for the managed endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-respondergateway.html#cfn-rtbfabric-respondergateway-managedendpointconfiguration
	//
	ManagedEndpointConfiguration interface{} `field:"optional" json:"managedEndpointConfiguration" yaml:"managedEndpointConfiguration"`
	// The networking port to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-respondergateway.html#cfn-rtbfabric-respondergateway-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The networking protocol to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-respondergateway.html#cfn-rtbfabric-respondergateway-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The unique identifiers of the security groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-respondergateway.html#cfn-rtbfabric-respondergateway-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The unique identifiers of the subnets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-respondergateway.html#cfn-rtbfabric-respondergateway-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// A map of the key-value pairs of the tag or tags to assign to the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-respondergateway.html#cfn-rtbfabric-respondergateway-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The configuration of the trust store.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-respondergateway.html#cfn-rtbfabric-respondergateway-truststoreconfiguration
	//
	TrustStoreConfiguration interface{} `field:"optional" json:"trustStoreConfiguration" yaml:"trustStoreConfiguration"`
	// The unique identifier of the Virtual Private Cloud (VPC).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-respondergateway.html#cfn-rtbfabric-respondergateway-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

