package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   corsConfigurationProperty := &CorsConfigurationProperty{
//   	AllowOrigin: jsii.String("allowOrigin"),
//
//   	// the properties below are optional
//   	AllowCredentials: jsii.Boolean(false),
//   	AllowHeaders: jsii.String("allowHeaders"),
//   	AllowMethods: jsii.String("allowMethods"),
//   	MaxAge: jsii.String("maxAge"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-api-corsconfiguration.html
//
type CfnApi_CorsConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-api-corsconfiguration.html#cfn-serverless-api-corsconfiguration-alloworigin
	//
	AllowOrigin *string `field:"required" json:"allowOrigin" yaml:"allowOrigin"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-api-corsconfiguration.html#cfn-serverless-api-corsconfiguration-allowcredentials
	//
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-api-corsconfiguration.html#cfn-serverless-api-corsconfiguration-allowheaders
	//
	AllowHeaders *string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-api-corsconfiguration.html#cfn-serverless-api-corsconfiguration-allowmethods
	//
	AllowMethods *string `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-api-corsconfiguration.html#cfn-serverless-api-corsconfiguration-maxage
	//
	MaxAge *string `field:"optional" json:"maxAge" yaml:"maxAge"`
}

