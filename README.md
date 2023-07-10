## Steps 
1. Create a plugin 
`go build -buildmode=plugin -o <plugin_location> <plugin_files_path>`
<br>Example:
`go build -buildmode=plugin -o ./plugins/parser/parser.so ./plugins/parser/parser.go`

2. Use plugin
`go run main.go <plugin_name>`


## Reference links :
https://pkg.go.dev/plugin
https://dev.to/jacktt/plugin-in-golang-4m67
