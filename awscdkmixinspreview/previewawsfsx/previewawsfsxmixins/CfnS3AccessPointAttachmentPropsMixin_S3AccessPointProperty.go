package previewawsfsxmixins


// Describes the S3 access point configuration of the S3 access point attachment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnS3AccessPointAttachmentPropsMixin_S3AccessPointProperty struct {
	// The S3 access point's alias.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-s3accesspoint.html#cfn-fsx-s3accesspointattachment-s3accesspoint-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The S3 access point's policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-s3accesspoint.html#cfn-fsx-s3accesspointattachment-s3accesspoint-policy
	//
	Policy interface{} `field:"optional" json:"policy" yaml:"policy"`
	// The S3 access point's ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-s3accesspoint.html#cfn-fsx-s3accesspointattachment-s3accesspoint-resourcearn
	//
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// The S3 access point's virtual private cloud (VPC) configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-s3accesspoint.html#cfn-fsx-s3accesspointattachment-s3accesspoint-vpcconfiguration
	//
	VpcConfiguration interface{} `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

