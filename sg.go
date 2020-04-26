package main

import (
    "fmt"
    "log"
    "time"
    "strings"
    "io/ioutil"
)

var (
    cflags = "-Wall -Wextra -Werror -Ofast"
)

type Target struct {
    Name string
    Lang string
    Inputs []string
    Includes []string
}

func assertErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func buildTarget(contents []string) *Target {
    previous := ""
    var target, lang string
    inputs := make([]string, 0, 2)
    includes := make([]string, 0, 2)
    for _, str := range contents {
        switch previous {
            case "target":
                target = str
            
            case "lang":
                lang = str
            
            case "input":
                inputs = append(inputs, str)
            
            case "include":
                includes = append(includes, str)
        }
        
        previous = str
    }
    
    return &Target{Name: target, Lang: lang, Inputs: inputs, Includes: includes}
}

func gen(target *Target, compiler string) []byte {
    buf := "ninja_required_version = 1.8\n"
    
    buf += "cflags = " + cflags + "\n"
    
    buf += "cincludes = "
    for _, include := range target.Includes {
        buf += "-I" + include + " "
    }
    buf += "\n"
    
    buf += "\nrule cc\n  command = " + compiler + " -MMD -MF $out.d $cflags $cincludes -c $in -o $out\n\n"
    buf += "\nrule cl\n  command = " + compiler + " -MMD -MF $out.d $cflags $cincludes -o $out $in\n\n"
    
    inputs := ""
    for _, in := range target.Inputs {
        buf += "build " + in + ".o: cc ../" + in + "\n"
        inputs += in + ".o "
    }
    
    buf += "build " + target.Name + ": cl " + inputs + "\n"
    buf += "\ndefault " + target.Name + "\n"
    
    return []byte(buf)
}

func main() {
    t := time.Now()
    
    buildFile, err := ioutil.ReadFile("build.sg")
    assertErr(err)
    
    target := buildTarget(strings.Fields(string(buildFile)))
    
    compiler := ""
    switch target.Lang {
        case "c":
            compiler = "gcc"
        case "cxx":
            compiler = "g++"
    }
    
    err = ioutil.WriteFile("out/build.ninja", gen(target, compiler), 0755)
    assertErr(err)
    
    fmt.Println("Target built in " + time.Now().Sub(t).String() + " milliseconds.")
}