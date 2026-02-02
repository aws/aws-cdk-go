package previewawsemrcontainersmixins


// Authentication configuration for the security configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   authenticationConfigurationProperty := &AuthenticationConfigurationProperty{
//   	IamConfiguration: map[string]*string{
//   		"systemRole": jsii.String("systemRole"),
//   	},
//   	IdentityCenterConfiguration: &IdentityCenterConfigurationProperty{
//   		EnableIdentityCenter: jsii.Boolean(false),
//   		IdentityCenterApplicationAssignmentRequired: jsii.Boolean(false),
//   		IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-authenticationconfiguration.html
//
type CfnSecurityConfigurationPropsMixin_AuthenticationConfigurationProperty struct {
	// IAM configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-authenticationconfiguration.html#cfn-emrcontainers-securityconfiguration-authenticationconfiguration-iamconfiguration
	//
	IamConfiguration interface{} `field:"optional" json:"iamConfiguration" yaml:"iamConfiguration"`
	// Identity Center configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-authenticationconfiguration.html#cfn-emrcontainers-securityconfiguration-authenticationconfiguration-identitycenterconfiguration
	//
	IdentityCenterConfiguration interface{} `field:"optional" json:"identityCenterConfiguration" yaml:"identityCenterConfiguration"`
}

