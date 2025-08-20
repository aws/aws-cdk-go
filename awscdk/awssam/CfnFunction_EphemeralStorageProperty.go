package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ephemeralStorageProperty := &EphemeralStorageProperty{
//   	Size: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-ephemeralstorage.html
//
type CfnFunction_EphemeralStorageProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-ephemeralstorage.html#cfn-serverless-function-ephemeralstorage-size
	//
	Size *float64 `field:"required" json:"size" yaml:"size"`
}

