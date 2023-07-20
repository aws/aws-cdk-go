package awss3


// Specifies the container element for Object Ownership rules.
//
// S3 Object Ownership is an Amazon S3 bucket-level setting that you can use to disable access control lists (ACLs) and take ownership of every object in your bucket, simplifying access management for data stored in Amazon S3. For more information, see [Controlling ownership of objects and disabling ACLs](https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ownershipControlsProperty := &OwnershipControlsProperty{
//   	Rules: []interface{}{
//   		&OwnershipControlsRuleProperty{
//   			ObjectOwnership: jsii.String("objectOwnership"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-ownershipcontrols.html
//
type CfnBucket_OwnershipControlsProperty struct {
	// Specifies the container element for Object Ownership rules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-ownershipcontrols.html#cfn-s3-bucket-ownershipcontrols-rules
	//
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

