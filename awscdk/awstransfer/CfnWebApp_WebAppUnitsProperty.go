package awstransfer


// Contains an integer value that represents the value for number of concurrent connections or the user sessions on your web app.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webAppUnitsProperty := &WebAppUnitsProperty{
//   	Provisioned: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-webapp-webappunits.html
//
type CfnWebApp_WebAppUnitsProperty struct {
	// An integer that represents the number of units for your desired number of concurrent connections, or the number of user sessions on your web app at the same time.
	//
	// Each increment allows an additional 250 concurrent sessions: a value of `1` sets the number of concurrent sessions to 250; `2` sets a value of 500, and so on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-webapp-webappunits.html#cfn-transfer-webapp-webappunits-provisioned
	//
	Provisioned *float64 `field:"required" json:"provisioned" yaml:"provisioned"`
}

