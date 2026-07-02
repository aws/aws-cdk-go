package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   resourcesProperty := &ResourcesProperty{
//   	MinimumMemoryInMiB: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-resources.html
//
type CfnMicrovmImagePropsMixin_ResourcesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-microvmimage-resources.html#cfn-lambda-microvmimage-resources-minimummemoryinmib
	//
	MinimumMemoryInMiB *float64 `field:"optional" json:"minimumMemoryInMiB" yaml:"minimumMemoryInMiB"`
}

