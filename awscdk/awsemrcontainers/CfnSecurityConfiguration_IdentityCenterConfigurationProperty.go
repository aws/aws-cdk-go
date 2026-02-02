package awsemrcontainers


// Identity Center configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identityCenterConfigurationProperty := &IdentityCenterConfigurationProperty{
//   	EnableIdentityCenter: jsii.Boolean(false),
//   	IdentityCenterApplicationAssignmentRequired: jsii.Boolean(false),
//   	IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-identitycenterconfiguration.html
//
type CfnSecurityConfiguration_IdentityCenterConfigurationProperty struct {
	// Whether to enable Identity Center integration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-identitycenterconfiguration.html#cfn-emrcontainers-securityconfiguration-identitycenterconfiguration-enableidentitycenter
	//
	EnableIdentityCenter interface{} `field:"optional" json:"enableIdentityCenter" yaml:"enableIdentityCenter"`
	// Whether Identity Center application assignment is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-identitycenterconfiguration.html#cfn-emrcontainers-securityconfiguration-identitycenterconfiguration-identitycenterapplicationassignmentrequired
	//
	IdentityCenterApplicationAssignmentRequired interface{} `field:"optional" json:"identityCenterApplicationAssignmentRequired" yaml:"identityCenterApplicationAssignmentRequired"`
	// The ARN of the Identity Center instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-identitycenterconfiguration.html#cfn-emrcontainers-securityconfiguration-identitycenterconfiguration-identitycenterinstancearn
	//
	IdentityCenterInstanceArn *string `field:"optional" json:"identityCenterInstanceArn" yaml:"identityCenterInstanceArn"`
}

