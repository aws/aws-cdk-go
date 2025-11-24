package mixinsawsappintegrations


// The configuration for what files should be pulled from the source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var filters interface{}
//
//   fileConfigurationProperty := &FileConfigurationProperty{
//   	Filters: filters,
//   	Folders: []*string{
//   		jsii.String("folders"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-dataintegration-fileconfiguration.html
//
type CfnDataIntegrationPropsMixin_FileConfigurationProperty struct {
	// Restrictions for what files should be pulled from the source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-dataintegration-fileconfiguration.html#cfn-appintegrations-dataintegration-fileconfiguration-filters
	//
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// Identifiers for the source folders to pull all files from recursively.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-dataintegration-fileconfiguration.html#cfn-appintegrations-dataintegration-fileconfiguration-folders
	//
	Folders *[]*string `field:"optional" json:"folders" yaml:"folders"`
}

