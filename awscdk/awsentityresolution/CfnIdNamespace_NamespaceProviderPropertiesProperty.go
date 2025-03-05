package awsentityresolution


// An object containing `ProviderConfiguration` and `ProviderServiceArn` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   namespaceProviderPropertiesProperty := &NamespaceProviderPropertiesProperty{
//   	ProviderServiceArn: jsii.String("providerServiceArn"),
//
//   	// the properties below are optional
//   	ProviderConfiguration: map[string]*string{
//   		"providerConfigurationKey": jsii.String("providerConfiguration"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-namespaceproviderproperties.html
//
type CfnIdNamespace_NamespaceProviderPropertiesProperty struct {
	// The Amazon Resource Name (ARN) of the provider service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-namespaceproviderproperties.html#cfn-entityresolution-idnamespace-namespaceproviderproperties-providerservicearn
	//
	ProviderServiceArn *string `field:"required" json:"providerServiceArn" yaml:"providerServiceArn"`
	// An object which defines any additional configurations required by the provider service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-namespaceproviderproperties.html#cfn-entityresolution-idnamespace-namespaceproviderproperties-providerconfiguration
	//
	ProviderConfiguration interface{} `field:"optional" json:"providerConfiguration" yaml:"providerConfiguration"`
}

