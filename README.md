#monitor.go

##Description:
  Monitoring running docker container's network bandwidth and report to elasticsearch server.

##Usage:
sudo run monitor.go

#throttle.go

##Description:
  Set the download or upload speed limit of a container using tc and container's network namespace.

##Usage:
sudo run throttle.go 

#weighted.go

##Description:
  Throttle container's bandwidth based on weighted shares.

##Usage:
sudo run weighted.go

The content of map to be passed should look like this:

	container_1 weight_1
	container_2 weight_2
	...

For example, the following configuration will force container c1 use 50% of bandwidth, c2 and c3 share the remaining 50%.

	c1 2
	c2 1
	c3 1
Bandwidth to be shared between containers should be passed as a seperate value
