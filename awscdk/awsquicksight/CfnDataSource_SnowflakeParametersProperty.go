package awsquicksight


// The parameters for Snowflake.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snowflakeParametersProperty := &SnowflakeParametersProperty{
//   	Database: jsii.String("database"),
//   	Host: jsii.String("host"),
//   	Warehouse: jsii.String("warehouse"),
//
//   	// the properties below are optional
//   	AuthenticationType: jsii.String("authenticationType"),
//   	DatabaseAccessControlRole: jsii.String("databaseAccessControlRole"),
//   	OAuthParameters: &OAuthParametersProperty{
//   		TokenProviderUrl: jsii.String("tokenProviderUrl"),
//
//   		// the properties below are optional
//   		IdentityProviderResourceUri: jsii.String("identityProviderResourceUri"),
//   		IdentityProviderVpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   			VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   		},
//   		OAuthScope: jsii.String("oAuthScope"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-snowflakeparameters.html
//
type CfnDataSource_SnowflakeParametersProperty struct {
	// Database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-snowflakeparameters.html#cfn-quicksight-datasource-snowflakeparameters-database
	//
	Database *string `field:"required" json:"database" yaml:"database"`
	// Host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-snowflakeparameters.html#cfn-quicksight-datasource-snowflakeparameters-host
	//
	Host *string `field:"required" json:"host" yaml:"host"`
	// Warehouse.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-snowflakeparameters.html#cfn-quicksight-datasource-snowflakeparameters-warehouse
	//
	Warehouse *string `field:"required" json:"warehouse" yaml:"warehouse"`
	// The authentication type that you want to use for your connection.
	//
	// This parameter accepts OAuth and non-OAuth authentication types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-snowflakeparameters.html#cfn-quicksight-datasource-snowflakeparameters-authenticationtype
	//
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// The database access control role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-snowflakeparameters.html#cfn-quicksight-datasource-snowflakeparameters-databaseaccesscontrolrole
	//
	DatabaseAccessControlRole *string `field:"optional" json:"databaseAccessControlRole" yaml:"databaseAccessControlRole"`
	// An object that contains information needed to create a data source connection between an Quick Sight account and Snowflake.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-snowflakeparameters.html#cfn-quicksight-datasource-snowflakeparameters-oauthparameters
	//
	OAuthParameters interface{} `field:"optional" json:"oAuthParameters" yaml:"oAuthParameters"`
}

