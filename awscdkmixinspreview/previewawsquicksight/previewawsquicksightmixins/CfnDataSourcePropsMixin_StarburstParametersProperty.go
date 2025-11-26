package previewawsquicksightmixins


// The parameters that are required to connect to a Starburst data source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   starburstParametersProperty := &StarburstParametersProperty{
//   	AuthenticationType: jsii.String("authenticationType"),
//   	Catalog: jsii.String("catalog"),
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
//   	Port: jsii.Number(123),
//   	ProductType: jsii.String("productType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html
//
type CfnDataSourcePropsMixin_StarburstParametersProperty struct {
	// The authentication type that you want to use for your connection.
	//
	// This parameter accepts OAuth and non-OAuth authentication types.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html#cfn-quicksight-datasource-starburstparameters-authenticationtype
	//
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// The catalog name for the Starburst data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html#cfn-quicksight-datasource-starburstparameters-catalog
	//
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// The database access control role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html#cfn-quicksight-datasource-starburstparameters-databaseaccesscontrolrole
	//
	DatabaseAccessControlRole *string `field:"optional" json:"databaseAccessControlRole" yaml:"databaseAccessControlRole"`
	// The host name of the Starburst data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html#cfn-quicksight-datasource-starburstparameters-host
	//
	Host *string `field:"optional" json:"host" yaml:"host"`
	// An object that contains information needed to create a data source connection between an Quick Sight account and Starburst.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html#cfn-quicksight-datasource-starburstparameters-oauthparameters
	//
	OAuthParameters interface{} `field:"optional" json:"oAuthParameters" yaml:"oAuthParameters"`
	// The port for the Starburst data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html#cfn-quicksight-datasource-starburstparameters-port
	//
	// Default: - 0.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The product type for the Starburst data source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-starburstparameters.html#cfn-quicksight-datasource-starburstparameters-producttype
	//
	ProductType *string `field:"optional" json:"productType" yaml:"productType"`
}

