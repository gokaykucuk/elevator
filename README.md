# elevator
A golang library for requesting self-sudo.

Simply put, the idea is in the runtime, if the process is not started with superuser permissions, request them
and then restart the application. The only rule is, this libraries `CheckAndRequestElevation()` function *MUST* be
called as the first thing before the main application loads.

To supply the sudo username and password during development locally, just use the environment variables defined as,

    ELEVATOR_SUDO_PASSWORD=12341234

As you should already know, what i described a moment ago is a terrible idea, and you should never do it. But, if
you feel like not entering the password again and again on every restart while debugging your application, it's
actually a quite useful tool. Maybe I'll provide more elegant ways of supplying a default password and username in
future.


### Installation

Just as any other go package, you can get it with `go get`

    go get -u github.com/gokaykucuk/elevator


### Example

    package main
    
    import (
    	"github.com/gokaykucuk/elevator"
    	"gitlab.com/..."
    )
    
    var (
    	version   = "No version provided"
    	commit    = "No commit provided"
    	buildTime = "No build timestamp provided"
    )
    
    func main() {
    	elevator.CheckAndRequestElevation()
    	cmd.SetVersion(&cmd.Version{
    		Version:   version,
    		Commit:    commit,  
    		BuildTime: buildTime,
    	})
    	cmd.Execute()
    }


### Contributing
You can create a PR and as long as it makes sense, and doesnt' enlarge the scope of the project by 2, it'll be merged.

### Testing
Currently there are no real tests because i don't have a good idea how to write them. If you have any good ideas please
open a pull request.

### TODO
- Make it windows compatibale?
- Make another function which will only check but not ask for permissions