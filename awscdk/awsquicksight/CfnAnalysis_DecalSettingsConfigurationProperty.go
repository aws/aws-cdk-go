package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-decalsettingsconfiguration.html
//
type CfnAnalysis_DecalSettingsConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-decalsettingsconfiguration.html#cfn-quicksight-analysis-decalsettingsconfiguration-customdecalsettings
	//
	CustomDecalSettings interface{} `field:"optional" json:"customDecalSettings" yaml:"customDecalSettings"`
}

