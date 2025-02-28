package customresources


// Properties used to initialize Logging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   loggingProps := &LoggingProps{
//   	LogApiResponseData: jsii.Boolean(false),
//   }
//
type LoggingProps struct {
	// Whether or not to log data associated with the API call response.
	// Default: true.
	//
	LogApiResponseData *bool `field:"optional" json:"logApiResponseData" yaml:"logApiResponseData"`
}

