package awsglue


// Properties for defining a `CfnIdentityCenterConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIdentityCenterConfigurationProps := &CfnIdentityCenterConfigurationProps{
//   	InstanceArn: jsii.String("instanceArn"),
//
//   	// the properties below are optional
//   	Scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   	UserBackgroundSessionsEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-identitycenterconfiguration.html
//
type CfnIdentityCenterConfigurationProps struct {
	// The IAM identity center instance arn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-identitycenterconfiguration.html#cfn-glue-identitycenterconfiguration-instancearn
	//
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The downstream scopes that Glue identity center configuration can access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-identitycenterconfiguration.html#cfn-glue-identitycenterconfiguration-scopes
	//
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// Enable or disable user background sessions for Glue Identity Center.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-identitycenterconfiguration.html#cfn-glue-identitycenterconfiguration-userbackgroundsessionsenabled
	//
	UserBackgroundSessionsEnabled interface{} `field:"optional" json:"userBackgroundSessionsEnabled" yaml:"userBackgroundSessionsEnabled"`
}

