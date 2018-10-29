# elevator
A golang library for requesting self-sudo.

Simply put, the idea is in the runtime, if the process is not started with superuser permissions, request them
and then restart the application.

To supply the sudo username and password during development locally, just use the environment variables defined as,

SUDO_PASSWORD=12341234

As you should already know, what i described a moment ago is a terrible idea, and you should never do it. But, if
you feel like not entering the password again and again on every restart while debugging your application, it's
actually a quite useful tool. Maybe I'll provide more elegant ways of supplying a default password and username in
future.


### Contributing
You can create a PR and as long as it makes sense, and doesnt' enlarge the scope of the project by 2, it'll be merged.