package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var authorizers interface{}
//
//   httpApiAuthProperty := &httpApiAuthProperty{
//   	authorizers: authorizers,
//   	defaultAuthorizer: jsii.String("defaultAuthorizer"),
//   }
//
type CfnHttpApi_HttpApiAuthProperty struct {
	// `CfnHttpApi.HttpApiAuthProperty.Authorizers`.
	Authorizers interface{} `field:"optional" json:"authorizers" yaml:"authorizers"`
	// `CfnHttpApi.HttpApiAuthProperty.DefaultAuthorizer`.
	DefaultAuthorizer *string `field:"optional" json:"defaultAuthorizer" yaml:"defaultAuthorizer"`
}

