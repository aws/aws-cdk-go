package awss3outposts


// The container for the lifecycle configuration for the objects stored in an S3 on Outposts bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var filter interface{}
//
//   lifecycleConfigurationProperty := &lifecycleConfigurationProperty{
//   	rules: []interface{}{
//   		&ruleProperty{
//   			status: jsii.String("status"),
//
//   			// the properties below are optional
//   			abortIncompleteMultipartUpload: &abortIncompleteMultipartUploadProperty{
//   				daysAfterInitiation: jsii.Number(123),
//   			},
//   			expirationDate: jsii.String("expirationDate"),
//   			expirationInDays: jsii.Number(123),
//   			filter: filter,
//   			id: jsii.String("id"),
//   		},
//   	},
//   }
//
type CfnBucket_LifecycleConfigurationProperty struct {
	// The container for the lifecycle configuration rules for the objects stored in the S3 on Outposts bucket.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

