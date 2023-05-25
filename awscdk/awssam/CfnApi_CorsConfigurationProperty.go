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
type CfnApi_CorsConfigurationProperty struct {
	// `CfnApi.CorsConfigurationProperty.AllowOrigin`.
	AllowOrigin *string `field:"required" json:"allowOrigin" yaml:"allowOrigin"`
	// `CfnApi.CorsConfigurationProperty.AllowCredentials`.
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// `CfnApi.CorsConfigurationProperty.AllowHeaders`.
	AllowHeaders *string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// `CfnApi.CorsConfigurationProperty.AllowMethods`.
	AllowMethods *string `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// `CfnApi.CorsConfigurationProperty.MaxAge`.
	MaxAge *string `field:"optional" json:"maxAge" yaml:"maxAge"`
}

