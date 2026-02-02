package previewawsemrcontainersmixins


// Container provider information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   containerProviderProperty := &ContainerProviderProperty{
//   	Id: jsii.String("id"),
//   	Info: &ContainerInfoProperty{
//   		EksInfo: &EksInfoProperty{
//   			Namespace: jsii.String("namespace"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-containerprovider.html
//
type CfnSecurityConfigurationPropsMixin_ContainerProviderProperty struct {
	// The container provider ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-containerprovider.html#cfn-emrcontainers-securityconfiguration-containerprovider-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Container information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-containerprovider.html#cfn-emrcontainers-securityconfiguration-containerprovider-info
	//
	Info interface{} `field:"optional" json:"info" yaml:"info"`
	// The container provider type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-securityconfiguration-containerprovider.html#cfn-emrcontainers-securityconfiguration-containerprovider-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

