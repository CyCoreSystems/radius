# radius

Go RADIUS client library

RADIUS Accounting / RFC2866

## Client usage

	cl := radius.NewClient(&radius.Opts{
		Host: "",
		SharedSecret: "",
	})

	// send arbitrary packets
	resp, err := cl.Send(...)

## Session Usage
	
	id := ...

	// create a new accounting session
	sess := cl.NewSession(id, radius.TextString(radius.AccessUsername, "username"))

	// add some data to the session
	sess.AddInputOctet(12)
	sess.AddOutputOctet(33)

	// close the session
	sess.Stop(radius.UserRequest)
