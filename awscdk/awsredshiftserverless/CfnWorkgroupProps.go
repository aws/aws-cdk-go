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
//   	MaxCapacity: jsii.Number(123),
//   	NamespaceName: jsii.String("namespaceName"),
//   	Port: jsii.Number(123),
//   	PricePerformanceTarget: &PerformanceTargetProperty{
//   		Level: jsii.Number(123),
//   		Status: jsii.String("status"),
//   	},
//   	PubliclyAccessible: jsii.Boolean(false),
//   	RecoveryPointId: jsii.String("recoveryPointId"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SnapshotArn: jsii.String("snapshotArn"),
//   	SnapshotName: jsii.String("snapshotName"),
//   	SnapshotOwnerAccount: jsii.String("snapshotOwnerAccount"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TrackName: jsii.String("trackName"),
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
//   		MaxCapacity: jsii.Number(123),
//   		NamespaceName: jsii.String("namespaceName"),
//   		PricePerformanceTarget: &PerformanceTargetProperty{
//   			Level: jsii.Number(123),
//   			Status: jsii.String("status"),
//   		},
//   		PubliclyAccessible: jsii.Boolean(false),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Status: jsii.String("status"),
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		TrackName: jsii.String("trackName"),
//   		WorkgroupArn: jsii.String("workgroupArn"),
//   		WorkgroupId: jsii.String("workgroupId"),
//   		WorkgroupName: jsii.String("workgroupName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html
//
type CfnWorkgroupProps struct {
	// The name of the workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html#cfn-redshiftserverless-workgroup-workgroupname
	//
	WorkgroupName *string `field:"required" json:"workgroupName" yaml:"workgroupName"`
	// The base compute capacity of the workgroup in Redshift Processing Units (RPUs).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html#cfn-redshiftserverless-workgroup-basecapacity
	//
	BaseCapacity *float64 `field:"optional" json:"baseCapacity" yaml:"baseCapacity"`
	// The key of the parameter.
	//
	// The options are `auto_mv` , `datestyle` , `enable_case_sensitive_identifier` , `enable_user_activity_logging` , `query_group` , `search_path` , `require_ssl` , `use_fips_ssl` , and query monitoring metrics that let you define performance boundaries. For more information about query monitoring rules and available metrics, see [Query monitoring metrics for Amazon Redshift Serverless](https://docs.aws.amazon.com/redshift/latest/dg/cm-c-wlm-query-monitoring-rules.html#cm-c-wlm-query-monitoring-metrics-serverless) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html#cfn-redshiftserverless-workgroup-configparameters
	//
	ConfigParameters interface{} `field:"optional" json:"configParameters" yaml:"configParameters"`
	// The value that specifies whether to enable enhanced virtual private cloud (VPC) routing, which forces Amazon Redshift Serverless to route traffic through your VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html#cfn-redshiftserverless-workgroup-enhancedvpcrouting
	//
	// Default: - false.
	//
	EnhancedVpcRouting interface{} `field:"optional" json:"enhancedVpcRouting" yaml:"enhancedVpcRouting"`
	// The maximum data-warehouse capacity Amazon Redshift Serverless uses to serve queries.
	//
	// The max capacity is specified in RPUs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html#cfn-redshiftserverless-workgroup-maxcapacity
	//
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The namespace the workgroup is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html#cfn-redshiftserverless-workgroup-namespacename
	//
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
	// The custom port to use when connecting to a workgroup.
	//
	// Valid port ranges are 5431-5455 and 8191-8215. The default is 5439.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html#cfn-redshiftserverless-workgroup-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// An object that represents the price performance target settings for the workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html#cfn-redshiftserverless-workgroup-priceperformancetarget
	//
	PricePerformanceTarget interface{} `field:"optional" json:"pricePerformanceTarget" yaml:"pricePerformanceTarget"`
	// A value that specifies whether the workgroup can be accessible from a public network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html#cfn-redshiftserverless-workgroup-publiclyaccessible
	//
	// Default: - false.
	//
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The recovery point id to restore from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html#cfn-redshiftserverless-workgroup-recoverypointid
	//
	RecoveryPointId *string `field:"optional" json:"recoveryPointId" yaml:"recoveryPointId"`
	// A list of security group IDs to associate with the workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html#cfn-redshiftserverless-workgroup-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The Amazon Resource Name (ARN) of the snapshot to restore from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html#cfn-redshiftserverless-workgroup-snapshotarn
	//
	SnapshotArn *string `field:"optional" json:"snapshotArn" yaml:"snapshotArn"`
	// The snapshot name to restore from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html#cfn-redshiftserverless-workgroup-snapshotname
	//
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
	// The Amazon Web Services account that owns the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html#cfn-redshiftserverless-workgroup-snapshotowneraccount
	//
	SnapshotOwnerAccount *string `field:"optional" json:"snapshotOwnerAccount" yaml:"snapshotOwnerAccount"`
	// A list of subnet IDs the workgroup is associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html#cfn-redshiftserverless-workgroup-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The map of the key-value pairs used to tag the workgroup.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html#cfn-redshiftserverless-workgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// An optional parameter for the name of the track for the workgroup.
	//
	// If you don't provide a track name, the workgroup is assigned to the current track.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html#cfn-redshiftserverless-workgroup-trackname
	//
	TrackName *string `field:"optional" json:"trackName" yaml:"trackName"`
	// The collection of computing resources from which an endpoint is created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshiftserverless-workgroup.html#cfn-redshiftserverless-workgroup-workgroup
	//
	Workgroup interface{} `field:"optional" json:"workgroup" yaml:"workgroup"`
}

