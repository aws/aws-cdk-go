package awsredshiftserverless


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
	// `CfnWorkgroup.WorkgroupProperty.BaseCapacity`.
	BaseCapacity *float64 `field:"optional" json:"baseCapacity" yaml:"baseCapacity"`
	// `CfnWorkgroup.WorkgroupProperty.ConfigParameters`.
	ConfigParameters interface{} `field:"optional" json:"configParameters" yaml:"configParameters"`
	// `CfnWorkgroup.WorkgroupProperty.CreationDate`.
	CreationDate *string `field:"optional" json:"creationDate" yaml:"creationDate"`
	// `CfnWorkgroup.WorkgroupProperty.Endpoint`.
	Endpoint interface{} `field:"optional" json:"endpoint" yaml:"endpoint"`
	// `CfnWorkgroup.WorkgroupProperty.EnhancedVpcRouting`.
	EnhancedVpcRouting interface{} `field:"optional" json:"enhancedVpcRouting" yaml:"enhancedVpcRouting"`
	// `CfnWorkgroup.WorkgroupProperty.NamespaceName`.
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
	// `CfnWorkgroup.WorkgroupProperty.PubliclyAccessible`.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// `CfnWorkgroup.WorkgroupProperty.SecurityGroupIds`.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// `CfnWorkgroup.WorkgroupProperty.Status`.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// `CfnWorkgroup.WorkgroupProperty.SubnetIds`.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// `CfnWorkgroup.WorkgroupProperty.WorkgroupArn`.
	WorkgroupArn *string `field:"optional" json:"workgroupArn" yaml:"workgroupArn"`
	// `CfnWorkgroup.WorkgroupProperty.WorkgroupId`.
	WorkgroupId *string `field:"optional" json:"workgroupId" yaml:"workgroupId"`
	// `CfnWorkgroup.WorkgroupProperty.WorkgroupName`.
	WorkgroupName *string `field:"optional" json:"workgroupName" yaml:"workgroupName"`
}

