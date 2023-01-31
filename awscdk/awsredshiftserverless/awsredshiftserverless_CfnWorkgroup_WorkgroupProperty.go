package awsredshiftserverless


// The collection of computing resources from which an endpoint is created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workgroupProperty := &workgroupProperty{
//   	baseCapacity: jsii.Number(123),
//   	configParameters: []interface{}{
//   		&configParameterProperty{
//   			parameterKey: jsii.String("parameterKey"),
//   			parameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	creationDate: jsii.String("creationDate"),
//   	endpoint: &endpointProperty{
//   		address: jsii.String("address"),
//   		port: jsii.Number(123),
//   		vpcEndpoints: []interface{}{
//   			&vpcEndpointProperty{
//   				networkInterfaces: []interface{}{
//   					&networkInterfaceProperty{
//   						availabilityZone: jsii.String("availabilityZone"),
//   						networkInterfaceId: jsii.String("networkInterfaceId"),
//   						privateIpAddress: jsii.String("privateIpAddress"),
//   						subnetId: jsii.String("subnetId"),
//   					},
//   				},
//   				vpcEndpointId: jsii.String("vpcEndpointId"),
//   				vpcId: jsii.String("vpcId"),
//   			},
//   		},
//   	},
//   	enhancedVpcRouting: jsii.Boolean(false),
//   	namespaceName: jsii.String("namespaceName"),
//   	publiclyAccessible: jsii.Boolean(false),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	status: jsii.String("status"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	workgroupArn: jsii.String("workgroupArn"),
//   	workgroupId: jsii.String("workgroupId"),
//   	workgroupName: jsii.String("workgroupName"),
//   }
//
type CfnWorkgroup_WorkgroupProperty struct {
	// The base data warehouse capacity of the workgroup in Redshift Processing Units (RPUs).
	BaseCapacity *float64 `field:"optional" json:"baseCapacity" yaml:"baseCapacity"`
	// An array of parameters to set for advanced control over a database.
	//
	// The options are `auto_mv` , `datestyle` , `enable_case_sensitivity_identifier` , `enable_user_activity_logging` , `query_group` , , `search_path` , and query monitoring metrics that let you define performance boundaries. For more information about query monitoring rules and available metrics, see [Query monitoring metrics for Amazon Redshift Serverless](https://docs.aws.amazon.com/redshift/latest/dg/cm-c-wlm-query-monitoring-rules.html#cm-c-wlm-query-monitoring-metrics-serverless) .
	ConfigParameters interface{} `field:"optional" json:"configParameters" yaml:"configParameters"`
	// The creation date of the workgroup.
	CreationDate *string `field:"optional" json:"creationDate" yaml:"creationDate"`
	// The endpoint that is created from the workgroup.
	Endpoint interface{} `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The value that specifies whether to enable enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC.
	EnhancedVpcRouting interface{} `field:"optional" json:"enhancedVpcRouting" yaml:"enhancedVpcRouting"`
	// The namespace the workgroup is associated with.
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
	// A value that specifies whether the workgroup can be accessible from a public network.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// An array of security group IDs to associate with the workgroup.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The status of the workgroup.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// An array of subnet IDs the workgroup is associated with.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The Amazon Resource Name (ARN) that links to the workgroup.
	WorkgroupArn *string `field:"optional" json:"workgroupArn" yaml:"workgroupArn"`
	// The unique identifier of the workgroup.
	WorkgroupId *string `field:"optional" json:"workgroupId" yaml:"workgroupId"`
	// The name of the workgroup.
	WorkgroupName *string `field:"optional" json:"workgroupName" yaml:"workgroupName"`
}

