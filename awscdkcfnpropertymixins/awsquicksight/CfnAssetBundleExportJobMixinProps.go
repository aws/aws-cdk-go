package awsquicksight


// Properties for CfnAssetBundleExportJobPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnAssetBundleExportJobMixinProps := &CfnAssetBundleExportJobMixinProps{
//   	AssetBundleExportJobId: jsii.String("assetBundleExportJobId"),
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	ExportFormat: jsii.String("exportFormat"),
//   	IncludeAllDependencies: jsii.Boolean(false),
//   	IncludeFolderMembers: jsii.String("includeFolderMembers"),
//   	IncludeFolderMemberships: jsii.Boolean(false),
//   	IncludePermissions: jsii.Boolean(false),
//   	IncludeTags: jsii.Boolean(false),
//   	ResourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleexportjob.html
//
type CfnAssetBundleExportJobMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleexportjob.html#cfn-quicksight-assetbundleexportjob-assetbundleexportjobid
	//
	AssetBundleExportJobId *string `field:"optional" json:"assetBundleExportJobId" yaml:"assetBundleExportJobId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleexportjob.html#cfn-quicksight-assetbundleexportjob-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleexportjob.html#cfn-quicksight-assetbundleexportjob-exportformat
	//
	ExportFormat *string `field:"optional" json:"exportFormat" yaml:"exportFormat"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleexportjob.html#cfn-quicksight-assetbundleexportjob-includealldependencies
	//
	IncludeAllDependencies interface{} `field:"optional" json:"includeAllDependencies" yaml:"includeAllDependencies"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleexportjob.html#cfn-quicksight-assetbundleexportjob-includefoldermembers
	//
	IncludeFolderMembers *string `field:"optional" json:"includeFolderMembers" yaml:"includeFolderMembers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleexportjob.html#cfn-quicksight-assetbundleexportjob-includefoldermemberships
	//
	IncludeFolderMemberships interface{} `field:"optional" json:"includeFolderMemberships" yaml:"includeFolderMemberships"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleexportjob.html#cfn-quicksight-assetbundleexportjob-includepermissions
	//
	IncludePermissions interface{} `field:"optional" json:"includePermissions" yaml:"includePermissions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleexportjob.html#cfn-quicksight-assetbundleexportjob-includetags
	//
	IncludeTags interface{} `field:"optional" json:"includeTags" yaml:"includeTags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleexportjob.html#cfn-quicksight-assetbundleexportjob-resourcearns
	//
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
}

