package previewawssagemakerevents


// Type definition for ResourceConfig.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceConfig := &ResourceConfig{
//   	InstanceCount: []*string{
//   		jsii.String("instanceCount"),
//   	},
//   	InstanceType: []*string{
//   		jsii.String("instanceType"),
//   	},
//   	VolumeSizeInGb: []*string{
//   		jsii.String("volumeSizeInGb"),
//   	},
//   }
//
// Experimental.
type SageMakerTrainingJobStateChange_ResourceConfig struct {
	// InstanceCount property.
	//
	// Specify an array of string values to match this event if the actual value of InstanceCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceCount *[]*string `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// InstanceType property.
	//
	// Specify an array of string values to match this event if the actual value of InstanceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceType *[]*string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// VolumeSizeInGB property.
	//
	// Specify an array of string values to match this event if the actual value of VolumeSizeInGB is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VolumeSizeInGb *[]*string `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

