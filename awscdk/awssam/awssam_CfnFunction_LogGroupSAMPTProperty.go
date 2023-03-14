package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logGroupSAMPTProperty := &logGroupSAMPTProperty{
//   	logGroupName: jsii.String("logGroupName"),
//   }
//
type CfnFunction_LogGroupSAMPTProperty struct {
	// `CfnFunction.LogGroupSAMPTProperty.LogGroupName`.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
}

