package awsecs


// Firelens log router type, fluentbit or fluentd.
//
// https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html
// Experimental.
type FirelensLogRouterType string

const (
	// fluentbit.
	// Experimental.
	FirelensLogRouterType_FLUENTBIT FirelensLogRouterType = "FLUENTBIT"
	// fluentd.
	// Experimental.
	FirelensLogRouterType_FLUENTD FirelensLogRouterType = "FLUENTD"
)

