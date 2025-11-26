package previewawsfsxmixins


// If included, Amazon S3 restricts access to this access point to requests from the specified virtual private cloud (VPC).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   s3AccessPointVpcConfigurationProperty := &S3AccessPointVpcConfigurationProperty{
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-s3accesspointvpcconfiguration.html
//
type CfnS3AccessPointAttachmentPropsMixin_S3AccessPointVpcConfigurationProperty struct {
	// Specifies the virtual private cloud (VPC) for the S3 access point VPC configuration, if one exists.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-s3accesspointattachment-s3accesspointvpcconfiguration.html#cfn-fsx-s3accesspointattachment-s3accesspointvpcconfiguration-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

