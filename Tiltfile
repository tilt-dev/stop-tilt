# these resources CAN be stopped via `stop-tilt`
local_resource('hammer-time', serve_cmd='./my-script.sh 🔨')
local_resource('in-the-name-of-love', serve_cmd='./my-script.sh ❤️')

# this resources CANNOT be stopped via `stop-tilt` (because it has no serve_cmd)
local_resource('dont-stop-me-now', 'echo i\'m havin\' such a good time')
