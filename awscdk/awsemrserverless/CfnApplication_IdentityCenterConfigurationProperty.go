package awsemrserverless


// The IAM IdentityCenter configuration for trusted-identity-propagation on this application.
//
// Supported with release labels emr-7.8.0 and above.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-identitycenterconfiguration.html
//
type CfnApplication_IdentityCenterConfigurationProperty struct {
	// The IAM IdentityCenter instance arn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-identitycenterconfiguration.html#cfn-emrserverless-application-identitycenterconfiguration-identitycenterinstancearn
	//
	IdentityCenterInstanceArn *string `field:"optional" json:"identityCenterInstanceArn" yaml:"identityCenterInstanceArn"`
}

