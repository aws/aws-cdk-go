package previewawsquicksightmixins


// The parameters for Snowflake.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   snowflakeParametersProperty := &SnowflakeParametersProperty{
//   	AuthenticationType: jsii.String("authenticationType"),
//   	Database: jsii.String("database"),
//   	DatabaseAccessControlRole: jsii.String("databaseAccessControlRole"),
//   	Host: jsii.String("host"),
//   	OAuthParameters: &OAuthParametersProperty{
//   		IdentityProviderResourceUri: jsii.String("identityProviderResourceUri"),
//   		IdentityProviderVpcConnectionProperties: &VpcConnectionPropertiesProperty{
//   			VpcConnectionArn: jsii.String("vpcConnectionArn"),
//   		},
//   		OAuthScope: jsii.String("oAuthScope"),
//   		TokenProviderUrl: jsii.String("tokenProviderUrl"),
//   	},
//   	Warehouse: jsii.String("warehouse"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-snowflakeparameters.html
//
type CfnDataSourcePropsMixin_SnowflakeParametersProperty struct {
	// The authentication type that you want to use for your connection.
	//
	// This parameter accepts OAuth and non-OAuth authentication types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-snowflakeparameters.html#cfn-quicksight-datasource-snowflakeparameters-authenticationtype
	//
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// Database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-snowflakeparameters.html#cfn-quicksight-datasource-snowflakeparameters-database
	//
	Database *string `field:"optional" json:"database" yaml:"database"`
	// The database access control role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-snowflakeparameters.html#cfn-quicksight-datasource-snowflakeparameters-databaseaccesscontrolrole
	//
	DatabaseAccessControlRole *string `field:"optional" json:"databaseAccessControlRole" yaml:"databaseAccessControlRole"`
	// Host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-snowflakeparameters.html#cfn-quicksight-datasource-snowflakeparameters-host
	//
	Host *string `field:"optional" json:"host" yaml:"host"`
	// An object that contains information needed to create a data source connection between an Quick Sight account and Snowflake.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-snowflakeparameters.html#cfn-quicksight-datasource-snowflakeparameters-oauthparameters
	//
	OAuthParameters interface{} `field:"optional" json:"oAuthParameters" yaml:"oAuthParameters"`
	// Warehouse.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-snowflakeparameters.html#cfn-quicksight-datasource-snowflakeparameters-warehouse
	//
	Warehouse *string `field:"optional" json:"warehouse" yaml:"warehouse"`
}

