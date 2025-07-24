package awsstepfunctions


// Values allowed in the retrier JitterStrategy field.
//
// Example:
//   parallel := sfn.NewParallel(this, jsii.String("Do the work in parallel"))
//
//   // Add branches to be executed in parallel
//   shipItem := sfn.NewPass(this, jsii.String("ShipItem"))
//   sendInvoice := sfn.NewPass(this, jsii.String("SendInvoice"))
//   restock := sfn.NewPass(this, jsii.String("Restock"))
//   parallel.Branch(shipItem)
//   parallel.Branch(sendInvoice)
//   parallel.Branch(restock)
//
//   // Retry the whole workflow if something goes wrong with exponential backoff
//   parallel.AddRetry(&RetryProps{
//   	MaxAttempts: jsii.Number(1),
//   	MaxDelay: awscdk.Duration_Seconds(jsii.Number(5)),
//   	JitterStrategy: sfn.JitterType_FULL,
//   })
//
//   // How to recover from errors
//   sendFailureNotification := sfn.NewPass(this, jsii.String("SendFailureNotification"))
//   parallel.AddCatch(sendFailureNotification)
//
//   // What to do in case everything succeeded
//   closeOrder := sfn.NewPass(this, jsii.String("CloseOrder"))
//   parallel.Next(closeOrder)
//
type JitterType string

const (
	// Calculates the delay to be a random number between 0 and the computed backoff for the given retry attempt count.
	JitterType_FULL JitterType = "FULL"
	// Calculates the delay to be the computed backoff for the given retry attempt count (equivalent to if Jitter was not declared - i.e. the default value).
	JitterType_NONE JitterType = "NONE"
)

