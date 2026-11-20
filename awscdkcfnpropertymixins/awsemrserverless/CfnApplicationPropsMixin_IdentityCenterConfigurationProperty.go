package awsemrserverless


// The IAM Identity Center Configuration accepts the Identity Center instance parameter required to enable trusted identity propagation.
//
// This configuration allows identity propagation between integrated services and the Identity Center instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   identityCenterConfigurationProperty := &IdentityCenterConfigurationProperty{
//   	IdentityCenterInstanceArn: jsii.String("identityCenterInstanceArn"),
//   	UserBackgroundSessionsEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-identitycenterconfiguration.html
//
type CfnApplicationPropsMixin_IdentityCenterConfigurationProperty struct {
	// The ARN of the IAM Identity Center instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-identitycenterconfiguration.html#cfn-emrserverless-application-identitycenterconfiguration-identitycenterinstancearn
	//
	IdentityCenterInstanceArn *string `field:"optional" json:"identityCenterInstanceArn" yaml:"identityCenterInstanceArn"`
	// Enables user background sessions for this application so sessions can continue running after users log out of their interactive notebook or their Identity Center sessions expire.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrserverless-application-identitycenterconfiguration.html#cfn-emrserverless-application-identitycenterconfiguration-userbackgroundsessionsenabled
	//
	// Default: - false.
	//
	UserBackgroundSessionsEnabled interface{} `field:"optional" json:"userBackgroundSessionsEnabled" yaml:"userBackgroundSessionsEnabled"`
}

