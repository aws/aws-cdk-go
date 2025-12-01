package previewawssagemakerevents


// Type definition for TransformResources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   transformResources := &TransformResources{
//   	InstanceCount: []*string{
//   		jsii.String("instanceCount"),
//   	},
//   	InstanceType: []*string{
//   		jsii.String("instanceType"),
//   	},
//   }
//
// Experimental.
type ModelEvents_AWSAPICallViaCloudTrail_TransformResources struct {
	// instanceCount property.
	//
	// Specify an array of string values to match this event if the actual value of instanceCount is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceCount *[]*string `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// instanceType property.
	//
	// Specify an array of string values to match this event if the actual value of instanceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceType *[]*string `field:"optional" json:"instanceType" yaml:"instanceType"`
}

