package awsroute53


// Properties for a CAA record value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   caaRecordValue := &caaRecordValue{
//   	flag: jsii.Number(123),
//   	tag: awscdk.Aws_route53.caaTag_ISSUE,
//   	value: jsii.String("value"),
//   }
//
type CaaRecordValue struct {
	// The flag.
	Flag *float64 `field:"required" json:"flag" yaml:"flag"`
	// The tag.
	Tag CaaTag `field:"required" json:"tag" yaml:"tag"`
	// The value associated with the tag.
	Value *string `field:"required" json:"value" yaml:"value"`
}

