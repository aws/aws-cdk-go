package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3KeyFilterProperty := &s3KeyFilterProperty{
//   	rules: []interface{}{
//   		&s3KeyFilterRuleProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnFunction_S3KeyFilterProperty struct {
	// `CfnFunction.S3KeyFilterProperty.Rules`.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

