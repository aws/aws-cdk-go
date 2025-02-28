package awsappintegrations


// The configuration for what files should be pulled from the source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var filters interface{}
//
//   fileConfigurationProperty := &FileConfigurationProperty{
//   	Folders: []*string{
//   		jsii.String("folders"),
//   	},
//
//   	// the properties below are optional
//   	Filters: filters,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-dataintegration-fileconfiguration.html
//
type CfnDataIntegration_FileConfigurationProperty struct {
	// Identifiers for the source folders to pull all files from recursively.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-dataintegration-fileconfiguration.html#cfn-appintegrations-dataintegration-fileconfiguration-folders
	//
	Folders *[]*string `field:"required" json:"folders" yaml:"folders"`
	// Restrictions for what files should be pulled from the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-dataintegration-fileconfiguration.html#cfn-appintegrations-dataintegration-fileconfiguration-filters
	//
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}

