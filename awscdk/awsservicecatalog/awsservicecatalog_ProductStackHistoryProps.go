package awsservicecatalog


// Properties for a ProductStackHistory.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//
//   type s3BucketProduct struct {
//   	productStack
//   }
//
//   func newS3BucketProduct(scope cdk.Construct, id *string) *s3BucketProduct {
//   	this := &s3BucketProduct{}
//   	servicecatalog.NewProductStack_Override(this, scope, id)
//
//   	s3.NewBucket(this, jsii.String("BucketProductV2"))
//   	return this
//   }
//
//   productStackHistory := servicecatalog.NewProductStackHistory(this, jsii.String("ProductStackHistory"), &productStackHistoryProps{
//   	productStack: NewS3BucketProduct(this, jsii.String("S3BucketProduct")),
//   	currentVersionName: jsii.String("v2"),
//   	currentVersionLocked: jsii.Boolean(true),
//   })
//
//   product := servicecatalog.NewCloudFormationProduct(this, jsii.String("MyFirstProduct"), &cloudFormationProductProps{
//   	productName: jsii.String("My Product"),
//   	owner: jsii.String("Product Owner"),
//   	productVersions: []cloudFormationProductVersion{
//   		productStackHistory.currentVersion(),
//   	},
//   })
//
type ProductStackHistoryProps struct {
	// If this is set to true, the ProductStack will not be overwritten if a snapshot is found for the currentVersionName.
	CurrentVersionLocked *bool `field:"required" json:"currentVersionLocked" yaml:"currentVersionLocked"`
	// The current version name of the ProductStack.
	CurrentVersionName *string `field:"required" json:"currentVersionName" yaml:"currentVersionName"`
	// The ProductStack whose history will be retained as a snapshot.
	ProductStack ProductStack `field:"required" json:"productStack" yaml:"productStack"`
	// The description of the product version.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The directory where template snapshots will be stored.
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// Whether the specified product template will be validated by CloudFormation.
	//
	// If turned off, an invalid template configuration can be stored.
	ValidateTemplate *bool `field:"optional" json:"validateTemplate" yaml:"validateTemplate"`
}

