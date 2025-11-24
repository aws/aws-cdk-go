package mixinsawss3


// Describes the cross-origin access configuration for objects in an Amazon S3 bucket.
//
// For more information, see [Enabling Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   corsConfigurationProperty := &CorsConfigurationProperty{
//   	CorsRules: []interface{}{
//   		&CorsRuleProperty{
//   			AllowedHeaders: []*string{
//   				jsii.String("allowedHeaders"),
//   			},
//   			AllowedMethods: []*string{
//   				jsii.String("allowedMethods"),
//   			},
//   			AllowedOrigins: []*string{
//   				jsii.String("allowedOrigins"),
//   			},
//   			ExposedHeaders: []*string{
//   				jsii.String("exposedHeaders"),
//   			},
//   			Id: jsii.String("id"),
//   			MaxAge: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-corsconfiguration.html
//
type CfnBucketPropsMixin_CorsConfigurationProperty struct {
	// A set of origins and methods (cross-origin access that you want to allow).
	//
	// You can add up to 100 rules to the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-corsconfiguration.html#cfn-s3-bucket-corsconfiguration-corsrules
	//
	CorsRules interface{} `field:"optional" json:"corsRules" yaml:"corsRules"`
}

