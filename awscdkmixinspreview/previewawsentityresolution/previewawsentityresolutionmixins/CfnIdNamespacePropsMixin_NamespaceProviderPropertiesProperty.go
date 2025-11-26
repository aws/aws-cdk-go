package previewawsentityresolutionmixins


// An object containing `providerConfiguration` and `providerServiceArn` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   namespaceProviderPropertiesProperty := &NamespaceProviderPropertiesProperty{
//   	ProviderConfiguration: map[string]*string{
//   		"providerConfigurationKey": jsii.String("providerConfiguration"),
//   	},
//   	ProviderServiceArn: jsii.String("providerServiceArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-namespaceproviderproperties.html
//
type CfnIdNamespacePropsMixin_NamespaceProviderPropertiesProperty struct {
	// An object which defines any additional configurations required by the provider service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-namespaceproviderproperties.html#cfn-entityresolution-idnamespace-namespaceproviderproperties-providerconfiguration
	//
	ProviderConfiguration interface{} `field:"optional" json:"providerConfiguration" yaml:"providerConfiguration"`
	// The Amazon Resource Name (ARN) of the provider service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-namespaceproviderproperties.html#cfn-entityresolution-idnamespace-namespaceproviderproperties-providerservicearn
	//
	ProviderServiceArn *string `field:"optional" json:"providerServiceArn" yaml:"providerServiceArn"`
}

