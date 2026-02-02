package awsemrcontainers


// Authorization configuration for the security configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   authorizationConfigurationProperty := &AuthorizationConfigurationProperty{
//   	LakeFormationConfiguration: &LakeFormationConfigurationProperty{
//   		AuthorizedSessionTagValue: jsii.String("authorizedSessionTagValue"),
//   		QueryAccessControlEnabled: jsii.Boolean(false),
//   		QueryEngineRoleArn: jsii.String("queryEngineRoleArn"),
//   		SecureNamespaceInfo: &SecureNamespaceInfoProperty{
//   			ClusterId: jsii.String("clusterId"),
//   			Namespace: jsii.String("namespace"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-authorizationconfiguration.html
//
type CfnSecurityConfiguration_AuthorizationConfigurationProperty struct {
	// Lake Formation configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-authorizationconfiguration.html#cfn-emrcontainers-securityconfiguration-authorizationconfiguration-lakeformationconfiguration
	//
	LakeFormationConfiguration interface{} `field:"optional" json:"lakeFormationConfiguration" yaml:"lakeFormationConfiguration"`
}

