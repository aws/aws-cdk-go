package awsredshiftserverless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnWorkgroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkgroupProps := &CfnWorkgroupProps{
//   	WorkgroupName: jsii.String("workgroupName"),
//
//   	// the properties below are optional
//   	BaseCapacity: jsii.Number(123),
//   	ConfigParameters: []interface{}{
//   		&ConfigParameterProperty{
//   			ParameterKey: jsii.String("parameterKey"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	EnhancedVpcRouting: jsii.Boolean(false),
//   	NamespaceName: jsii.String("namespaceName"),
//   	PubliclyAccessible: jsii.Boolean(false),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Workgroup: &WorkgroupProperty{
//   		BaseCapacity: jsii.Number(123),
//   		ConfigParameters: []interface{}{
//   			&ConfigParameterProperty{
//   				ParameterKey: jsii.String("parameterKey"),
//   				ParameterValue: jsii.String("parameterValue"),
//   			},
//   		},
//   		CreationDate: jsii.String("creationDate"),
//   		Endpoint: &EndpointProperty{
//   			Address: jsii.String("address"),
//   			Port: jsii.Number(123),
//   			VpcEndpoints: []interface{}{
//   				&VpcEndpointProperty{
//   					NetworkInterfaces: []interface{}{
//   						&NetworkInterfaceProperty{
//   							AvailabilityZone: jsii.String("availabilityZone"),
//   							NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   							PrivateIpAddress: jsii.String("privateIpAddress"),
//   							SubnetId: jsii.String("subnetId"),
//   						},
//   					},
//   					VpcEndpointId: jsii.String("vpcEndpointId"),
//   					VpcId: jsii.String("vpcId"),
//   				},
//   			},
//   		},
//   		EnhancedVpcRouting: jsii.Boolean(false),
//   		NamespaceName: jsii.String("namespaceName"),
//   		PubliclyAccessible: jsii.Boolean(false),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Status: jsii.String("status"),
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		WorkgroupArn: jsii.String("workgroupArn"),
//   		WorkgroupId: jsii.String("workgroupId"),
//   		WorkgroupName: jsii.String("workgroupName"),
//   	},
//   }
//
type CfnWorkgroupProps struct {
	// The name of the workgroup.
	WorkgroupName *string `field:"required" json:"workgroupName" yaml:"workgroupName"`
	// The base compute capacity of the workgroup in Redshift Processing Units (RPUs).
	BaseCapacity *float64 `field:"optional" json:"baseCapacity" yaml:"baseCapacity"`
	// A list of parameters to set for finer control over a database.
	//
	// Available options are `datestyle` , `enable_user_activity_logging` , `query_group` , `search_path` , and `max_query_execution_time` .
	ConfigParameters interface{} `field:"optional" json:"configParameters" yaml:"configParameters"`
	// The value that specifies whether to enable enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC.
	EnhancedVpcRouting interface{} `field:"optional" json:"enhancedVpcRouting" yaml:"enhancedVpcRouting"`
	// The namespace the workgroup is associated with.
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
	// A value that specifies whether the workgroup can be accessible from a public network.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// A list of security group IDs to associate with the workgroup.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of subnet IDs the workgroup is associated with.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The map of the key-value pairs used to tag the workgroup.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::RedshiftServerless::Workgroup.Workgroup`.
	Workgroup interface{} `field:"optional" json:"workgroup" yaml:"workgroup"`
}

