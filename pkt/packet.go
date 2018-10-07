package main

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"os/exec"
	"bytes"
	"strings"
)

func help(){
	fmt.Println("Usage for packet tool.");
	fmt.Println("    install 安装依赖包");
	fmt.Println("    build 编译主程序");
}

func exec_shell(exe string,s ...string) (string, error){
    //函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
    cmd := exec.Command(exe,s...)

    //读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
    var out bytes.Buffer
    cmd.Stdout = &out

    //Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
    err := cmd.Run()

    return out.String(), err
}

func exec_cmd(cmd []string){
	command := "go " + strings.Trim(fmt.Sprint(cmd), "[]");

	out,err := exec_shell("go",cmd...);
	if (err != nil){
		fmt.Println(command);
		fmt.Println(err);
	}else {
		if(len(out) > 0){
			fmt.Println(command);
			fmt.Println(out);
		}else{
			fmt.Print(command);
			fmt.Println(".");
		}
	}
}

func install(name string,t string) {
	var cmd []string;
	if t == "u" {
		cmd = make([]string,3)
		cmd[0] = "install"
		cmd[1] = "-u"
		cmd[2] = name;
	} else if(t == "g"){
		cmd = make([]string,3)
		cmd[0] = "install"
		cmd[1] = "-g"
		cmd[2] = name;
	} else {
		cmd = make([]string,2)
		cmd[0] = "install"
		cmd[1] = name;
	}

	exec_cmd(cmd);
}

func main(){
	args := os.Args
	if len(args) == 1 {
		help();
		return;
	}
	
	str,err := ioutil.ReadFile("packet.json");
	if(err != nil){
		fmt.Print("io error:");
		fmt.Print(err);
		return;
	}

	var obj interface{}
	err = json.Unmarshal(str,&obj);
	if(err != nil){
		fmt.Print("json error:");
		fmt.Print(err);
		return;
	}
	//fmt.Println(obj);

	param := args[1];
	fmt.Println(param);
	if param == "install" {
		jtable := obj.(map[string]interface{});
		dep := jtable["dependencies"];
		if(dep == nil){
			return;
		}
		deps := dep.(map[string]interface{});
		for k,v := range deps {
			install(k,v.(string));
		}
	} else if(param == "build"){
		jtable := obj.(map[string]interface{});
		name := jtable["name"];
		main := jtable["main"];

		if(main == nil){
			return;
		}
		c := 2;
		if(name != nil){
			c+=2;
		}
		var cmd []string;
		cmd = make([]string,c)
		i := 0;
		cmd[i] = "build";
		i++;
		if(name != nil){
			cmd[i] = "-o"
			i++;
			cmd[i] = name.(string) + ".exe";
			i++;
		}
		cmd[i] = main.(string);
		
		exec_cmd(cmd);
	}
}
