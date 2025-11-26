package previewawsbatchmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   efsVolumeConfigurationProperty := &EfsVolumeConfigurationProperty{
//   	AuthorizationConfig: &AuthorizationConfigProperty{
//   		AccessPointId: jsii.String("accessPointId"),
//   		Iam: jsii.String("iam"),
//   	},
//   	FileSystemId: jsii.String("fileSystemId"),
//   	RootDirectory: jsii.String("rootDirectory"),
//   	TransitEncryption: jsii.String("transitEncryption"),
//   	TransitEncryptionPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-efsvolumeconfiguration.html
//
type CfnJobDefinitionPropsMixin_EfsVolumeConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-efsvolumeconfiguration.html#cfn-batch-jobdefinition-efsvolumeconfiguration-authorizationconfig
	//
	AuthorizationConfig interface{} `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-efsvolumeconfiguration.html#cfn-batch-jobdefinition-efsvolumeconfiguration-filesystemid
	//
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-efsvolumeconfiguration.html#cfn-batch-jobdefinition-efsvolumeconfiguration-rootdirectory
	//
	RootDirectory *string `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-efsvolumeconfiguration.html#cfn-batch-jobdefinition-efsvolumeconfiguration-transitencryption
	//
	TransitEncryption *string `field:"optional" json:"transitEncryption" yaml:"transitEncryption"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-efsvolumeconfiguration.html#cfn-batch-jobdefinition-efsvolumeconfiguration-transitencryptionport
	//
	TransitEncryptionPort *float64 `field:"optional" json:"transitEncryptionPort" yaml:"transitEncryptionPort"`
}

