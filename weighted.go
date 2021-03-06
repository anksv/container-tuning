package main

import(
"fmt"
)

func process_cmd(str []string) (string){

	cmd := exec.Command(str[0],str[1:]...)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	  if err != nil {
         log.Printf(err.Error())
      } else {
         //log.Printf("success %s", string(cmdOutput.Bytes()))
         return string(cmdOutput.Bytes())
      }
   return ""

}

func get_veth( cname string) string {
	var err error
      vethid := 0
      veth := ""
     // veth0 := ""
      _vethid := ""
	iplink_cmd := []string{"sudo", "ip", "netns", "exec", "cname", "ip", "link", "show",}
    iplink := process_cmd(iplink_cmd)
    result := strings.Split(iplink, "\n")

    for i := range result {
        if strings.Contains(result[i], " eth0:"){   
          res := strings.Split(result[i],":")
          vethid,err = strconv.Atoi(res[0])
          if err == nil {
           vethid = vethid + 1
           _vethid = strconv.Itoa(vethid)
    	  }
        }
    }

    iplink_cmd1 := []string{"sudo", "ip", "link", "show",}
    iplink1 := process_cmd(iplink_cmd1)
    result1 := strings.Split(iplink1, "\n")
    for i := range result1 {
    	res := strings.Split(result1[i],":")
        if res[0] == _vethid {
           veth := res[1][1:]
           at := strings.Index(veth, "@")
           if at == -1 {
             at = 0
           }
           if at == 0 {
               veth =  veth[:]
           }else {
               veth =  veth[:at]
           } 
        }
    }
    return veth
}



func link_netns(finalmap map[string]string){
    
	type State struct {
		        Pid    int  `json:"pid"`
	}

	type Container struct {
		        ConState   State      `json:"state"`
		        Name       string     `json:"name"`
	}
    
    //If the container is not found, it gets removed from the map
    mk_dir := []string{"sudo", "mkdir", "-p", "/var/run/netns"}
    
    process_cmd(mk_dir)

    for cid,_ := range finalmap {
      
	  fmt.Println(cid)
      inspect_data := []string{"docker" , "inspect", cid}
      result := process_cmd(inspect_data)
	  
	    if result != "" {
	  
      b := []byte(result)

      test := make([]Container, 0)
      json.Unmarshal(b, &test)
      //fmt.Println("json\n",json[0].Name, "\n", json[0].ConState.Pid)
      
      pid := test[0].ConState.Pid
      _pid := strconv.Itoa(pid)

      cname := test[0].Name
      
      rm_dir := []string{"sudo", "rm", "/var/run/netns/", cname}
      symb_link := []string{ "sudo", "ln", "-s", "/proc/", _pid, "/ns/net", "/var/run/netns/", cname}
      process_cmd(rm_dir)
      process_cmd(symb_link)
	    }else{
		    delete(finalmap,cid)
		    continue
	    }

    }

}

func allot_netwwork_share(finalmap map[string]string, rnb int) {
     
    total := 100
    //total = rnb

	link_netns(finalmap)

	total_weight := 0

	for key,_ := range finalmap {
		 val,_ := strconv.Atoi(finalmap[key])
		 total_weight += val
	}

	for cid,_ := range finalmap {

		veth := get_veth(cid)
                v,_ := strconv.Atoi(finalmap[cid])
		drate := total * 1024 * 1024 * v / total_weight
		_drate := strconv.Itoa(drate)
        d := drate/250
		_d := strconv.Itoa(d)

		del_cmd := []string{"sudo" , "tc" , "qdisc" ,  "del" ,  "dev" , veth, "root"}
		add_cmd := []string{"sudo" , "tc" ,  "qdisc" ,  "add" ,  "dev" , veth, "root", "tbf", "rate", _drate, "bit latency 50ms burst", _d}

		process_cmd(del_cmd)

		process_cmd(add_cmd)


}
	func main() {
	        //Pass a map whose keys are container and value as network share , bandwidth to be distributed
		allot_netwwork_share(finalmap, bandwidth)
		
	}
