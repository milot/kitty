# Experiment Codename Kitty

I am publishing a number of experiments that I have done in the past here on Github instead of leaving them on my hard drives to collect dust in hopes that it helps somebody. I have experimented with this project two years ago on my spare time just to see how we could be building this using golang.

Codename kitty is a simple piece of software that is used in distributed systems in a leader/follower fashion. Basically what this software does is supports all followers to send messages to the leader. The followers have the ability to reconnect if a leader is lost. However, this leaves room for improvement i.e. having a consensus algorithm to automatically elect the leader if one leader goes down etc.

You have to build each program on its own i.e. `go build leader.go` and `go build follower.go` appropriately.