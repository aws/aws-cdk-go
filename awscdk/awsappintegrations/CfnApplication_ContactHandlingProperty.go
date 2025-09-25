package awsappintegrations


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contactHandlingProperty := &ContactHandlingProperty{
//   	Scope: jsii.String("scope"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-application-contacthandling.html
//
type CfnApplication_ContactHandlingProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-application-contacthandling.html#cfn-appintegrations-application-contacthandling-scope
	//
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

