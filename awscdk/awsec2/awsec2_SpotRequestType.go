package awsec2


// The Spot Instance request type.
// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-requests.html
//
type SpotRequestType string

const (
	// A one-time Spot Instance request remains active until Amazon EC2 launches the Spot Instance, the request expires, or you cancel the request.
	//
	// If the Spot price exceeds your maximum price
	// or capacity is not available, your Spot Instance is terminated and the Spot Instance request
	// is closed.
	SpotRequestType_ONE_TIME SpotRequestType = "ONE_TIME"
	// A persistent Spot Instance request remains active until it expires or you cancel it, even if the request is fulfilled.
	//
	// If the Spot price exceeds your maximum price or capacity is not available,
	// your Spot Instance is interrupted. After your instance is interrupted, when your maximum price exceeds
	// the Spot price or capacity becomes available again, the Spot Instance is started if stopped or resumed
	// if hibernated.
	SpotRequestType_PERSISTENT SpotRequestType = "PERSISTENT"
)

