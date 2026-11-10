package awscloudwatch


// Identity Center configuration.
//
// Required when identityProviders includes IDC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   identityCenterConfigurationProperty := &IdentityCenterConfigurationProperty{
//   	IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-domainfororganization-identitycenterconfiguration.html
//
type CfnDomainForOrganizationPropsMixin_IdentityCenterConfigurationProperty struct {
	// Identity Center instance ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-domainfororganization-identitycenterconfiguration.html#cfn-cloudwatch-domainfororganization-identitycenterconfiguration-identitycenterinstancearn
	//
	IdentityCenterInstanceArn *string `field:"optional" json:"identityCenterInstanceArn" yaml:"identityCenterInstanceArn"`
}

