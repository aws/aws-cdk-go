package previewawsemrcontainersmixins


// Lake Formation configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lakeFormationConfigurationProperty := &LakeFormationConfigurationProperty{
//   	AuthorizedSessionTagValue: jsii.String("authorizedSessionTagValue"),
//   	QueryAccessControlEnabled: jsii.Boolean(false),
//   	QueryEngineRoleArn: jsii.String("queryEngineRoleArn"),
//   	SecureNamespaceInfo: &SecureNamespaceInfoProperty{
//   		ClusterId: jsii.String("clusterId"),
//   		Namespace: jsii.String("namespace"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-lakeformationconfiguration.html
//
type CfnSecurityConfigurationPropsMixin_LakeFormationConfigurationProperty struct {
	// The session tag to authorize Lake Formation access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-lakeformationconfiguration.html#cfn-emrcontainers-securityconfiguration-lakeformationconfiguration-authorizedsessiontagvalue
	//
	AuthorizedSessionTagValue *string `field:"optional" json:"authorizedSessionTagValue" yaml:"authorizedSessionTagValue"`
	// Whether query access control is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-lakeformationconfiguration.html#cfn-emrcontainers-securityconfiguration-lakeformationconfiguration-queryaccesscontrolenabled
	//
	QueryAccessControlEnabled interface{} `field:"optional" json:"queryAccessControlEnabled" yaml:"queryAccessControlEnabled"`
	// The ARN of the query engine role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-lakeformationconfiguration.html#cfn-emrcontainers-securityconfiguration-lakeformationconfiguration-queryenginerolearn
	//
	QueryEngineRoleArn *string `field:"optional" json:"queryEngineRoleArn" yaml:"queryEngineRoleArn"`
	// Secure namespace information for Lake Formation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-lakeformationconfiguration.html#cfn-emrcontainers-securityconfiguration-lakeformationconfiguration-securenamespaceinfo
	//
	SecureNamespaceInfo interface{} `field:"optional" json:"secureNamespaceInfo" yaml:"secureNamespaceInfo"`
}

