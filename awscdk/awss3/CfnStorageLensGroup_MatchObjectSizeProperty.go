package awss3


// Filter to match all of the specified values for the minimum and maximum object size.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matchObjectSizeProperty := &MatchObjectSizeProperty{
//   	BytesGreaterThan: jsii.Number(123),
//   	BytesLessThan: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-matchobjectsize.html
//
type CfnStorageLensGroup_MatchObjectSizeProperty struct {
	// Minimum object size to which the rule applies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-matchobjectsize.html#cfn-s3-storagelensgroup-matchobjectsize-bytesgreaterthan
	//
	BytesGreaterThan *float64 `field:"optional" json:"bytesGreaterThan" yaml:"bytesGreaterThan"`
	// Maximum object size to which the rule applies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-matchobjectsize.html#cfn-s3-storagelensgroup-matchobjectsize-byteslessthan
	//
	BytesLessThan *float64 `field:"optional" json:"bytesLessThan" yaml:"bytesLessThan"`
}

