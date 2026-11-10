package awscloudwatch


// Identity Center configuration.
//
// Required when identityProviders includes IDC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identityCenterConfigurationProperty := &IdentityCenterConfigurationProperty{
//   	IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-domain-identitycenterconfiguration.html
//
type CfnDomain_IdentityCenterConfigurationProperty struct {
	// Identity Center instance ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-domain-identitycenterconfiguration.html#cfn-cloudwatch-domain-identitycenterconfiguration-identitycenterinstancearn
	//
	IdentityCenterInstanceArn *string `field:"optional" json:"identityCenterInstanceArn" yaml:"identityCenterInstanceArn"`
}

