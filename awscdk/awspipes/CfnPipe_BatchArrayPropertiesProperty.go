package awspipes


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

