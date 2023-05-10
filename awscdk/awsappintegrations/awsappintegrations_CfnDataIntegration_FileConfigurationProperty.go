package awsappintegrations


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
type CfnDataIntegration_FileConfigurationProperty struct {
	// `CfnDataIntegration.FileConfigurationProperty.Folders`.
	Folders *[]*string `field:"required" json:"folders" yaml:"folders"`
	// `CfnDataIntegration.FileConfigurationProperty.Filters`.
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
}

