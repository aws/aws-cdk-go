package awscdkamplifyalpha


// Available hosting platforms to use on the App.
//
// Example:
//   amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &AppProps{
//   	Platform: amplify.Platform_WEB_COMPUTE,
//   })
//
// Experimental.
type Platform string

const (
	// WEB - Used to indicate that the app is hosted using only static assets.
	// Experimental.
	Platform_WEB Platform = "WEB"
	// WEB_COMPUTE - Used to indicate the app is hosted using a combination of server side rendered and static assets.
	// Experimental.
	Platform_WEB_COMPUTE Platform = "WEB_COMPUTE"
	// WEB_DYNAMIC - Used to indicate the app is hosted using a fully dynamic architecture, where requests are processed at runtime by backend compute services.
	// Experimental.
	Platform_WEB_DYNAMIC Platform = "WEB_DYNAMIC"
)

