package awsquicksight


// Properties for defining a `CfnAssetBundleExportJob`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssetBundleExportJobProps := &CfnAssetBundleExportJobProps{
//   	AssetBundleExportJobId: jsii.String("assetBundleExportJobId"),
//   	ExportFormat: jsii.String("exportFormat"),
//   	ResourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//
//   	// the properties below are optional
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	IncludeAllDependencies: jsii.Boolean(false),
//   	IncludeFolderMembers: jsii.String("includeFolderMembers"),
//   	IncludeFolderMemberships: jsii.Boolean(false),
//   	IncludePermissions: jsii.Boolean(false),
//   	IncludeTags: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleexportjob.html
//
type CfnAssetBundleExportJobProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleexportjob.html#cfn-quicksight-assetbundleexportjob-assetbundleexportjobid
	//
	AssetBundleExportJobId *string `field:"required" json:"assetBundleExportJobId" yaml:"assetBundleExportJobId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleexportjob.html#cfn-quicksight-assetbundleexportjob-exportformat
	//
	ExportFormat *string `field:"required" json:"exportFormat" yaml:"exportFormat"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleexportjob.html#cfn-quicksight-assetbundleexportjob-resourcearns
	//
	ResourceArns *[]*string `field:"required" json:"resourceArns" yaml:"resourceArns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-assetbundleexportjob.html#cfn-quicksight-assetbundleexportjob-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
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
}

