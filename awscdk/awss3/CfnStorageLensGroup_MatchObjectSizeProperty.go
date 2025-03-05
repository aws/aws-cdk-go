package awss3


// This resource filters objects that match the specified object size range.
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
	// This property specifies the minimum object size in bytes.
	//
	// The value must be a positive number, greater than 0 and less than 5 TB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-matchobjectsize.html#cfn-s3-storagelensgroup-matchobjectsize-bytesgreaterthan
	//
	BytesGreaterThan *float64 `field:"optional" json:"bytesGreaterThan" yaml:"bytesGreaterThan"`
	// This property specifies the maximum object size in bytes.
	//
	// The value must be a positive number, greater than the minimum object size and less than 5 TB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-storagelensgroup-matchobjectsize.html#cfn-s3-storagelensgroup-matchobjectsize-byteslessthan
	//
	BytesLessThan *float64 `field:"optional" json:"bytesLessThan" yaml:"bytesLessThan"`
}

