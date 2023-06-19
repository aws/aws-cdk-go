package awsroute53


// Properties for a CAA record value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   caaRecordValue := &CaaRecordValue{
//   	Flag: jsii.Number(123),
//   	Tag: awscdk.Aws_route53.CaaTag_ISSUE,
//   	Value: jsii.String("value"),
//   }
//
// Experimental.
type CaaRecordValue struct {
	// The flag.
	// Experimental.
	Flag *float64 `field:"required" json:"flag" yaml:"flag"`
	// The tag.
	// Experimental.
	Tag CaaTag `field:"required" json:"tag" yaml:"tag"`
	// The value associated with the tag.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

