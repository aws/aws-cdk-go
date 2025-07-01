package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3KeyFilterProperty := &S3KeyFilterProperty{
//   	Rules: []interface{}{
//   		&S3KeyFilterRuleProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-s3keyfilter.html
//
type CfnFunction_S3KeyFilterProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-s3keyfilter.html#cfn-serverless-function-s3keyfilter-rules
	//
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

