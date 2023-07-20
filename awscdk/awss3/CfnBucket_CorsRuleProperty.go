package awss3


// Specifies a cross-origin access rule for an Amazon S3 bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   corsRuleProperty := &CorsRuleProperty{
//   	AllowedMethods: []*string{
//   		jsii.String("allowedMethods"),
//   	},
//   	AllowedOrigins: []*string{
//   		jsii.String("allowedOrigins"),
//   	},
//
//   	// the properties below are optional
//   	AllowedHeaders: []*string{
//   		jsii.String("allowedHeaders"),
//   	},
//   	ExposedHeaders: []*string{
//   		jsii.String("exposedHeaders"),
//   	},
//   	Id: jsii.String("id"),
//   	MaxAge: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-corsrule.html
//
type CfnBucket_CorsRuleProperty struct {
	// An HTTP method that you allow the origin to run.
	//
	// *Allowed values* : `GET` | `PUT` | `HEAD` | `POST` | `DELETE`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-corsrule.html#cfn-s3-bucket-corsrule-allowedmethods
	//
	AllowedMethods *[]*string `field:"required" json:"allowedMethods" yaml:"allowedMethods"`
	// One or more origins you want customers to be able to access the bucket from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-corsrule.html#cfn-s3-bucket-corsrule-allowedorigins
	//
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Headers that are specified in the `Access-Control-Request-Headers` header.
	//
	// These headers are allowed in a preflight OPTIONS request. In response to any preflight OPTIONS request, Amazon S3 returns any requested headers that are allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-corsrule.html#cfn-s3-bucket-corsrule-allowedheaders
	//
	AllowedHeaders *[]*string `field:"optional" json:"allowedHeaders" yaml:"allowedHeaders"`
	// One or more headers in the response that you want customers to be able to access from their applications (for example, from a JavaScript `XMLHttpRequest` object).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-corsrule.html#cfn-s3-bucket-corsrule-exposedheaders
	//
	ExposedHeaders *[]*string `field:"optional" json:"exposedHeaders" yaml:"exposedHeaders"`
	// A unique identifier for this rule.
	//
	// The value must be no more than 255 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-corsrule.html#cfn-s3-bucket-corsrule-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The time in seconds that your browser is to cache the preflight response for the specified resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-corsrule.html#cfn-s3-bucket-corsrule-maxage
	//
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
}

