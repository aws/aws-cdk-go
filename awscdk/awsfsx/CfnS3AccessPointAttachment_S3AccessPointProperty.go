package awsfsx


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policy interface{}
//
//   s3AccessPointProperty := &S3AccessPointProperty{
//   	Alias: jsii.String("alias"),
//   	Policy: policy,
//   	ResourceArn: jsii.String("resourceArn"),
//   	VpcConfiguration: &S3AccessPointVpcConfigurationProperty{
//   		VpcId: jsii.String("vpcId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-s3accesspoint.html
//
type CfnS3AccessPointAttachment_S3AccessPointProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-s3accesspoint.html#cfn-fsx-s3accesspointattachment-s3accesspoint-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-s3accesspoint.html#cfn-fsx-s3accesspointattachment-s3accesspoint-policy
	//
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-s3accesspoint.html#cfn-fsx-s3accesspointattachment-s3accesspoint-resourcearn
	//
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-s3accesspoint.html#cfn-fsx-s3accesspointattachment-s3accesspoint-vpcconfiguration
	//
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

