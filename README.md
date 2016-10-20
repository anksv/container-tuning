#monitor.go

##Description:
  Monitoring running docker container's network bandwidth and report to elasticsearch server.

##Usage:
monitor.go [OPTIONS]

##Options:
	-n, --name            Only monitor the container with given name
	-h, --help            This page


#throttle.go

##Description:
  Set the download or upload speed limit of a container using tc and container's network namespace.

##Usage:
throttle.go CONTAINER_NAME [OPTIONS]

##Options:
	-d, --download    #kmKM        Set download speed limit in specified unit
	-u, --upload      #kmKM        Set upload speed limit in specified unit
	-c, --clean                    Clean the speed limits
	-h, --help                     This page

#weighted.go

##Description:
  Throttle container's bandwidth based on weighted shares.

##Usage:
weighted.go [-f filename]

##Options:
	-f, --file        Read configuration from specified file
	-t, --total       Set total bandwidth in megabits
	-h, --help        This page

The content of configuration file should look like this:

	container_1 weight_1
	container_2 weight_2
	...

For example, the following configuration will force container c1 use 50% of bandwidth, c2 and c3 share the remaining 50%.

	c1 2
	c2 1
	c3 1
