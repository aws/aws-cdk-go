package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   decalSettingsConfigurationProperty := &DecalSettingsConfigurationProperty{
//   	CustomDecalSettings: []interface{}{
//   		&DecalSettingsProperty{
//   			DecalColor: jsii.String("decalColor"),
//   			DecalPatternType: jsii.String("decalPatternType"),
//   			DecalStyleType: jsii.String("decalStyleType"),
//   			DecalVisibility: jsii.String("decalVisibility"),
//   			ElementValue: jsii.String("elementValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-decalsettingsconfiguration.html
//
type CfnTemplatePropsMixin_DecalSettingsConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-decalsettingsconfiguration.html#cfn-quicksight-template-decalsettingsconfiguration-customdecalsettings
	//
	CustomDecalSettings interface{} `field:"optional" json:"customDecalSettings" yaml:"customDecalSettings"`
}

