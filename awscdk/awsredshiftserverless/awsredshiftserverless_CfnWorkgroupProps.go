package awsredshiftserverless

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
//   	Port: jsii.Number(123),
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
	// The custom port to use when connecting to a workgroup.
	//
	// Valid port ranges are 5431-5455 and 8191-8215. The default is 5439.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A value that specifies whether the workgroup can be accessible from a public network.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// A list of security group IDs to associate with the workgroup.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of subnet IDs the workgroup is associated with.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The map of the key-value pairs used to tag the workgroup.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

