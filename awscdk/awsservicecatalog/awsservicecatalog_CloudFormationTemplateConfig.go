package awsservicecatalog


// Result of binding `Template` into a `Product`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudFormationTemplateConfig := &cloudFormationTemplateConfig{
//   	httpUrl: jsii.String("httpUrl"),
//   }
//
type CloudFormationTemplateConfig struct {
	// The http url of the template in S3.
	HttpUrl *string `field:"required" json:"httpUrl" yaml:"httpUrl"`
}

