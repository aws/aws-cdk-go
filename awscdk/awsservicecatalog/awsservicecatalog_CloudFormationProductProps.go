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
//   product := servicecatalog.NewCloudFormationProduct(this, jsii.String("Product"), &CloudFormationProductProps{
//   	ProductName: jsii.String("My Product"),
//   	Owner: jsii.String("Product Owner"),
//   	ProductVersions: []cloudFormationProductVersion{
//   		&cloudFormationProductVersion{
//   			ProductVersionName: jsii.String("v1"),
//   			CloudFormationTemplate: servicecatalog.CloudFormationTemplate_FromProductStack(NewS3BucketProduct(this, jsii.String("S3BucketProduct"))),
//   		},
//   	},
//   })
//
// Experimental.
type CloudFormationProductProps struct {
	// The owner of the product.
	// Experimental.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The name of the product.
	// Experimental.
	ProductName *string `field:"required" json:"productName" yaml:"productName"`
	// The configuration of the product version.
	// Experimental.
	ProductVersions *[]*CloudFormationProductVersion `field:"required" json:"productVersions" yaml:"productVersions"`
	// The description of the product.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The distributor of the product.
	// Experimental.
	Distributor *string `field:"optional" json:"distributor" yaml:"distributor"`
	// The language code.
	//
	// Controls language for logging and errors.
	// Experimental.
	MessageLanguage MessageLanguage `field:"optional" json:"messageLanguage" yaml:"messageLanguage"`
	// Whether to give provisioning artifacts a new unique identifier when the product attributes or provisioning artifacts is updated.
	// Experimental.
	ReplaceProductVersionIds *bool `field:"optional" json:"replaceProductVersionIds" yaml:"replaceProductVersionIds"`
	// The support information about the product.
	// Experimental.
	SupportDescription *string `field:"optional" json:"supportDescription" yaml:"supportDescription"`
	// The contact email for product support.
	// Experimental.
	SupportEmail *string `field:"optional" json:"supportEmail" yaml:"supportEmail"`
	// The contact URL for product support.
	// Experimental.
	SupportUrl *string `field:"optional" json:"supportUrl" yaml:"supportUrl"`
	// TagOptions associated directly to a product.
	// Experimental.
	TagOptions TagOptions `field:"optional" json:"tagOptions" yaml:"tagOptions"`
}

