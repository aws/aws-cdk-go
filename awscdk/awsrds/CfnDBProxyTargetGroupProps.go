package awsrds


// Properties for defining a `CfnDBProxyTargetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBProxyTargetGroupProps := &CfnDBProxyTargetGroupProps{
//   	DbProxyName: jsii.String("dbProxyName"),
//   	TargetGroupName: jsii.String("targetGroupName"),
//
//   	// the properties below are optional
//   	ConnectionPoolConfigurationInfo: &ConnectionPoolConfigurationInfoFormatProperty{
//   		ConnectionBorrowTimeout: jsii.Number(123),
//   		InitQuery: jsii.String("initQuery"),
//   		MaxConnectionsPercent: jsii.Number(123),
//   		MaxIdleConnectionsPercent: jsii.Number(123),
//   		SessionPinningFilters: []*string{
//   			jsii.String("sessionPinningFilters"),
//   		},
//   	},
//   	DbClusterIdentifiers: []*string{
//   		jsii.String("dbClusterIdentifiers"),
//   	},
//   	DbInstanceIdentifiers: []*string{
//   		jsii.String("dbInstanceIdentifiers"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html
//
type CfnDBProxyTargetGroupProps struct {
	// The identifier of the `DBProxy` that is associated with the `DBProxyTargetGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbproxyname
	//
	DbProxyName *string `field:"required" json:"dbProxyName" yaml:"dbProxyName"`
	// The identifier for the target group.
	//
	// > Currently, this property must be set to `default` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-targetgroupname
	//
	TargetGroupName *string `field:"required" json:"targetGroupName" yaml:"targetGroupName"`
	// Settings that control the size and behavior of the connection pool associated with a `DBProxyTargetGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-connectionpoolconfigurationinfo
	//
	ConnectionPoolConfigurationInfo interface{} `field:"optional" json:"connectionPoolConfigurationInfo" yaml:"connectionPoolConfigurationInfo"`
	// One or more DB cluster identifiers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbclusteridentifiers
	//
	DbClusterIdentifiers *[]*string `field:"optional" json:"dbClusterIdentifiers" yaml:"dbClusterIdentifiers"`
	// One or more DB instance identifiers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxytargetgroup.html#cfn-rds-dbproxytargetgroup-dbinstanceidentifiers
	//
	DbInstanceIdentifiers *[]*string `field:"optional" json:"dbInstanceIdentifiers" yaml:"dbInstanceIdentifiers"`
}

