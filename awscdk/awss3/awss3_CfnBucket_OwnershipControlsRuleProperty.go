package awss3


// Specifies an Object Ownership rule.
//
// S3 Object Ownership is an Amazon S3 bucket-level setting that you can use to disable access control lists (ACLs) and take ownership of every object in your bucket, simplifying access management for data stored in Amazon S3. For more information, see [Controlling ownership of objects and disabling ACLs](https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ownershipControlsRuleProperty := &ownershipControlsRuleProperty{
//   	objectOwnership: jsii.String("objectOwnership"),
//   }
//
type CfnBucket_OwnershipControlsRuleProperty struct {
	// Specifies an Object Ownership rule.
	//
	// *Allowed values* : `BucketOwnerEnforced` | `ObjectWriter` | `BucketOwnerPreferred`.
	ObjectOwnership *string `field:"optional" json:"objectOwnership" yaml:"objectOwnership"`
}

