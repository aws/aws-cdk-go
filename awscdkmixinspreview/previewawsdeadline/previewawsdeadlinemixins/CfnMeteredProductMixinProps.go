package previewawsdeadlinemixins


// Properties for CfnMeteredProductPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMeteredProductMixinProps := &CfnMeteredProductMixinProps{
//   	LicenseEndpointId: jsii.String("licenseEndpointId"),
//   	ProductId: jsii.String("productId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-meteredproduct.html
//
type CfnMeteredProductMixinProps struct {
	// The Amazon EC2 identifier of the license endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-meteredproduct.html#cfn-deadline-meteredproduct-licenseendpointid
	//
	LicenseEndpointId *string `field:"optional" json:"licenseEndpointId" yaml:"licenseEndpointId"`
	// The product ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-meteredproduct.html#cfn-deadline-meteredproduct-productid
	//
	ProductId *string `field:"optional" json:"productId" yaml:"productId"`
}

