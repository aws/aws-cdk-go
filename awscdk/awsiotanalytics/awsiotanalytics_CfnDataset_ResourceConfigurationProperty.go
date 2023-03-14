package awsiotanalytics


// The configuration of the resource used to execute the `containerAction` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceConfigurationProperty := &resourceConfigurationProperty{
//   	computeType: jsii.String("computeType"),
//   	volumeSizeInGb: jsii.Number(123),
//   }
//
type CfnDataset_ResourceConfigurationProperty struct {
	// The type of the compute resource used to execute the `containerAction` .
	//
	// Possible values are: `ACU_1` (vCPU=4, memory=16 GiB) or `ACU_2` (vCPU=8, memory=32 GiB).
	ComputeType *string `field:"required" json:"computeType" yaml:"computeType"`
	// The size, in GB, of the persistent storage available to the resource instance used to execute the `containerAction` (min: 1, max: 50).
	VolumeSizeInGb *float64 `field:"required" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

