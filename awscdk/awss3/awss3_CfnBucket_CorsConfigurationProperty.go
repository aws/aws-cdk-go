package awss3


// Describes the cross-origin access configuration for objects in an Amazon S3 bucket.
//
// For more information, see [Enabling Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   corsConfigurationProperty := &corsConfigurationProperty{
//   	corsRules: []interface{}{
//   		&corsRuleProperty{
//   			allowedMethods: []*string{
//   				jsii.String("allowedMethods"),
//   			},
//   			allowedOrigins: []*string{
//   				jsii.String("allowedOrigins"),
//   			},
//
//   			// the properties below are optional
//   			allowedHeaders: []*string{
//   				jsii.String("allowedHeaders"),
//   			},
//   			exposedHeaders: []*string{
//   				jsii.String("exposedHeaders"),
//   			},
//   			id: jsii.String("id"),
//   			maxAge: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnBucket_CorsConfigurationProperty struct {
	// A set of origins and methods (cross-origin access that you want to allow).
	//
	// You can add up to 100 rules to the configuration.
	CorsRules interface{} `field:"required" json:"corsRules" yaml:"corsRules"`
}

