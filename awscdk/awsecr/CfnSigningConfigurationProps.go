package awsecr


// Properties for defining a `CfnSigningConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSigningConfigurationProps := &CfnSigningConfigurationProps{
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			SigningProfileArn: jsii.String("signingProfileArn"),
//
//   			// the properties below are optional
//   			RepositoryFilters: []interface{}{
//   				&RepositoryFilterProperty{
//   					Filter: jsii.String("filter"),
//   					FilterType: jsii.String("filterType"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-signingconfiguration.html
//
type CfnSigningConfigurationProps struct {
	// Array of signing rules that define which repositories should be signed and with which signing profiles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-signingconfiguration.html#cfn-ecr-signingconfiguration-rules
	//
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

