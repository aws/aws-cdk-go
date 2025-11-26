package previewawsredshiftserverlessmixins


// The collection of computing resources from which an endpoint is created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   workgroupProperty := &WorkgroupProperty{
//   	BaseCapacity: jsii.Number(123),
//   	ConfigParameters: []interface{}{
//   		&ConfigParameterProperty{
//   			ParameterKey: jsii.String("parameterKey"),
//   			ParameterValue: jsii.String("parameterValue"),
//   		},
//   	},
//   	CreationDate: jsii.String("creationDate"),
//   	Endpoint: &EndpointProperty{
//   		Address: jsii.String("address"),
//   		Port: jsii.Number(123),
//   		VpcEndpoints: []interface{}{
//   			&VpcEndpointProperty{
//   				NetworkInterfaces: []interface{}{
//   					&NetworkInterfaceProperty{
//   						AvailabilityZone: jsii.String("availabilityZone"),
//   						NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   						PrivateIpAddress: jsii.String("privateIpAddress"),
//   						SubnetId: jsii.String("subnetId"),
//   					},
//   				},
//   				VpcEndpointId: jsii.String("vpcEndpointId"),
//   				VpcId: jsii.String("vpcId"),
//   			},
//   		},
//   	},
//   	EnhancedVpcRouting: jsii.Boolean(false),
//   	MaxCapacity: jsii.Number(123),
//   	NamespaceName: jsii.String("namespaceName"),
//   	PricePerformanceTarget: &PerformanceTargetProperty{
//   		Level: jsii.Number(123),
//   		Status: jsii.String("status"),
//   	},
//   	PubliclyAccessible: jsii.Boolean(false),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	Status: jsii.String("status"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	TrackName: jsii.String("trackName"),
//   	WorkgroupArn: jsii.String("workgroupArn"),
//   	WorkgroupId: jsii.String("workgroupId"),
//   	WorkgroupName: jsii.String("workgroupName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html
//
type CfnWorkgroupPropsMixin_WorkgroupProperty struct {
	// The base data warehouse capacity of the workgroup in Redshift Processing Units (RPUs).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-basecapacity
	//
	BaseCapacity *float64 `field:"optional" json:"baseCapacity" yaml:"baseCapacity"`
	// An array of parameters to set for advanced control over a database.
	//
	// The options are `auto_mv` , `datestyle` , `enable_case_sensitive_identifier` , `enable_user_activity_logging` , `query_group` , `search_path` , `require_ssl` , `use_fips_ssl` , and query monitoring metrics that let you define performance boundaries. For more information about query monitoring rules and available metrics, see [Query monitoring metrics for Amazon Redshift Serverless](https://docs.aws.amazon.com/redshift/latest/dg/cm-c-wlm-query-monitoring-rules.html#cm-c-wlm-query-monitoring-metrics-serverless) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-configparameters
	//
	ConfigParameters interface{} `field:"optional" json:"configParameters" yaml:"configParameters"`
	// The creation date of the workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-creationdate
	//
	CreationDate *string `field:"optional" json:"creationDate" yaml:"creationDate"`
	// The endpoint that is created from the workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-endpoint
	//
	Endpoint interface{} `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The value that specifies whether to enable enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-enhancedvpcrouting
	//
	EnhancedVpcRouting interface{} `field:"optional" json:"enhancedVpcRouting" yaml:"enhancedVpcRouting"`
	// The maximum data-warehouse capacity Amazon Redshift Serverless uses to serve queries.
	//
	// The max capacity is specified in RPUs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-maxcapacity
	//
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The namespace the workgroup is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-namespacename
	//
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
	// An object that represents the price performance target settings for the workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-priceperformancetarget
	//
	PricePerformanceTarget interface{} `field:"optional" json:"pricePerformanceTarget" yaml:"pricePerformanceTarget"`
	// A value that specifies whether the workgroup can be accessible from a public network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-publiclyaccessible
	//
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// An array of security group IDs to associate with the workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The status of the workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// An array of subnet IDs the workgroup is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The name of the track for the workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-trackname
	//
	TrackName *string `field:"optional" json:"trackName" yaml:"trackName"`
	// The Amazon Resource Name (ARN) that links to the workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-workgrouparn
	//
	WorkgroupArn *string `field:"optional" json:"workgroupArn" yaml:"workgroupArn"`
	// The unique identifier of the workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-workgroupid
	//
	WorkgroupId *string `field:"optional" json:"workgroupId" yaml:"workgroupId"`
	// The name of the workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-workgroup.html#cfn-redshiftserverless-workgroup-workgroup-workgroupname
	//
	WorkgroupName *string `field:"optional" json:"workgroupName" yaml:"workgroupName"`
}

