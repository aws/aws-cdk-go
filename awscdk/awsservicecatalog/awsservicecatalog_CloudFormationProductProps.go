package awsservicecatalog


// Properties for a Cloudformation Product.
//
// Example:
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//
//   type s3BucketProduct struct {
//   	productStack
//   }
//
//   func newS3BucketProduct(scope construct, id *string) *s3BucketProduct {
//   	this := &s3BucketProduct{}
//   	servicecatalog.NewProductStack_Override(this, scope, id)
//
//   	s3.NewBucket(this, jsii.String("BucketProduct"))
//   	return this
//   }
//
//   product := servicecatalog.NewCloudFormationProduct(this, jsii.String("Product"), &cloudFormationProductProps{
//   	productName: jsii.String("My Product"),
//   	owner: jsii.String("Product Owner"),
//   	productVersions: []cloudFormationProductVersion{
//   		&cloudFormationProductVersion{
//   			productVersionName: jsii.String("v1"),
//   			cloudFormationTemplate: servicecatalog.cloudFormationTemplate.fromProductStack(NewS3BucketProduct(this, jsii.String("S3BucketProduct"))),
//   		},
//   	},
//   })
//
type CloudFormationProductProps struct {
	// The owner of the product.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The name of the product.
	ProductName *string `field:"required" json:"productName" yaml:"productName"`
	// The configuration of the product version.
	ProductVersions *[]*CloudFormationProductVersion `field:"required" json:"productVersions" yaml:"productVersions"`
	// The description of the product.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The distributor of the product.
	Distributor *string `field:"optional" json:"distributor" yaml:"distributor"`
	// The language code.
	//
	// Controls language for logging and errors.
	MessageLanguage MessageLanguage `field:"optional" json:"messageLanguage" yaml:"messageLanguage"`
	// Whether to give provisioning artifacts a new unique identifier when the product attributes or provisioning artifacts is updated.
	ReplaceProductVersionIds *bool `field:"optional" json:"replaceProductVersionIds" yaml:"replaceProductVersionIds"`
	// The support information about the product.
	SupportDescription *string `field:"optional" json:"supportDescription" yaml:"supportDescription"`
	// The contact email for product support.
	SupportEmail *string `field:"optional" json:"supportEmail" yaml:"supportEmail"`
	// The contact URL for product support.
	SupportUrl *string `field:"optional" json:"supportUrl" yaml:"supportUrl"`
	// TagOptions associated directly to a product.
	TagOptions TagOptions `field:"optional" json:"tagOptions" yaml:"tagOptions"`
}

