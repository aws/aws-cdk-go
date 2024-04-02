package awsappintegrations


// The external URL source for the application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   externalUrlConfigProperty := &ExternalUrlConfigProperty{
//   	AccessUrl: jsii.String("accessUrl"),
//   	ApprovedOrigins: []*string{
//   		jsii.String("approvedOrigins"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-application-externalurlconfig.html
//
type CfnApplication_ExternalUrlConfigProperty struct {
	// The URL to access the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-application-externalurlconfig.html#cfn-appintegrations-application-externalurlconfig-accessurl
	//
	AccessUrl *string `field:"required" json:"accessUrl" yaml:"accessUrl"`
	// Additional URLs to allow list if different than the access URL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-application-externalurlconfig.html#cfn-appintegrations-application-externalurlconfig-approvedorigins
	//
	ApprovedOrigins *[]*string `field:"required" json:"approvedOrigins" yaml:"approvedOrigins"`
}

