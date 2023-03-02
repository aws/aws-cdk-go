package awsrds


// Properties for defining a `CfnDBProxyTargetGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBProxyTargetGroupProps := &cfnDBProxyTargetGroupProps{
//   	dbProxyName: jsii.String("dbProxyName"),
//   	targetGroupName: jsii.String("targetGroupName"),
//
//   	// the properties below are optional
//   	connectionPoolConfigurationInfo: &connectionPoolConfigurationInfoFormatProperty{
//   		connectionBorrowTimeout: jsii.Number(123),
//   		initQuery: jsii.String("initQuery"),
//   		maxConnectionsPercent: jsii.Number(123),
//   		maxIdleConnectionsPercent: jsii.Number(123),
//   		sessionPinningFilters: []*string{
//   			jsii.String("sessionPinningFilters"),
//   		},
//   	},
//   	dbClusterIdentifiers: []*string{
//   		jsii.String("dbClusterIdentifiers"),
//   	},
//   	dbInstanceIdentifiers: []*string{
//   		jsii.String("dbInstanceIdentifiers"),
//   	},
//   }
//
type CfnDBProxyTargetGroupProps struct {
	// The identifier of the `DBProxy` that is associated with the `DBProxyTargetGroup` .
	DbProxyName *string `field:"required" json:"dbProxyName" yaml:"dbProxyName"`
	// The identifier for the target group.
	//
	// > Currently, this property must be set to `default` .
	TargetGroupName *string `field:"required" json:"targetGroupName" yaml:"targetGroupName"`
	// Settings that control the size and behavior of the connection pool associated with a `DBProxyTargetGroup` .
	ConnectionPoolConfigurationInfo interface{} `field:"optional" json:"connectionPoolConfigurationInfo" yaml:"connectionPoolConfigurationInfo"`
	// One or more DB cluster identifiers.
	DbClusterIdentifiers *[]*string `field:"optional" json:"dbClusterIdentifiers" yaml:"dbClusterIdentifiers"`
	// One or more DB instance identifiers.
	DbInstanceIdentifiers *[]*string `field:"optional" json:"dbInstanceIdentifiers" yaml:"dbInstanceIdentifiers"`
}

