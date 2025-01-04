package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   volumesProperty := &VolumesProperty{
//   	EfsVolumeConfiguration: &EfsVolumeConfigurationProperty{
//   		FileSystemId: jsii.String("fileSystemId"),
//
//   		// the properties below are optional
//   		AuthorizationConfig: &AuthorizationConfigProperty{
//   			AccessPointId: jsii.String("accessPointId"),
//   			Iam: jsii.String("iam"),
//   		},
//   		RootDirectory: jsii.String("rootDirectory"),
//   		TransitEncryption: jsii.String("transitEncryption"),
//   		TransitEncryptionPort: jsii.Number(123),
//   	},
//   	Host: &VolumesHostProperty{
//   		SourcePath: jsii.String("sourcePath"),
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumes.html
//
type CfnJobDefinition_VolumesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumes.html#cfn-batch-jobdefinition-volumes-efsvolumeconfiguration
	//
	EfsVolumeConfiguration interface{} `field:"optional" json:"efsVolumeConfiguration" yaml:"efsVolumeConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumes.html#cfn-batch-jobdefinition-volumes-host
	//
	Host interface{} `field:"optional" json:"host" yaml:"host"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumes.html#cfn-batch-jobdefinition-volumes-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

