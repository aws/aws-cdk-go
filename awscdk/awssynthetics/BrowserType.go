package awssynthetics


// Browser types supported by CloudWatch Synthetics Canary.
//
// Example:
//   canary := synthetics.NewCanary(this, jsii.String("MyCanary"), &CanaryProps{
//   	Schedule: synthetics.Schedule_Rate(awscdk.Duration_Minutes(jsii.Number(5))),
//   	Test: synthetics.Test_Custom(&CustomTestOptions{
//   		Code: synthetics.Code_FromAsset(path.join(__dirname, jsii.String("canary"))),
//   		Handler: jsii.String("index.handler"),
//   	}),
//   	Runtime: synthetics.Runtime_SYNTHETICS_NODEJS_PUPPETEER_9_1(),
//   	BrowserConfigs: []browserType{
//   		synthetics.*browserType_CHROME,
//   		synthetics.*browserType_FIREFOX,
//   	},
//   })
//
type BrowserType string

const (
	// Google Chrome browser.
	BrowserType_CHROME BrowserType = "CHROME"
	// Mozilla Firefox browser.
	BrowserType_FIREFOX BrowserType = "FIREFOX"
)

