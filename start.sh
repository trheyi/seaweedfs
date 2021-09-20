#!/usr/bin/osascript
# https://superuser.com/questions/647672/split-window-programmatically-with-iterm
#
on run argv
    tell application "iTerm"
        tell current window
            create tab with default profile
            tell current session
                write text "cd /Users/max/Code/seaweedfs/" 
                write text "weed/weed master"
                split horizontally with default profile
            end tell

            tell last session of last tab
                write text "sleep 5"
                write text "cd /Users/max/Code/seaweedfs/"
                write text "weed/weed volume -dir=\"/Users/max/Code/seaweedfs/tmp/data\" -max=5  -mserver=\"localhost:9333\" -port=8080"
                split horizontally with default profile
            end tell

            tell last session of last tab
                write text "sleep 5"
                write text "cd /Users/max/Code/seaweedfs/"
                write text "weed/weed filer"
                split horizontally with default profile
            end tell

            tell last session of last tab
                write text "sleep 5"
                write text "cd /Users/max/Code/seaweedfs/"
                write text "weed/weed s3 -filer=\"localhost:8888\""
                split horizontally with default profile
            end tell
            
            tell last session of last tab
                write text "sleep 5"
                write text "cd /Users/max/Code/seaweedfs/"
                write text "weed/weed webdav -filer=\"localhost:8888\""
            end tell

        end tell
    end tell 
end

# #!/bin/bash
# weed/weed volume -dir="/Users/max/Code/seaweedfs/tmp/data" -max=5  -mserver="localhost:9333" -port=8080
# weed/weed server -dir="/Users/max/Code/seaweedfs/tmp/data" -s3
# weed/weed webdav -dir="/Users/max/Code/seaweedfs/tmp/data"
# weed volume -dir="/tmp/data1" -max=5  -mserver="localhost:9333" -port=8080
# weed/weed master
# weed/weed server -dir=/Users/max/Code/seaweedfs/tmp/data -s3

# weed/weed volume -dir="/Users/max/Code/seaweedfs/tmp/data" -max=5  -mserver="localhost:9333" -port=8081