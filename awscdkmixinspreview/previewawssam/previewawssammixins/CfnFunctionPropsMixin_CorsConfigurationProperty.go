package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   corsConfigurationProperty := &CorsConfigurationProperty{
//   	AllowCredentials: jsii.Boolean(false),
//   	AllowHeaders: jsii.String("allowHeaders"),
//   	AllowMethods: jsii.String("allowMethods"),
//   	AllowOrigin: jsii.String("allowOrigin"),
//   	MaxAge: jsii.String("maxAge"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-corsconfiguration.html
//
type CfnFunctionPropsMixin_CorsConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-corsconfiguration.html#cfn-serverless-function-corsconfiguration-allowcredentials
	//
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-corsconfiguration.html#cfn-serverless-function-corsconfiguration-allowheaders
	//
	AllowHeaders *string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-corsconfiguration.html#cfn-serverless-function-corsconfiguration-allowmethods
	//
	AllowMethods *string `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-corsconfiguration.html#cfn-serverless-function-corsconfiguration-alloworigin
	//
	AllowOrigin *string `field:"optional" json:"allowOrigin" yaml:"allowOrigin"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-corsconfiguration.html#cfn-serverless-function-corsconfiguration-maxage
	//
	MaxAge *string `field:"optional" json:"maxAge" yaml:"maxAge"`
}

