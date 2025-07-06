# gator
This is am RSS feed aggregator. It's a work in progress that started as guided project of the glorious Boot.dev, and I'm thinking to expand its capabilites. As of right now it does work, but is limited. It will run RSS feeds in your terminal.

## requirments
You'll need Go 1.17 or newer, and Postgres 

## install it

go install github.com/felixsolom/gator@latest

### add this to your path
export PATH="$PATH:$(go env GOBIN)"
#### or, if GOBIN is not set:
export PATH="$PATH:$HOME/go/bin"

## run it
GATOR has very simple user authentiaction. You just run 
### register _user_
and your done
### login 
works as you expect, no password needed at this point 
### addfeed
adds the feed by _single word name_ followed by _url_
it also follows the feed for the user that's logged in 
### follow 
allows you lo follow  by _url_ feeds that other users already follow
### following
displays the feeds that you follow
### unfollow
unfollows specific feed by adding its _url_ to the command
### agg 
is the forever green running command that aggregates your feeds by freshest.
it needs a time refresh constraint. 2s or 10h, or anything in the middle. it can run in separate 
### browse
needs a number of posts added to it, to know how many posts to display to you in the terminal

This is so far.. Enjoy!

