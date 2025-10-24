package awsemrserverless


// The IAM Identity Center Configuration accepts the Identity Center instance parameter required to enable trusted identity propagation.
//
// This configuration allows identity propagation between integrated services and the Identity Center instance.
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
	// The ARN of the IAM Identity Center instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-identitycenterconfiguration.html#cfn-emrserverless-application-identitycenterconfiguration-identitycenterinstancearn
	//
	IdentityCenterInstanceArn *string `field:"optional" json:"identityCenterInstanceArn" yaml:"identityCenterInstanceArn"`
}

