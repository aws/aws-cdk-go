package awspipes


// The array properties for the submitted job, such as the size of the array.
//
// The array size can be between 2 and 10,000. If you specify array properties for a job, it becomes an array job. This parameter is used only if the target is an AWS Batch job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   batchArrayPropertiesProperty := &BatchArrayPropertiesProperty{
//   	Size: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-batcharrayproperties.html
//
type CfnPipe_BatchArrayPropertiesProperty struct {
	// The size of the array, if this is an array batch job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-batcharrayproperties.html#cfn-pipes-pipe-batcharrayproperties-size
	//
	// Default: - 0.
	//
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}

