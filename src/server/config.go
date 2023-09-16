package server

/*
	this file describes server configurations
*/

type S3rverConfig struct {
	/* networking config */
	TLS  bool
	Port int

	/* filesystem config */
	Root string
}
