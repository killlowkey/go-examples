## Init Project
```shell
mkdir pprof
cd pprof-example
go mod init pprof-example
```

## Install ab
access https://www.apachelounge.com/download/ website, select ab version for your operating system.
in windows, after downloaded it, unzip it, add unzip path adding bin dir to windows environment, such as 'E:\Apache24\bin'

## Install graphviz 
> https://gitlab.com/api/v4/projects/4207231/packages/generic/graphviz-releases/10.0.1/windows_10_cmake_Release_Graphviz-10.0.1-win64.zip
access https://graphviz.org/download/ website, select and download it, add unzip path to windows environment


## Usage
> https://www.codereliant.io/memory-leaks-with-pprof/
> execute web command in pprof, to view the profile image in web browser
```shell
$ ab -n 1000 -c 10 http://localhost:8080/leaky-endpoint
...
Finished 1000 requests
$ go tool pprof -alloc_space http://localhost:8080/debug/pprof/heap
(pprof) top 
Showing nodes accounting for 348.95MB, 99.43% of 350.96MB total 
Dropped 14 nodes (cum <= 1.75MB)
      flat  flat%   sum%        cum   cum%
  348.95MB 99.43% 99.43%   349.45MB 99.57%  main.handleRequest
         0     0% 99.43%   349.45MB 99.57%  net/http.(*ServeMux).ServeHTTP
         0     0% 99.43%   350.96MB   100%  net/http.(*conn).serve
         0     0% 99.43%   349.45MB 99.57%  net/http.HandlerFunc.ServeHTTP
         0     0% 99.43%   349.45MB 99.57%  net/http.serverHandler.ServeHTTP
(pprof) web
(pprof) list handleRequest 
Total: 350.96MB                                                                           
ROUTINE ======================== main.handleRequest in F:\go-examples\pprof\main.go       
  348.95MB   349.45MB (flat, cum) 99.57% of Total                                         
         .          .     28:func handleRequest(w http.ResponseWriter, r *http.Request) { 
         .          .     29:   userCache.mu.Lock()                                       
         .          .     30:   defer userCache.mu.Unlock()                               
         .          .     31:                                                             
         .          .     32:   userData := &UserData{                                    
  348.95MB   348.95MB     33:           Data: make([]byte, 1000000),                      
         .          .     34:   }                                                         
         .          .     35:                                                             
         .          .     36:   userID := fmt.Sprintf("%d", len(userCache.Cache))         
         .          .     37:   userCache.Cache[userID] = userData
         .   512.50kB     38:   log.Printf("Added data for user %s. Total users: %d\n", userID, len(userCache.Cache))
         .          .     39:}
         .          .     40:
         .          .     41:func main() {
         .          .     42:   http.HandleFunc("/leaky-endpoint", handleRequest)
         .          .     43:   http.ListenAndServe(":8080", nil)
```

### Pprof Commands
```shell
(pprof) help 
  Commands: 
    callgrind        Outputs a graph in callgrind format
    comments         Output all profile comments
    disasm           Output assembly listings annotated with samples
    dot              Outputs a graph in DOT format
    eog              Visualize graph through eog
    evince           Visualize graph through evince
    gif              Outputs a graph image in GIF format
    gv               Visualize graph through gv
    kcachegrind      Visualize report in KCachegrind
    list             Output annotated source for functions matching regexp
    pdf              Outputs a graph in PDF format
    peek             Output callers/callees of functions matching regexp
    png              Outputs a graph image in PNG format
    proto            Outputs the profile in compressed protobuf format
    ps               Outputs a graph in PS format
    raw              Outputs a text representation of the raw profile
    svg              Outputs a graph in SVG format
    tags             Outputs all tags in the profile
    text             Outputs top entries in text form
    top              Outputs top entries in text form
    topproto         Outputs top entries in compressed protobuf format
    traces           Outputs all profile samples in text form
    tree             Outputs a text rendering of call graph
    web              Visualize graph through web browser
    weblist          Display annotated source in a web browser
    o/options        List options and their current values
    q/quit/exit/^D   Exit pprof

  Options:
    call_tree        Create a context-sensitive call tree
    compact_labels   Show minimal headers
    divide_by        Ratio to divide all samples before visualization
    drop_negative    Ignore negative differences
    edgefraction     Hide edges below <f>*total
    focus            Restricts to samples going through a node matching regexp
    hide             Skips nodes matching regexp
    ignore           Skips paths going through any nodes matching regexp
    intel_syntax     Show assembly in Intel syntax
    mean             Average sample value over first value (count)
    nodecount        Max number of nodes to show
    nodefraction     Hide nodes below <f>*total
    noinlines        Ignore inlines.
    normalize        Scales profile based on the base profile.
    output           Output filename for file-based outputs
    prune_from       Drops any functions below the matched frame.
    relative_percentages Show percentages relative to focused subgraph
    sample_index     Sample value to report (0-based index or name)
    show             Only show nodes matching regexp
    show_from        Drops functions above the highest matched frame.
    source_path      Search path for source files
    tagfocus         Restricts to samples with tags in range or matched by regexp
    taghide          Skip tags matching this regexp
    tagignore        Discard samples with tags in range or matched by regexp
    tagleaf          Adds pseudo stack frames for labels key/value pairs at the callstack leaf.
    tagroot          Adds pseudo stack frames for labels key/value pairs at the callstack root.
    tagshow          Only consider tags matching this regexp
    trim             Honor nodefraction/edgefraction/nodecount defaults
    trim_path        Path to trim from source paths before search
    unit             Measurement units to display

  Option groups (only set one per group):
    granularity
      functions        Aggregate at the function level.
      filefunctions    Aggregate at the function level.
      files            Aggregate at the file level.
      lines            Aggregate at the source code line level.
      addresses        Aggregate at the address level.
    sort
      cum              Sort entries based on cumulative weight
      flat             Sort entries based on own weight
  :   Clear focus/ignore/hide/tagfocus/tagignore

  type "help <cmd|option>" for more information
```

### Profiling CPU Usage
```shell
$ go tool pprof -http=localhost:8081 http://localhost:8080/debug/pprof/goroutine
$ go tool pprof -http=localhost:8081 http://localhost:8080/debug/pprof/threadcreate
```

### Profiling Mutex Contention
```shell
$ go tool pprof -http=localhost:8081 http://localhost:8080/debug/pprof/mutex
```

### Heap Diff Profiling
```shell
$ curl -s http://localhost:8080/debug/pprof/heap > base.heap
$ curl -s http://localhost:8080/debug/pprof/heap > current.heap
$ go tool pprof -http=localhost:8081 --base base.heap current.heap
```


## reference
> look https://www.codereliant.io/memory-leaks-with-pprof/ in more detail 
1. https://www.codereliant.io/memory-leaks-with-pprof/
2. https://www.apachelounge.com/download/
3. https://community.chocolatey.org/packages/apache-httpd#files
4. [How to install apache ab ?](https://gist.github.com/yolossn/20d86c79745acbd97125b9cca950cbf7)
5. [Diagnostics](https://go.dev/doc/diagnostics)