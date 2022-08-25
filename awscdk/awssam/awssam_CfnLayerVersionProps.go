package awssam


// Properties for defining a `CfnLayerVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLayerVersionProps := &cfnLayerVersionProps{
//   	compatibleRuntimes: []*string{
//   		jsii.String("compatibleRuntimes"),
//   	},
//   	contentUri: jsii.String("contentUri"),
//   	description: jsii.String("description"),
//   	layerName: jsii.String("layerName"),
//   	licenseInfo: jsii.String("licenseInfo"),
//   	retentionPolicy: jsii.String("retentionPolicy"),
//   }
//
type CfnLayerVersionProps struct {
	// `AWS::Serverless::LayerVersion.CompatibleRuntimes`.
	CompatibleRuntimes *[]*string `field:"optional" json:"compatibleRuntimes" yaml:"compatibleRuntimes"`
	// `AWS::Serverless::LayerVersion.ContentUri`.
	ContentUri interface{} `field:"optional" json:"contentUri" yaml:"contentUri"`
	// `AWS::Serverless::LayerVersion.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Serverless::LayerVersion.LayerName`.
	LayerName *string `field:"optional" json:"layerName" yaml:"layerName"`
	// `AWS::Serverless::LayerVersion.LicenseInfo`.
	LicenseInfo *string `field:"optional" json:"licenseInfo" yaml:"licenseInfo"`
	// `AWS::Serverless::LayerVersion.RetentionPolicy`.
	RetentionPolicy *string `field:"optional" json:"retentionPolicy" yaml:"retentionPolicy"`
}

