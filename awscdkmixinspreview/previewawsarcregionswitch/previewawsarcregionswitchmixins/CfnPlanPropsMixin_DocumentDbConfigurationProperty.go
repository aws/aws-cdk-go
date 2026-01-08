package previewawsarcregionswitchmixins


// Configuration for Amazon DocumentDB global clusters used in a Region switch plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   documentDbConfigurationProperty := &DocumentDbConfigurationProperty{
//   	Behavior: jsii.String("behavior"),
//   	CrossAccountRole: jsii.String("crossAccountRole"),
//   	DatabaseClusterArns: []*string{
//   		jsii.String("databaseClusterArns"),
//   	},
//   	ExternalId: jsii.String("externalId"),
//   	GlobalClusterIdentifier: jsii.String("globalClusterIdentifier"),
//   	TimeoutMinutes: jsii.Number(123),
//   	Ungraceful: &DocumentDbUngracefulProperty{
//   		Ungraceful: jsii.String("ungraceful"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-documentdbconfiguration.html
//
type CfnPlanPropsMixin_DocumentDbConfigurationProperty struct {
	// The behavior for a global cluster, that is, only allow switchover or also allow failover.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-documentdbconfiguration.html#cfn-arcregionswitch-plan-documentdbconfiguration-behavior
	//
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
	// The cross account role for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-documentdbconfiguration.html#cfn-arcregionswitch-plan-documentdbconfiguration-crossaccountrole
	//
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// The database cluster Amazon Resource Names (ARNs) for a DocumentDB global cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-documentdbconfiguration.html#cfn-arcregionswitch-plan-documentdbconfiguration-databaseclusterarns
	//
	DatabaseClusterArns *[]*string `field:"optional" json:"databaseClusterArns" yaml:"databaseClusterArns"`
	// The external ID (secret key) for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-documentdbconfiguration.html#cfn-arcregionswitch-plan-documentdbconfiguration-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// The global cluster identifier for a DocumentDB global cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-documentdbconfiguration.html#cfn-arcregionswitch-plan-documentdbconfiguration-globalclusteridentifier
	//
	GlobalClusterIdentifier *string `field:"optional" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// The timeout value specified for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-documentdbconfiguration.html#cfn-arcregionswitch-plan-documentdbconfiguration-timeoutminutes
	//
	// Default: - 60.
	//
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// The settings for ungraceful execution.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-documentdbconfiguration.html#cfn-arcregionswitch-plan-documentdbconfiguration-ungraceful
	//
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

