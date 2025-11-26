package previewawsquicksightmixins


// The parameters for an IAM Identity Center configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   identityCenterConfigurationProperty := &IdentityCenterConfigurationProperty{
//   	EnableIdentityPropagation: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-identitycenterconfiguration.html
//
type CfnDataSourcePropsMixin_IdentityCenterConfigurationProperty struct {
	// A Boolean option that controls whether Trusted Identity Propagation should be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-identitycenterconfiguration.html#cfn-quicksight-datasource-identitycenterconfiguration-enableidentitypropagation
	//
	EnableIdentityPropagation interface{} `field:"optional" json:"enableIdentityPropagation" yaml:"enableIdentityPropagation"`
}

