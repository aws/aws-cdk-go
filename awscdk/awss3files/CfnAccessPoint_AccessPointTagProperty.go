package awss3files


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessPointTagProperty := &AccessPointTagProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-accesspoint-accesspointtag.html
//
type CfnAccessPoint_AccessPointTagProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-accesspoint-accesspointtag.html#cfn-s3files-accesspoint-accesspointtag-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3files-accesspoint-accesspointtag.html#cfn-s3files-accesspoint-accesspointtag-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

