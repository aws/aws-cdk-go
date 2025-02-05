package awsappintegrations


// The configuration for where the application should be loaded from.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   applicationSourceConfigProperty := &ApplicationSourceConfigProperty{
//   	ExternalUrlConfig: &ExternalUrlConfigProperty{
//   		AccessUrl: jsii.String("accessUrl"),
//
//   		// the properties below are optional
//   		ApprovedOrigins: []*string{
//   			jsii.String("approvedOrigins"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-application-applicationsourceconfig.html
//
type CfnApplication_ApplicationSourceConfigProperty struct {
	// The external URL source for the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-application-applicationsourceconfig.html#cfn-appintegrations-application-applicationsourceconfig-externalurlconfig
	//
	ExternalUrlConfig interface{} `field:"required" json:"externalUrlConfig" yaml:"externalUrlConfig"`
}

