package previewawsappintegrationsmixins


// The configuration for where the application should be loaded from.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   applicationSourceConfigProperty := &ApplicationSourceConfigProperty{
//   	ExternalUrlConfig: &ExternalUrlConfigProperty{
//   		AccessUrl: jsii.String("accessUrl"),
//   		ApprovedOrigins: []*string{
//   			jsii.String("approvedOrigins"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-application-applicationsourceconfig.html
//
type CfnApplicationPropsMixin_ApplicationSourceConfigProperty struct {
	// The external URL source for the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appintegrations-application-applicationsourceconfig.html#cfn-appintegrations-application-applicationsourceconfig-externalurlconfig
	//
	ExternalUrlConfig interface{} `field:"optional" json:"externalUrlConfig" yaml:"externalUrlConfig"`
}

