package awssam


// Properties for defining a `CfnLayerVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLayerVersionProps := &CfnLayerVersionProps{
//   	CompatibleRuntimes: []*string{
//   		jsii.String("compatibleRuntimes"),
//   	},
//   	ContentUri: jsii.String("contentUri"),
//   	Description: jsii.String("description"),
//   	LayerName: jsii.String("layerName"),
//   	LicenseInfo: jsii.String("licenseInfo"),
//   	RetentionPolicy: jsii.String("retentionPolicy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-layerversion.html
//
type CfnLayerVersionProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-layerversion.html#cfn-serverless-layerversion-compatibleruntimes
	//
	CompatibleRuntimes *[]*string `field:"optional" json:"compatibleRuntimes" yaml:"compatibleRuntimes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-layerversion.html#cfn-serverless-layerversion-contenturi
	//
	ContentUri interface{} `field:"optional" json:"contentUri" yaml:"contentUri"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-layerversion.html#cfn-serverless-layerversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-layerversion.html#cfn-serverless-layerversion-layername
	//
	LayerName *string `field:"optional" json:"layerName" yaml:"layerName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-layerversion.html#cfn-serverless-layerversion-licenseinfo
	//
	LicenseInfo *string `field:"optional" json:"licenseInfo" yaml:"licenseInfo"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-layerversion.html#cfn-serverless-layerversion-retentionpolicy
	//
	RetentionPolicy *string `field:"optional" json:"retentionPolicy" yaml:"retentionPolicy"`
}

