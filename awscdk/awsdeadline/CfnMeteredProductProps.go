package awsdeadline


// Properties for defining a `CfnMeteredProduct`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMeteredProductProps := &CfnMeteredProductProps{
//   	Family: jsii.String("family"),
//   	LicenseEndpointId: jsii.String("licenseEndpointId"),
//   	Port: jsii.Number(123),
//   	ProductId: jsii.String("productId"),
//   	Vendor: jsii.String("vendor"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-meteredproduct.html
//
type CfnMeteredProductProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-meteredproduct.html#cfn-deadline-meteredproduct-family
	//
	Family *string `field:"optional" json:"family" yaml:"family"`
	// The Amazon EC2 identifier of the license endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-meteredproduct.html#cfn-deadline-meteredproduct-licenseendpointid
	//
	LicenseEndpointId *string `field:"optional" json:"licenseEndpointId" yaml:"licenseEndpointId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-meteredproduct.html#cfn-deadline-meteredproduct-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The product ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-meteredproduct.html#cfn-deadline-meteredproduct-productid
	//
	ProductId *string `field:"optional" json:"productId" yaml:"productId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-meteredproduct.html#cfn-deadline-meteredproduct-vendor
	//
	Vendor *string `field:"optional" json:"vendor" yaml:"vendor"`
}

