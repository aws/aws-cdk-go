package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainSAMPTProperty := &DomainSAMPTProperty{
//   	DomainName: jsii.String("domainName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-domainsampt.html
//
type CfnFunction_DomainSAMPTProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-domainsampt.html#cfn-serverless-function-domainsampt-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
}

