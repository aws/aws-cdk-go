package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationLocationProperty := &ApplicationLocationProperty{
//   	ApplicationId: jsii.String("applicationId"),
//   	SemanticVersion: jsii.String("semanticVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-application-applicationlocation.html
//
type CfnApplication_ApplicationLocationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-application-applicationlocation.html#cfn-serverless-application-applicationlocation-applicationid
	//
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-application-applicationlocation.html#cfn-serverless-application-applicationlocation-semanticversion
	//
	SemanticVersion *string `field:"required" json:"semanticVersion" yaml:"semanticVersion"`
}

