package previewawsiotanalyticsmixins


// The configuration of the resource used to execute the `containerAction` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceConfigurationProperty := &ResourceConfigurationProperty{
//   	ComputeType: jsii.String("computeType"),
//   	VolumeSizeInGb: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-resourceconfiguration.html
//
type CfnDatasetPropsMixin_ResourceConfigurationProperty struct {
	// The type of the compute resource used to execute the `containerAction` .
	//
	// Possible values are: `ACU_1` (vCPU=4, memory=16 GiB) or `ACU_2` (vCPU=8, memory=32 GiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-resourceconfiguration.html#cfn-iotanalytics-dataset-resourceconfiguration-computetype
	//
	ComputeType *string `field:"optional" json:"computeType" yaml:"computeType"`
	// The size, in GB, of the persistent storage available to the resource instance used to execute the `containerAction` (min: 1, max: 50).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-resourceconfiguration.html#cfn-iotanalytics-dataset-resourceconfiguration-volumesizeingb
	//
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

